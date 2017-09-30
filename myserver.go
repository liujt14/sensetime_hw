package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"os"
	"syscall"
)

func main() {
	ln, err := net.Listen("tcp", ":2121")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	//wait for a connection.
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	for {
		Addr2 := conn.LocalAddr()
		remoteAddr := conn.RemoteAddr()
		fmt.Println("connection build successfully!\nlocal address:", Addr2, "\nclientAddr: ", remoteAddr)
		fmt.Println("Waiting for login......")
		//cmdReader in server
		conn.Write([]byte("please login!" + "\n"))
		//store the log info
		fmt.Print("loginfo:\n")
		userReader := bufio.NewReader(conn)
		Readusername, err := userReader.ReadString('\n')
		fmt.Print("Username:", Readusername)
		passReader := bufio.NewReader(conn)
		Readpass, _ := passReader.ReadString('\n')
		fmt.Print("Password:", Readpass)
		//tell client success
		fmt.Fprint(conn, "login allowed!"+"\n")
		fmt.Println("Waiting for cmd...")
		cmdReader := bufio.NewReader(conn)
		cmdreadfromclient, err := cmdReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		//cmd received
		fmt.Print("cmd received from client: ", cmdreadfromclient)
		//cmd action
		cmdfromclient := strings.Split(cmdreadfromclient, "\n")[0]
		switch cmdfromclient {
		case "list":
			fmt.Println("received cmd :list")
			cmd := exec.Command("ls")
			currentfilesanddirs, _ := cmd.Output()
			thefilesandfolders := string(currentfilesanddirs[:])
			fmtstringslicefiles := strings.Split(thefilesandfolders, "\n")
			fmtstringfiles := strings.Join(fmtstringslicefiles, "  ")
			fmt.Println(fmtstringfiles)
			conn.Write([]byte(fmtstringfiles + "\n"))
			cmd.Dir = "/home/sensetime/go/bin"
			fmt.Println("current dir: ", cmd.Dir)
		case "cwd":
			fmt.Println("received cmd :cwd")
			aimdirbyte := make([]byte, 20)
			n, _ := conn.Read(aimdirbyte)
			//aimDir,_ := bufio.NewReader(conn).ReadString('\n')
			aimDir := string(aimdirbyte[:n-2])
			fmt.Printf("Received aimDir is :%s\n", aimDir)
			err := syscall.Chdir(aimDir)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Changed to the DIR:%s\n", aimDir)
			//cmd := exec.Command("cd %s",aimDir)
			//suffer to supply cd dir?
			//test
			//cmd := exec.Command("cd",cmd.Dir)
		case "pwd":
			fmt.Println("received cmd :pwd")
			cmd := exec.Command("pwd")
			currentDir, _ := cmd.Output()
			currentworkdir := string(currentDir[:])
			fmt.Print("The current work Dir: ", currentworkdir)
			//dirsplit := strings.Split(currentworkdir,"\n")
			conn.Write([]byte(currentworkdir + "\n"))

			////test make
			//aimfilename := make([]byte,40)
			//n,_ := conn.Read(aimfilename)
			//fmt.Print(n,aimfilename)
		case "retr":
			fmt.Println("received cmd :retr")
			cmd := exec.Command("ls")
			currentfilesanddirs, _ := cmd.Output()
			thefilesandfolders := string(currentfilesanddirs[:])
			fmtstringslicefiles := strings.Split(thefilesandfolders, "\n")
			fmtstringfiles := strings.Join(fmtstringslicefiles, "  ")
			fmt.Println(fmtstringfiles)
			conn.Write([]byte(fmtstringfiles + "\n"))
			//receive the aimfilename
			aimfilename := make([]byte, 50)
			n, _ := conn.Read(aimfilename)
			aimfile := string(aimfilename[:n-2])
			//test print aimfile
			fmt.Printf("Opened file is :%s\n", aimfile)

			file, err := os.Open(aimfile)
			//file,err:=os.Open("test.txt")
			if err != nil {
				log.Fatal(err)
			}
			data := make([]byte, 100)
			count, _ := file.Read(data)
			datastring := string(data[:count])
			fmt.Printf("Read %d bytes:%q\n", count, data[:count])
			conn.Write([]byte(datastring))
		case "stor":
			fmt.Println("received cmd :stor")
			filenametomake := make([]byte, 50)
			filecontents := make([]byte, 200)
			namen, _ := conn.Read(filenametomake)
			contentsn, _ := conn.Read(filecontents)
			newfilename := string(filenametomake[:namen-2])
			newfile, _ := os.Create(newfilename)
			filecontentsString := string(filecontents[:contentsn-2])
			newfile.WriteString(filecontentsString)
			fmt.Println("Strings read:", filecontentsString)
			newfile.Sync()
			newfile.Close()
			cmd := exec.Command("ls")
			currentfilesanddirs, _ := cmd.Output()
			thefilesandfolders := string(currentfilesanddirs[:])
			fmtstringslicefiles := strings.Split(thefilesandfolders, "\n")
			fmtstringfiles := strings.Join(fmtstringslicefiles, "  ")
			fmt.Println(fmtstringfiles)
			fmt.Printf("%s Built successfully!\n", newfilename)

		default:
			fmt.Println("The cmd hasn't been built,Please enter the simple cmd!")
		}
	}
}
