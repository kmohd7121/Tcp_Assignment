package main

import (
	"net"
	"fmt"
	"log"
)
func main(){
	conn,err:=net.Dial("tcp","localhost:3000")
	//string(data)
	if err!=nil {
		log.Fatalln(err)
	}

	_ , err = conn.Write([]byte("hello server"))
	if err!=nil {
		log.Fatalln(err)
	}
	fmt.Println("Message sent: hello server")

	_,err = conn.Write([]byte("how are you"))
	if err!=nil {
		log.Fatalln(err)
	}
	fmt.Println("Message sent: how are you")

	for {
		buffer:=make([]byte,1400)
		dataSize,err:=conn.Read(buffer)
		if err!=nil{
			fmt.Println("connection close")
			return
		}

		data:=buffer[:dataSize]
		fmt.Println("received message:",string(data))
	}
}