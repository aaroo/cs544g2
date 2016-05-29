package server;

/**
 * Created by avenkat000 on 5/29/16.
 */
public class MCGPMessageHandler {



    public static Message getOutputForInput(Message message) {
       if(message.getOp().equalsIgnoreCase("status")) {
            getStatus(message);
       } else if(message.getOp().equalsIgnoreCase("control")) {
           executeControl(message);
       }
        return null;

    }

    private static Message getStatus(Message message) {
        Message output = getStatus(message);
        output.setVersion(message.getVersion());
        output.setOp("confirm");
        output.setErr("");
        Payload outputPayload = new Payload();
        Device device1 = new Device("device1", "open");
        outputPayload.addDevice(device1);
        Device device2 = new Device("device1", "open");
        output.setPayload(outputPayload);
        return output;
    }

    private static Message executeControl(Message message) {
        return null;
    }
}
