package main

import (
	"net"
	"fmt"
	"comet/libs"
	"encoding/json"
)

type Auth struct{
	UserId int64 `json:"uid"`
	RoomId string `json:"rid"`
	Token string `json:"token"`
}

func main(){
	addr,_ := net.ResolveTCPAddr("tcp","localhost:5333")
	conn,err := net.DialTCP("tcp",nil,addr)
	if err != nil{
		fmt.Println(err)
		return
	}
	a := &Auth{UserId:100,RoomId:"123",Token:""}
	body,_ := json.Marshal(a)
	p := &libs.Proto{}
	p.Operation = 1;
	p.Body = body

	p.WriteTcp(conn)

	for{
		if err:=p.ReadTcp(conn);err != nil{
			fmt.Println(err)
			break
		}
		fmt.Println(p.SeqId)
	}
	conn.Close()

}
