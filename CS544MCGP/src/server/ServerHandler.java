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

		ServerThread thread = new ServerThread(serverSocket);
		thread.start();
	}

}
