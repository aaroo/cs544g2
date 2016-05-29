package server;

public class Message {

    String version;

    String op;

    String err;

    Payload payload;

    //IPayloadData payloadData;

  /*  public Message() {
        if(op.equalsIgnoreCase("control")) {
            ObjectMapper mapper = new ObjectMapper();
            try {
                Payload controlPayload = mapper.readValue(payload, Payload.class);
            } catch (IOException e) {
                e.printStackTrace();
            }
        } else if(op.equalsIgnoreCase("control")) {
            ObjectMapper mapper = new ObjectMapper();
            try {
                Payload controlPayload = mapper.readValue(payload, Payload.class);
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
*/
  /*  public class ControlPayload implements  IPayloadData{


    }

    public class StatusPayload implements IPayloadData{

        ArrayList<Device> deviceList = new ArrayList<Device>();

    }
*/
/*

    private interface IPayloadData {
    }
*/

    public String getVersion() {
        return version;
    }

    public void setVersion(String version) {
        this.version = version;
    }

    public String getOp() {
        return op;
    }

    public void setOp(String op) {
        this.op = op;
    }

    public String getErr() {
        return err;
    }

    public void setErr(String err) {
        this.err = err;
    }

    public Payload getPayload() {
        return payload;
    }

    public void setPayload(Payload payload) {
        this.payload = payload;
    }
}
