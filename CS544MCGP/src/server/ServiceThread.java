/** 
 * CS544 Computer Networks
 * create time£º2016/5/22
 * group member: 
 *   Kenneth Balogh
 *   Arudra Venkat
 *   Xiaofeng Zhou
 *   Daniel Speichert
 * purpose:
 *    The ServiceThread class is created when the handler received an request of connection 
 *    from the client end, the handler will create a new serviceThread to deal with the specific 
 *    connection, the servicethread will receive data sent by the client and send data to the client 
 *    through byte stream.
 */
package server;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.net.Socket;

import utils.Utils;

public class ServiceThread extends Thread {

	private Socket socket;
	private BufferedReader bufIn;
	private BufferedWriter bufOut;
	private int ID;

	public ServiceThread(Socket socket, int ID) {
		this.socket = socket;
		this.ID = ID;
	}

	public void run() {

		try {
			// using a bufferedReader to store data received by the socket
			bufIn = new BufferedReader(new InputStreamReader(
					socket.getInputStream()));
			// using a bufferedWriter to store data that will be sent to the
			// client
			// through the socket
			bufOut = new BufferedWriter(new OutputStreamWriter(
					socket.getOutputStream()));

			while (true) {
				String line = null;
				while ((line = bufIn.readLine()) != null) {
					System.out.println("data received: " + line);
					// convert the received data into byte array for command
					// usage.
					byte[] cmd = Utils.toByteStream(line);
					// System.out.println("data reveived in byte:" + cmd);
				}
			}
		} catch (IOException e) {
			e.printStackTrace();
		} finally {
			try {
				if (bufIn != null)
					bufIn.close();
			} catch (IOException e) {
				throw new RuntimeException("read close failed");
			}
			try {
				if (bufOut != null)
					bufOut.close();
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

	/**
	 * method to disconnect
	 */
	private void shutdown() {
		try {
			socket.close();
		} catch (IOException e) {
			throw new RuntimeException("connection close failed");
		}
	}
}
