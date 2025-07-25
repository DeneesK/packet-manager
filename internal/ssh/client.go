package ssh

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/sftp"
	gossh "golang.org/x/crypto/ssh"
)

type SSHClient struct {
	Host string
	User string
	Key  string
}

func (c *SSHClient) connect() (*sftp.Client, error) {
	key, err := os.ReadFile(c.Key)
	if err != nil {
		return nil, fmt.Errorf("read key: %w", err)
	}
	signer, err := gossh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}

	config := &gossh.ClientConfig{
		User: c.User,
		Auth: []gossh.AuthMethod{
			gossh.PublicKeys(signer),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}

	conn, err := gossh.Dial("tcp", c.Host+":22", config)
	if err != nil {
		return nil, fmt.Errorf("dial ssh: %w", err)
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, fmt.Errorf("new sftp client: %w", err)
	}
	return client, nil
}

func (c *SSHClient) Upload(localPath, remotePath string) error {
	client, err := c.connect()
	if err != nil {
		return err
	}
	defer client.Close()

	srcFile, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := client.Create(remotePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func (c *SSHClient) Download(remotePath, localPath string) error {
	client, err := c.connect()
	if err != nil {
		return err
	}
	defer client.Close()

	srcFile, err := client.Open(remotePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func (c *SSHClient) ListFiles(remoteDir string) ([]string, error) {
	client, err := c.connect()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	files, err := client.ReadDir(remoteDir)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, f := range files {
		if !f.IsDir() && !strings.HasPrefix(f.Name(), ".") {
			result = append(result, f.Name())
		}
	}
	return result, nil
}
