package client;

import java.net.Socket;
/**
 * client used to do the test
 *
 */
public class Client {

	public static void main(String[] args) throws Exception{
		final int length = 10;

		Socket[] sockets = new Socket[length];
		
		for (int i = 0; i < sockets.length; i++) {
			sockets[i] = new Socket("localhost", 9070);
			System.out.println("No."+(i+1)+"connection succeed.");
		}
		Thread.sleep(3000);
		for (int i = 0; i < sockets.length; i++) {
			sockets[i].close();
		}
	}
}
