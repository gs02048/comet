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
	select {

	}
}
