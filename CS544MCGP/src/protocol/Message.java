package protocol;
/**
 * definition of the PDU 
 *
 */
public class Message {
	//message codes
	public static final byte  MSG_REQUEST_CONN = 0;
	public static final byte  MSG_ERROR = 1;
	public static final byte  MSG_CONN_READY = 2;
	public static final byte  MSG_READ_STATUS = 3;
	public static final byte  MSG_CONTROL_DEVICES = 4;
	public static final byte  MSG_DONE = 5;
	public static final byte  MSG_CONFIRM = 6;
	public static final byte  MSG_VERSION = 7;
	public static final byte  MSG_AUTHEN = 8;
	public static final byte  MSG_SERVICE_DONE = 9;
	public static final byte  MSG_NUM_DEVICES = 10;
	//error codes
	public static final byte  ERR_CONN_REFUSED = 0;
	public static final byte  ERR_AUTH = 1;
	public static final byte  ERR_SERV = 2;
	public static final byte  ERR_CEHCK_STATUS = 3;
	public static final byte  ERR_SERV_CTRL = 4;
	public static final byte  ERR_CLIENT_CMM = 5;
	public static final byte  ERR_SERV_EXE = 6;
	
	private  byte[] msgs = new byte[18];
	
	/**
	 * constructor of control message
	 * @param msg
	 */
	public Message(byte[] msgs) {
		this.msgs = msgs;
	}
	
	
	
		
}
