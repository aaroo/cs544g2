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
	}

	public static void main(String[] args) {
		try {
			Server server = new Server();
			server.service();
			System.out.println("server started");
		} catch (IOException e) {
			e.printStackTrace();
		}
	}

	/**
	 * set up the server, the method will create a server handler that deal with all clients connection
	 */
	public void service() {
		ServerHandler handler = new ServerHandler(serverSocket);
		Thread t = new Thread(handler);
		t.start();
	}
}
