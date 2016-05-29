package server;

/**
 * Created by avenkat000 on 5/29/16.
 */
public class Device {

    String deviceName;

    String status;

    String param;

    public Device(String deviceName, String stats) {
        this.deviceName = deviceName;
        this.status = stats;
    }
    public Device(String deviceName, String stats, String param) {
        this.deviceName = deviceName;
        this.status = stats;
        this.param = param;
    }

    public String getDeviceName() {
        return deviceName;
    }

    public void setDeviceName(String deviceName) {
        this.deviceName = deviceName;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public String getParam() {
        return param;
    }

    public void setParam(String param) {
        this.param = param;
    }
}
