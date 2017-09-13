package main

import (
	"container/list"
	//"net"
	//"time"
	//"comet/libs"
	//log "github.com/alecthomas/log4go"
	"net"
	"comet/define"
)




type RoomBucket struct {
	bucket map[string]*list.List
	pusharg chan *define.BoardcastRoomArg
}

func NewBucket() *RoomBucket{
	c := make(chan *define.BoardcastRoomArg,100)
	bucket := &RoomBucket{
		bucket:make(map[string]*list.List,1000),
		pusharg:c,
	}
	go bucket.roomproc(c)
	return bucket
}

func (b *RoomBucket) Add(rid string,conn *net.TCPConn)(*list.Element,error){
	if _,ok := b.bucket[rid];!ok{
		b.bucket[rid] = list.New()
	}
	elem := b.bucket[rid].PushFront(conn)
	return elem,nil
}

func (b *RoomBucket) Delete(rid string,element *list.Element)(error){
	b.bucket[rid].Remove(element)
	return nil
}

func (b *RoomBucket) roomproc(c chan *define.BoardcastRoomArg){
	for{
		arg := <-c
		if _,ok := b.bucket[arg.Rid];ok{
			for elem := b.bucket[arg.Rid].Front();elem != nil;elem = elem.Next(){
				conn := elem.Value.(*net.TCPConn)
				arg.P.WriteTcp(conn)
			}
		}
	}
}
