# Copyright (c) 2023 Valters Jansons

jetty-version = 12.0.1

java-options = \
	-cp '.:$\
	jetty-home-$(jetty-version)/lib/http2/jetty-http2-client-$(jetty-version).jar:$\
	jetty-home-$(jetty-version)/lib/jetty-alpn-client-$(jetty-version).jar:$\
	jetty-home-$(jetty-version)/lib/jetty-http-$(jetty-version).jar:$\
	jetty-home-$(jetty-version)/lib/http2/jetty-http2-common-$(jetty-version).jar:$\
	jetty-home-$(jetty-version)/lib/http2/jetty-http2-hpack-$(jetty-version).jar:$\
	jetty-home-$(jetty-version)/lib/jetty-io-$(jetty-version).jar:$\
	jetty-home-$(jetty-version)/lib/jetty-jmx-$(jetty-version).jar:$\
	jetty-home-$(jetty-version)/lib/jetty-util-$(jetty-version).jar:$\
	jetty-home-$(jetty-version)/lib/slf4j-api-2.0.7.jar'

java-options = \
	-cp .:jetty-home-$(jetty-version)/lib/*:jetty-home-$(jetty-version)/lib/http2/*:jetty-home-$(jetty-version)/lib/logging/*

.PHONY: all build clean run

all: build
build: Test.class

clean:
	rm -f *.class jetty-home-$(jetty-version).tar.gz
	rm -rf jetty-home-$(jetty-version)
distclean: clean

jetty-home-$(jetty-version).tar.gz:
	wget -nv --no-use-server-timestamps -O $@ https://repo1.maven.org/maven2/org/eclipse/jetty/jetty-home/$(jetty-version)/jetty-home-$(jetty-version).tar.gz

jetty-home-$(jetty-version)/: jetty-home-$(jetty-version).tar.gz
	tar -xmf jetty-home-$(jetty-version).tar.gz

jetty-home-$(jetty-version)/lib/http2/jetty-http2-client-$(jetty-version).jar: jetty-home-$(jetty-version)/
	wget -nv --no-use-server-timestamps -O $@ https://repo1.maven.org/maven2/org/eclipse/jetty/http2/jetty-http2-client/$(jetty-version)/jetty-http2-client-$(jetty-version).jar

%.class: %.java jetty-home-$(jetty-version)/ jetty-home-$(jetty-version)/lib/http2/jetty-http2-client-$(jetty-version).jar
	javac $(java-options) $(basename $@).java

run: build
	java $(java-options) Test
