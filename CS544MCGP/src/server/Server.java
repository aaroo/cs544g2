/** 
* CS544 Computer Networks
* create time£º2016/5/22
* group member: 
*   Kenneth Balogh
*   Arudra Venkat
*   Xiaofeng Zhou
*   Daniel Speichert
* purpose:
*   The server class conatins main method, it is the entry of the server end.
*   After the server started, it will listen on port 9070, and deal with all 
*   connection request comming from that port. The server class will create a new thread
*   that actually deal with all the connections.
*/

package server;

import java.io.IOException;
import java.net.ServerSocket;

import javax.net.ssl.SSLServerSocketFactory;

public class Server {
	//server socket that used for TCP connection
	private ServerSocket serverSocket;
	//server's original version 
	public static final int VERSION = 1;
	//specific port of the protocol
	public static final int PORT = 9070;
	/**
	 * constructor of the Server class, it will designate port 9070
	 * as the service port
	 * @throws IOException
	 */
	public Server() throws IOException {
		serverSocket = new ServerSocket(Server.PORT);
	}
	
	/**
	 * main method to start the server
	 */
	public static void main(String[] args) {
		try {
			Server server = new Server();
			//create a new thread that deal with all coming connections
			server.service();
			System.out.println("server started");
		} catch (IOException e) {
			e.printStackTrace();
		}
	}

	/**
	 * the server will start to provide services, the method will create a connection handler that deal with all clients'connections
	 */
	public void service() {
		ConnectionHandler handler = new ConnectionHandler(serverSocket);
		handler.start();
	}
}
