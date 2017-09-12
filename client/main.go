package main

import (
	"net"
	"fmt"
)

func main(){
	conn,err := net.Dial("tcp","localhost:5333")
	if err != nil{
		fmt.Println(err)
		return
	}

	p := &Proto{}
	p.SeqId = 1001;
	p.WriteTcp(conn)
	select {

	}
}
