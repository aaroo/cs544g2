# How to build after GO is setup (See Setup Below)

```bash\windows
go build -ldflags "-X main.buildtime '`date`'"
```
Windows Setup to Compile Source

1) Install Go https://golang.org/dl/
2) Setup GO in path
	Windows 7
		From the desktop, right click the Computer icon.
		Choose Properties from the context menu.
		Click the Advanced system settings link.
		Click Environment Variables. ...
		In the EDIT System Variable (or New System Variable) window, specify the value of the PATH environment variable.
			-example add C:\Go\bin
3)Add GOPATH Variable
	Windows 7
		From the desktop, right click the Computer icon.
		Choose Properties from the context menu.
		Click the Advanced system settings link.
		Click Environment Variables. ...
		In the NEW System Variable (or New System Variable) window, specify the value of the PATH environment variable.
			-example add C:\GOPAT
4)Setup GIT
5)Add GIT to Command Line
	Get the Git URL
		-need to get the url of the Git \cmd directory your computer. Git is located here:
			C:\Users\<user>\AppData\Local\GitHub\PortableGit_<guid>\cmd\git.exe
			On your computer, replace <user> with your user and find out what the <guid> is for your computer. (The guid may change each time GitHub updates PortableGit, but they're working on a solution to that.)
	Copy it and paste it into a command prompt (right-click > paste to paste in the terminal) to verify that it works. You should see the Git help response that lists common Git commands. 
	If you see The system cannot find the path specified. Then the URL isn’t right. Once you have it right, create the link to the directory using this format: C:\Users\<user>\AppData\Local\GitHub\PortableGit_<guid>\cmd
	(Note: \cmd at the end, not \cmd\git.exe anymore!)
	
6)Install CLI Package 
	- Run go get github.com/urfave/cli from Command Line
	- WIll Add files to directory specified in GOPATH

How to Run Executable

1) After Running Build Command Above should get gomcgp.exe
2) Run gomcgp.exe with proper command line arguement 
3) To Start Server
	To Run as local host type: gomcgp.exe server 
	To Setup IP Address add ipadress and port after server: gomcgp.exe server 192.168.1.140:6666
4) To Get Devices
	-type gomcgp.exe device retrieve filename (filename is location of devices files to be stored
		-example type: gomcgp device retrieve c:\devices.txt
5) To Update Devices
		-type gomcgp.exe device update