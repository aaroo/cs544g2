package server;

import utils.Utils;

import javax.net.ssl.SSLServerSocket;
import javax.net.ssl.SSLServerSocketFactory;
import javax.net.ssl.SSLSocket;
import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.net.ServerSocket;

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
		//Socket socket = null;
        SSLServerSocket sslserversocket = null;
        SSLServerSocketFactory sslserversocketfactory =
                (SSLServerSocketFactory) SSLServerSocketFactory.getDefault();
        try {
             sslserversocket = (SSLServerSocket) sslserversocketfactory.createServerSocket(9999);
        } catch (IOException e) {
            e.printStackTrace();
        }

        while (true) {
			try {
				//socket = serversocket.accept();
                SSLSocket sslsocket = (SSLSocket) sslserversocket.accept();

                System.out.println("New connection accepted "
						+ sslserversocket.getInetAddress() + ":9999");
				bufr = new BufferedReader(new InputStreamReader(
						sslsocket.getInputStream()));
				bufw = new BufferedWriter(new OutputStreamWriter(
						sslsocket.getOutputStream()));

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
				if (sslserversocket != null) {
					try {
						sslserversocket.close();
					} catch (IOException e) {
						e.printStackTrace();
					}
				}
			}

		}
	}
}
