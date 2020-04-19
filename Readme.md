# Тестовый апи на gRPC на golang

### Необходимое ПО:
**protoc** - компилятор `.proto` файлов  
```
## Установка

# protoc
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip
unzip protoc-3.11.4-linux-x86_64.zip -d $HOME/.local
export PATH="$PATH:$HOME/.local/bin"

# protoc-gen-go
go get github.com/golang/protobuf/protoc-gen-go
export PATH="$PATH:$GOPATH/bin"
```

**grpc** - сам фреймворк от Google
```
## Установка

export GO111MODULE=on # Enable module-aware mode
go get google.golang.org/grpc@v1.28.1
```

---

### Пример команд
**Генерация go исходников из .proto**
```
protoc -I /path/to/proto-files/dir/ /path/to/file.proto --go_out=plugins=grpc:/path/to/output/dir/
```