# gRPC OPEN

Proof of Concept project that provides a basic gRPC server and a gRPC client that doesn't send `END_STREAM` flag right away with `HEADER` frame. This happens when the client wants to send additional messages to the same target server, without worrying about a load balancer directing it somewhere else.

The PoC project is intended to assist in investigating an issue with HAProxy, reported on the mailing list as [mux-h2: Backend stream is not fully closed if frontend keeps stream open]. The PoC application has two endpoints for this: `Ping` and `Foobar`.

The `Ping` endpoint is a basic ping-pong showing the happy path: its request payload carries a `Content` string, which the server returns in response payload. The gRPC endpoint is configured to support bi-directional streaming. Currently, only a single `Ping` request is sent, but this does not affect the issue on hand.

The `Foobar` endpoint is not implemented, and returns a gRPC status with code Unimplemented. The server does this right away, when it sees any client activity. The case may be that client is still sending data when this occurs, so the client has not yet closed the stream from its side. To ensure such a simulated situation, without actually sending a payload, we configure the client to sleep for a second before closing its sending side.

[mux-h2: Backend stream is not fully closed if frontend keeps stream open]: https://www.mail-archive.com/haproxy@formilux.org/msg44010.html

## How-to

### Configuring HAProxy

Building HAProxy is outside of the scope of this README. Consult HAProxy's [INSTALL] documentation for how to build it.

A base configuration that can be used for testing:

```text
global
   log stdout local0
   stats socket /run/haproxy/admin.sock mode 660 level admin expose-fd listeners

defaults
   mode http
   log global
   option httplog
   timeout connect 5000
   timeout client 50000
   timeout server 50000

frontend fe
   bind 127.0.0.1:8080 proto h2
   default_backend be

backend be
   server srv 127.0.0.1:8088 proto h2
```

### Installing Go

```bash
# https://go.dev/dl/

wget -nv https://go.dev/dl/go1.21.1.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
rm go1.21.1.linux-amd64.tar.gz

source <(echo 'export PATH="$PATH:/usr/local/go/bin"' | tee -a /etc/profile)
source <(echo 'export PATH="$PATH:$(go env GOPATH)/bin"' | tee -a /etc/profile)
```

### Starting the application

```bash
go run server/main.go -addr :8088
```

```bash
go run client/main.go -addr localhost:8088 # direct
go run client/main.go -addr localhost:8080 # HAProxy
```
