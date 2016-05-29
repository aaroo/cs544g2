package server;

import java.util.ArrayList;

/**
 * Created by avenkat000 on 5/29/16.
 */
class Payload {
    ArrayList<Device> deviceList = new ArrayList<Device>();

    public void addDevice(Device device) {
        deviceList.add(device);
    }

    public ArrayList<Device> getDeviceList() {
        return deviceList;
    }

    public void setDeviceList(ArrayList<Device> deviceList) {
        this.deviceList = deviceList;
    }
}
