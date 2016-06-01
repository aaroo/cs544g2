# Manual testing

```bash
# should print help:
gomcgp
# should print device list:
gomcgp device list
# should print authentication failure:
gomcgp -i admin device list
# should turn on thermometer
gomcgp action --on 6
# should close the garage door:
gomcgp action --close 1
# should show above changes:
gomcgp device list

```

# How to build after GO is setup (See Setup Below)

```bash\windows
go build -ldflags "-X main.buildtime '`date`'"
```
# Windows Setup to Compile Source

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

### Version
1.0.0