/** 
 * CS544 Computer Networks
 * create time��2016/5/22
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

import org.codehaus.jackson.JsonGenerator;
import org.codehaus.jackson.JsonParser;
import org.codehaus.jackson.map.ObjectMapper;

import javax.net.ssl.SSLSocket;
import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;

public class ServiceThread extends Thread {

	private SSLSocket sslsocket;
	private BufferedReader bufIn;
	private BufferedWriter bufOut;
	private int ID;

	public ServiceThread(SSLSocket sslsocket, int ID) {
		this.sslsocket = sslsocket;
		this.ID = ID;
	}

	public void run() {

		try {
			// using a bufferedReader to store data received by the socket
			bufIn = new BufferedReader(new InputStreamReader(
					sslsocket.getInputStream()));
            while(true) {
                try {
                    ObjectMapper mapper = new ObjectMapper();
                    mapper.configure(JsonParser.Feature.AUTO_CLOSE_SOURCE, false);
                    mapper.configure(JsonGenerator.Feature.AUTO_CLOSE_TARGET, false);
                    Door door = mapper.readValue(bufIn, Door.class);
                    System.out.println("Door name = " + door.getName());
                    System.out.println("Door status = " + door.getStatus());

                    //Some class methods to do message handling

                    Door output = MCGPMessageHandler.getOutputForInput(door);



                    bufOut = new BufferedWriter(new OutputStreamWriter(
                            sslsocket.getOutputStream()));

                    mapper.writeValue(bufOut, door);
                    bufOut.flush();


                } catch (Exception e) {
                    e.printStackTrace();
                }
            }



			// using a bufferedWriter to store data that will be sent to the
			// client
			// through the socket


			/*while (true) {
				String line = null;
				while ((line = bufIn.readLine()) != null) {
					System.out.println("data received: " + line);
					// convert the received data into byte array for command
					// usage.
					byte[] cmd = Utils.toByteStream(line);
					// System.out.println("data reveived in byte:" + cmd);

				}
			}*/
		} catch (Exception e) {
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
			if (sslsocket != null) {
				try {
					sslsocket.close();
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
			sslsocket.close();
		} catch (IOException e) {
			throw new RuntimeException("connection close failed");
		}
	}
}
