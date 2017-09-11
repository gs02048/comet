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
	conn.Write([]byte("hello"))

	p := &Proto{}
	p.ReadTcp(conn)
	fmt.Println(p.Operation)
	fmt.Println(p.SeqId)
	fmt.Println(p.Ver)
	select {

	}
}
