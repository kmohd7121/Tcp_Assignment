package main

import(
	"fmt"
	"net"
)

func main(){
	fmt.Println("server listening on 3000")
	listener,err:=net.Listen("tcp","localhost:3000")
	if err!=nil {
		fmt.Println(err)
	}
	defer listener.Close()

	for {
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println("new connection")
		go listenConnection(conn)
	}
}

func listenConnection(conn net.Conn) {
	for{
		buffer:=make([]byte,1400)
		dataSize,err:=conn.Read(buffer)
		if err!=nil{
			fmt.Println("connection close")
			return
		}

		data:=buffer[:dataSize]
		fmt.Println("received message:",string(data))

		_,err=conn.Write(data)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println("Message sent:",string(data))
	}
}