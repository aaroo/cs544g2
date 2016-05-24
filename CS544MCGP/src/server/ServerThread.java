package server;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.net.Socket;

import utils.Utils;

/**
 * thread for a specific client connection
 * 
 */
public class ServerThread extends Thread {

	private Socket socket;
	private BufferedReader bufr;
	private BufferedWriter bufw;

	public ServerThread(Socket socket) {
		this.socket = socket;
	}

	public void run() {
		try {
			bufr = new BufferedReader(new InputStreamReader(
					socket.getInputStream()));
			bufw = new BufferedWriter(new OutputStreamWriter(
					socket.getOutputStream()));

			while (true) {
				String line = null;

				while (line == null) {
					line = bufr.readLine();
					byte[] cmd = Utils.toByteStream(line);
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
		}
	}
}
