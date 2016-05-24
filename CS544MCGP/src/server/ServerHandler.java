package server;

import java.io.IOException;
import java.net.ServerSocket;
import java.net.Socket;

/**
 * Server Handler will manage all connections between clients and the server
 * 
 */
public class ServerHandler implements Runnable {

	private ServerSocket serverSocket;

	public ServerHandler(ServerSocket socket) {
		this.serverSocket = socket;
	}

	public void run() {
		while (true) {
			Socket socket = null;
			try {
				socket = serverSocket.accept();
//				System.out.println("New connection accepted "
//						+ socket.getInetAddress() + ":" + socket.getPort());
				ServerThread thread = new ServerThread(socket);
				thread.start();
			} catch (IOException e) {
				e.printStackTrace();
			} finally {
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
