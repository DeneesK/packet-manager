# Packet Manager

[![Go Test](https://github.com/Deneesk/packet-manager/actions/workflows/pm-tests.yml/badge.svg)](https://github.com/Deneesk/packet-manager/actions/workflows/pm-tests.yml.yml)
![Go version](https://img.shields.io/badge/go-1.23-blue)

Утилита для пакетной упаковки файлов, загрузки и скачивания архивов по SSH с поддержкой JSON/YAML конфигураций.

---

## Возможности

- Создание zip-архивов из файлов, заданных в конфигурационном JSON/YAML
- Загрузка архивов на удалённый сервер по SSH
- Скачивание и распаковка архивов с сервера по SSH
- Поддержка версии пакетов и зависимостей
- Удобный CLI с командами `pm create` и `pm update`

---

## Быстрый старт

заполните ```config.yaml```

```yaml
host: "example.com"
user: "deploy"
key: "/path/to/private_key"
remote_dir: "/remote/archive/dir"
```

### Установка

Склонируйте репозиторий и соберите бинарник:

```bash
git clone https://github.com/yourusername/packet-manager.git
cd packet-manager
go build -o pm main.go
```

### Использование
- *Создать архив по описанию в packet.json | yaml:*\
```./pm create ./packet.json```\
```./pm create ./packet.json -c ./path-to-your-config.yaml```

- *Скачать и распаковать пакеты из packages.json | yaml:*\
```./pm update ./packet.json```\
```./pm update ./packet.json -c ./path-to-your-config.yaml```

### Available Commands:
- completion  Generate the autocompletion script for the specified shell
- create      Create archive and upload
- help        Help about any command
- update      Download and extract archives from server

- #### Flags:
  -c, --config string   Path to config file (default "./config.yaml")\
  -h, --help            help for pm

### Тесты
- Автоматический запуск тестов через GitHub Actions при push