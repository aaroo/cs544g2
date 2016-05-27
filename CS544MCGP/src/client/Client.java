package client;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.net.Socket;

import javax.net.ssl.SSLSocket;
import javax.net.ssl.SSLSocketFactory;

/**
 * client used to do the test
 * 
 */
public class Client {

	public static void main(String[] args) throws Exception {
		final int length = 10;

		/*
		 * Socket[] sockets = new Socket[length];
		 * 
		 * for (int i = 0; i < sockets.length; i++) { sockets[i] = new
		 * Socket("localhost", 9070); System.out.println("No." + (i + 1) +
		 * "connection succeed.");
		 * 
		 * }
		 * 
		 * Thread.sleep(3000); for (int i = 0; i < sockets.length; i++) {
		 * sockets[i].close(); }
		 */
		SSLSocketFactory sslsocketfactory = (SSLSocketFactory) SSLSocketFactory.getDefault();
		SSLSocket sslsocket = (SSLSocket) sslsocketfactory.createSocket("localhost", 9070);
		// object to read data from keyboard
		BufferedReader bufr = new BufferedReader(new InputStreamReader(
				System.in));

		// destination, write data into outputstream of the socket and send to
		// server
		BufferedWriter bufOut = new BufferedWriter(new OutputStreamWriter(
				sslsocket.getOutputStream()));

		// reader to read data sent by the server
		BufferedReader bufIn = new BufferedReader(new InputStreamReader(
				sslsocket.getInputStream()));

		String line = null;

		while ((line = bufr.readLine()) != null) {
			if ("over".equals(line)) {
				break;
			}

			bufOut.write(line);
			bufOut.newLine();
			bufOut.flush();

		}

		// String str = bufIn.readLine();
		// System.out.println("server:" + str);
		bufr.close();
		sslsocket.close();
	}
}
