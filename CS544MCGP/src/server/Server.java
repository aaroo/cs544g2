package server;

import java.io.IOException;
import java.net.ServerSocket;
import java.net.Socket;
/**
 * Server 
 *
 */
public class Server {
	private ServerSocket serverSocket;
	
	public static final int VERSION = 1;
	
	public static final int PORT = 9070;

	public Server() throws IOException {
		serverSocket = new ServerSocket(Server.PORT);
		System.out.println("server started");
	}

	public static void main(String[] args) {
		try {
			Server server = new Server();
			server.service();
		} catch (IOException e) {
			e.printStackTrace();
		}
	}

	public void service() {
		while (true) {
			Socket socket = null;
			try {
				socket = serverSocket.accept();
				System.out.println("New connection accepted " + socket.getInetAddress() + ":" + socket.getPort());
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
