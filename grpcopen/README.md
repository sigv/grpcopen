The single Protocol Buffer file that exists in this repository is already pre-compiled to Go code. However, in the rare case you need to modify the `.proto`, here are helpful instructions on how to compile those changes for Go. (The instructions are primarily based on [gRPC's official quick-start documentation].)

[gRPC's official quick-start documentation]: https://grpc.io/docs/languages/go/quickstart/#prerequisites

## Installing Go

```bash
# https://go.dev/dl/

wget -nv https://go.dev/dl/go1.21.1.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
rm go1.21.1.linux-amd64.tar.gz

source <(echo 'export PATH="$PATH:/usr/local/go/bin"' | tee -a /etc/profile)
source <(echo 'export PATH="$PATH:$(go env GOPATH)/bin"' | tee -a /etc/profile)
```

## Installing protoc (Protocol Buffer compiler)

```bash
# https://github.com/protocolbuffers/protobuf/releases/latest

wget -nv https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
unzip protoc-3.15.8-linux-x86_64.zip -d $HOME/.local
rm protoc-3.15.8-linux-x86_64.zip

source <(echo 'export PATH="$PATH:$HOME/.local/bin"' | tee -a ~/.profile)
```

## Installing Go plugins for protoc

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3
```

## Rebuilding `.pb.go` files

```bash
find . -type f -name '*.proto' -exec protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative {} \;
```
