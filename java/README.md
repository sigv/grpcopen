# gRPC OPEN (Java)

A minimal example of Java with [Eclipse Jetty] exhibiting a similar issue against servers such as [nginx].

This example has a lower failure rate, but it does test HTTP/2 GET operations directly.

[Eclipse Jetty]: https://eclipse.dev/jetty/documentation/jetty-12/programming-guide/index.html#pg-client-http2
[nginx]: https://nginx.org/

## How-to

### nginx

Grab `nginx` with HTTP/2 support. For example, local build:

```bash
wget -nv https://nginx.org/download/nginx-1.25.2.tar.gz
tar -xf nginx-1.25.2.tar.gz
rm nginx-1.25.2.tar.gz

cd nginx-1.25.2
./configure --prefix=$HOME/.local --with-http_v2_module
make
make install

sed -i -r -e 's/listen[ \t]+80;/listen 8088;\n        http2 on;/' ~/.local/conf/nginx.conf
~/.local/sbin/nginx -g 'daemon off;'
```

### Java

Grab some variant of Java, such as [Eclipse Temurin].

[Eclipse Temurin]: https://adoptium.net/temurin/releases/?os=linux&arch=x64&package=jdk&version=17

### The application

An archive of [Eclipse Jetty Home](https://eclipse.dev/jetty/download.php) will be downloaded and extracted by Make. Java code will be compiled and run.

```bash
make run
```

The expected output:

```text
- nginx (with ES)
   response HTTP/2.0 status 200
- nginx (w/o ES)
   response HTTP/2.0 status 200
- haproxy (with ES)
   response HTTP/2.0 status 200
- haproxy (w/o ES)
   response HTTP/2.0 status 502
```
