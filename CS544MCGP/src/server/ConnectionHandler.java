/** 
 * CS544 Computer Networks
 * create time��2016/5/22
 * group member: 
 *   Kenneth Balogh
 *   Arudra Venkat
 *   Xiaofeng Zhou
 *   Daniel Speichert
 * purpose:
 *  The ConnectionHandler class is manager that deal with all the clients' connections.
 *  Once the handler receive an connection request from the server port, it will create a 
 *  new thread to deal with this specific connection.
 */

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
 * Server Handler will manage all connections between clients and the server
 * 
 */
public class ConnectionHandler extends Thread {

	// the TCP server socket that that designate 9070 as its port
	private ServerSocket serverSocket;
	// read data from the client
	private BufferedReader bufIn;
	// send data to the client
	private BufferedWriter bufOut;
	// socket that will be used to deal with specific connection
	private Socket socket;
	// the ID that combined with the connection, it is used to
	// realize concurrent function
	private int ID;

	/**
	 * constructor of the handler, it will receive the serversocket from the
	 * Server class and uses it to listen on port 9070 for coming connections.
	 * 
	 * @param socket
	 */
	public ConnectionHandler(ServerSocket socket) {
		this.serverSocket = socket;
	}

	public void run() {
		socket = null;
		while (socket == null) {
			try {
				//open a new socket to deal with the coming connection request
				socket = serverSocket.accept();
				System.out.println("New connection accepted "
						+ socket.getInetAddress() + ":" + socket.getPort());
			} catch (IOException e) {
				e.printStackTrace();
			}
		}
		ServiceThread serviceThread = new ServiceThread(socket, ID);
		serviceThread.start();
	}
}

