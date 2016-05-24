package utils;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;

public class Utils {
	public static String getTextFromStream(InputStream is) {

		byte[] b = new byte[1024];
		int len = 0;
		//create byteArrayOutputStream
		//read data from inputstream and store it in bos
		ByteArrayOutputStream bos = new ByteArrayOutputStream();
		try {
			while ((len = is.read(b)) != -1) {
				bos.write(b, 0, len);
			}
			// turn data in bos to string
			String text = new String(bos.toByteArray());
			return text;
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		finally {
			if (bos != null) {
				try {
					bos.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}
		}
		return null;
	}
	
	public static byte[] toByteStream(String hexStr) {
		String[] bytes = hexStr.split(" ");
		byte[] res = new byte[bytes.length];
		for (int i = 0; i < bytes.length; i++)
			res[i] = (byte) Integer.parseInt(bytes[i], 16);
		return res;
	}
}
