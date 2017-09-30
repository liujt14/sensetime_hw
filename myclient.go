package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:2121")
	if err != nil {
		os.Exit(-1)
	}
	localAddr := conn.LocalAddr().String()
	fmt.Println("localAddress is: ", localAddr)
	fmt.Println("remoteAddress is: ", conn.RemoteAddr())
	for {
		cmdReader := bufio.NewReader(os.Stdin)
		//listen
		messageReader := bufio.NewReader(conn)
		messagefromserver, err := messageReader.ReadString('\n')
		if err != nil {
			os.Exit(-1)
		}
		fmt.Print("Message from server: ", messagefromserver)
		fmt.Print("Enter username: ")
		username, _ := cmdReader.ReadString('\n')
		fmt.Fprintf(conn, username+"\n")
		fmt.Print("Enter Pass: ")
		password, _ := cmdReader.ReadString('\n')
		fmt.Fprintf(conn, password+"\n")
		loginReader := bufio.NewReader(conn)
		loginmessage, _ := loginReader.ReadString('\n')
		fmt.Println(loginmessage)
		//enter cmd
		fmt.Print("Enter cmd: ")
		cmdtosend, err := cmdReader.ReadString('\n')
		if err != nil {
			os.Exit(-1)
		}
		fmt.Fprintf(conn, cmdtosend+"\n")
		//send to server
		//two methods  to send cmd to server  conn.Write([]byte(cmd"\n"))&blow
		switch cmdtosend {
		case "list\n":
			fmt.Println("list from server:")
			listReader := bufio.NewReader(conn)
			listmessage, _ := listReader.ReadString('\n')
			fmt.Println(listmessage)

		case "cwd\n":
			fmt.Println("Enter the aim Dir: ")
			aimDir, _ := cmdReader.ReadString('\n')
			fmt.Fprintf(conn, aimDir+"\n")
		case "pwd\n":
			fmt.Println("CurrentworkDir in server: ")
			dirReader := bufio.NewReader(conn)
			dir, _ := dirReader.ReadString('\n')
			fmt.Println(dir)
		case "retr\n":
			filelistReader := bufio.NewReader(conn)
			files, _ := filelistReader.ReadString('\n')
			fmt.Println(files)
			fmt.Println("Select the file to copy: ")
			aimfile, _ := cmdReader.ReadString('\n')
			fmt.Fprintf(conn, aimfile+"\n")
			//file content Note that filecontent contains lots of \n,we need to received bytes + "\n" then convent it
			data := make([]byte, 100)
			count, _ := conn.Read(data)
			filecontent := string(data[:count])
			fmt.Print("The filecontents: ", filecontent)
			//filecontentReader :=bufio.NewReader(conn)
			//filecontent,_:= filecontentReader.ReadString('\n')
			//fmt.Print(filecontent)
			//save to local?
		case "stor\n":
			fmt.Println("Please enter the file name and contents with \\n as interruption:")
			fmt.Print("Enter the file name:")
			filename, _ := cmdReader.ReadString('\n')
			fmt.Fprintf(conn, filename+"\n")
			fmt.Print("Enter the file contents:")
			filecontents, _ := cmdReader.ReadString('\n')
			fmt.Fprintf(conn, filecontents+"\n")
		}
	}
}
