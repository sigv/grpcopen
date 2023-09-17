// Copyright (c) 2023 Valters Jansons

import java.net.InetSocketAddress;
import java.net.SocketAddress;
import java.nio.ByteBuffer;
import java.nio.CharBuffer;
import java.nio.charset.StandardCharsets;
import java.util.concurrent.CompletableFuture;

import org.eclipse.jetty.http.HttpFields;
import org.eclipse.jetty.http.HttpURI;
import org.eclipse.jetty.http.HttpVersion;
import org.eclipse.jetty.http.MetaData;
import org.eclipse.jetty.http2.api.Session;
import org.eclipse.jetty.http2.api.Stream;
import org.eclipse.jetty.http2.client.HTTP2Client;
import org.eclipse.jetty.http2.frames.HeadersFrame;

public class Test {
	private static final String HOSTNAME = "localhost";

	private static final int HAPROXY_PORT = 8080;

	private static final int NGINX_PORT = 8088;

	public static void get(String hostname, int port, boolean headersEndStream) throws Exception {
		// create default Jetty HTTP/2 client
		HTTP2Client http2Client = new HTTP2Client();
		http2Client.start();

		try {
			// connect to the remote
			SocketAddress serverAddress = new InetSocketAddress(hostname, port);
			CompletableFuture<Session> sessionCF = http2Client.connect(serverAddress, new Session.Listener() {
			});
			Session session = sessionCF.get();

			// build request HEADERS frame
			HttpFields requestHeaders = HttpFields.build();
			MetaData.Request request = new MetaData.Request("GET", HttpURI.from("http://example.com/"),
					HttpVersion.HTTP_2, requestHeaders);
			HeadersFrame headersFrame = new HeadersFrame(request, null, headersEndStream);

			// handler logic for HEADERS and DATA. generally from Eclipse documentation:
			// https://eclipse.dev/jetty/documentation/jetty-12/programming-guide/index.html#pg-client-http2-response
			Stream.Listener responseListener = new Stream.Listener() {
				@Override
				public void onHeaders(Stream stream, HeadersFrame frame) {
					try {
						// Delay processing of data, in turn delaying the stream close.
						Thread.sleep(2000);
					} catch (InterruptedException e) {
						throw new RuntimeException(e);
					}

					MetaData metaData = frame.getMetaData();

					if (metaData.isResponse()) {
						MetaData.Response response = (MetaData.Response) metaData;
						System.out.printf("   response HTTP/2.0 status %d\n", response.getStatus());

						if (!frame.isEndStream()) {
							stream.demand();
						}
					}
				}

				@Override
				public void onDataAvailable(Stream stream) {
					Stream.Data data = stream.readData();

					if (data == null) {
						stream.demand();
						return;
					}

					data.release();
					if (!data.frame().isEndStream()) {
						stream.demand();
					}
				}
			};

			// create the stream and send HEADERS
			CompletableFuture<Stream> streamCF = session.newStream(headersFrame, responseListener);
			Stream stream = streamCF.get();

			long startTime = System.nanoTime();
			while (!stream.isClosed()) {
				// 10 seconds timeout, just in case something has hung.
				if (System.nanoTime() - startTime > 10 * 1000000000L) {
					Thread.sleep(10);
				}
			}
		} finally {
			http2Client.stop();
		}
	}

	public static void main(String[] args) throws Exception {
		System.out.println("- nginx (with ES)");
		get(HOSTNAME, NGINX_PORT, true);

		System.out.println("- nginx (w/o ES)");
		get(HOSTNAME, NGINX_PORT, false);

		System.out.println("- haproxy (with ES)");
		get(HOSTNAME, HAPROXY_PORT, true);

		System.out.println("- haproxy (w/o ES)");
		get(HOSTNAME, HAPROXY_PORT, false);
	}
}
