package main

import (
	"fmt"
	"bufio"
	"net"
	"os"
	"strings"
	"strconv"
)
var count = 0
func main(){
	Argument:=os.Args
	if len(Argument)==1{
		fmt.Println("Please Provide port number")
		return
	}
	Port:= ":"+ Argument[1]
	conn,err:=net.Listen("tcp",Port)
	if(err!=nil){
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		data,err :=conn.Accept()
		if err!=nil{
			fmt.Println(err)
		}

		count++

		go handleConnection(data,count)
	}
}

func handleConnection(conn net.Conn, count int){
	for{
		netData, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		counter := strconv.Itoa(count)

		fmt.Print("CLIENT SAYS: ", counter, ": ", string(netData))

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("SERVER: ", counter, ": ")
		text, _ := reader.ReadString('\n')
		//user input is sent to the TCP server over the network
		fmt.Fprintf(conn, text+"\n")

	}
	conn.Close()
}