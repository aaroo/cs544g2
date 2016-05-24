package server;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.net.ServerSocket;
import java.net.Socket;

import utils.Utils;

/**
 * thread for a specific client connection
 * 
 */
public class ServerThread extends Thread {

	private ServerSocket serversocket;
	private BufferedReader bufr;
	private BufferedWriter bufw;

	public ServerThread(ServerSocket serversocket) {
		this.serversocket = serversocket;
	}

	public void run() {
		Socket socket = null;
		while (true) {
			try {
				socket = serversocket.accept();
				System.out.println("New connection accepted "
						+ socket.getInetAddress() + ":" + socket.getPort());
				bufr = new BufferedReader(new InputStreamReader(
						socket.getInputStream()));
				bufw = new BufferedWriter(new OutputStreamWriter(
						socket.getOutputStream()));

				while (true) {
					String line = null;

					while ((line = bufr.readLine()) != null) {
						System.out.println("data received: " + line);
						byte[] cmd = Utils.toByteStream(line);
						System.out.println("data reveived in byte:" + cmd);
					}
				}
			} catch (IOException e) {
				e.printStackTrace();
			} finally {
				try {
					if (bufr != null)
						bufr.close();
				} catch (IOException e) {
					throw new RuntimeException("read close failed");
				}
				try {
					if (bufw != null)
						bufw.close();
				} catch (IOException e) {
					throw new RuntimeException("write close failed");
				}
				if (socket != null) {
					try {
						socket.close();
					} catch (IOException e) {
						e.printStackTrace();
					}
				}
			}

		}
	}
}
