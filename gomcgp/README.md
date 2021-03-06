# Running/Testing

A compiled executable for windows and linux are included.  There are three included *.pem files.
THese files must be in the same directory as the executable when its being run.  

Follow the instructions below to install GO.  

If you want to compile the source follow the rest of the instructions after installing GO.

The below commands along with others are demoed in the videos.

```bash
# should print help:
gomcgp_x64.exe
# should print device list:
gomcgp_x64.exe
# should print authentication failure:
gomcgp_x64.exe -i admin device list
# should turn on thermometer
gomcgp_x64.exe device action --on 6
# should close the garage door:
gomcgp_x64.exe device action --close 1
# should show above changes:
gomcgp_x64.exe device list

```

# How to build after Go is setup (See Setup Below)

```bash
Simple Build
go build -ldflags "-X main.buildtime '`date`'"
Advanced Build Options
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.buildtime '`date`' -X main.buildver '`git log --pretty=format:'%h' -n 1`'" -o gomcgp
GOOS=windows GOARCH=amd64 go build -ldflags "-X main.buildtime '`date`' -X main.buildver '`git log --pretty=format:'%h' -n 1`'" -o gomcgp_x64.exe
GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.buildtime '`date`' -X main.buildver '`git log --pretty=format:'%h' -n 1`'" -o gomcgp_darwin-amd64
```

# Windows Setup to Run/Compile Source

Install [Go](https://golang.org/dl/)

Windows 7

  - Setup GO in windows path
  - From the desktop, right click the Computer icon
  - Choose Properties from the context menu
  - Click the Advanced system settings link
  - Click Environment Variables
  - In the EDIT System Variable window, specify the value in the PATH environment variable
  -     C:\Go\bin
  - Add GOPATH Variable to path
  - In the NEW System Variable window, specify the value of the GOPATH environment variable
  -     C:\GOPAT
- Install [Git](https://help.github.com/desktop/guides/getting-started/installing-github-desktop/)
- Add Git to Path same instructions as above

>Need to locate the Git \cmd directory your computer. Git is typically located here:
-       C:\Users\<user>\AppData\Local\GitHub\PortableGit_<guid>\cmd\git.exe
>On your computer, replace <user> with your user and find out what the <guid> is for your computer.
- The guid may change each time GitHub updates PortableGit, but they're working on a solution to that
- Copy it and paste it into a command prompt (right-click > paste to paste in the terminal) to verify that it works. You should see the Git help response that lists common Git commands.
-  If you see that the system cannot find the path specified. Then the URL isn’t right. Once you have it right, create the link to the directory using this format:
-       C:\Users\<user>\AppData\Local\GitHub\PortableGit_<guid>\cmd
	(Note: \cmd at the end, not \cmd\git.exe anymore!)

- Install And Build at Once - If GOPATH and Git are Setup in path
```bash
	# git clone git@github.com:aaroo/cs544g2.git
	# cd cs544g2/gomcgp
	# go get
	# go build
```
> Optional Install package one at a time
```bash
	# go get github.com/urfave/cli from Command Line
	# go get github.com/olekukonko/tablewriter
	# go build -ldflags "-X main.buildtime '`date`'"
```
This will install the required packages into your GOPATH

### Robustness

We have tested our implementation against basic use cases fuzzing.  The protocol
has not gone through extensive testing to say it is fully robust.  Some basic testing
was done.
- Correctness:
  Tested  functionality of the different components of the code.  The implmentation of
  version checks, autentication and security.
- Protocol correctness:
  In addition we ran tests to check how the protocol is implemented against
  the specification and DFA.
- Robustness:
   A few tests that we ran were sendinng
  - invalid messages
  - changed the user but supplied the wrong certificate
  - sending invalid ports
This was not an exahustive list of testing but showed we could handle basic errors.
- Concurrency:
  We tested that the server could handle multiple connections.
- Environment:
  The server-client was tested on different operating systems including windows and linux.

### Extra Credit
The extra credit was not implmented.

### Version
1.0.1
