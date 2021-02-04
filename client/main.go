package main

import (
	"net"
	"fmt"
)
func main(){
	conn,err:=net.Dial("tcp","localhost:3000")s
	if err!=nil {
		fmt.Println(err)
	}

	_ , err = conn.Write([]byte("hello server"))
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("Message sent: hello server")

	_,err = conn.Write([]byte("how are you"))
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("Message sent: how are you")s

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