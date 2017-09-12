package main

import (
	"net"
	"comet/libs"
	log "github.com/alecthomas/log4go"
	"time"
)

func startTcp(bind string)error{
	log.Info("start listen tcp addr:%s",bind)
	go tcpListen(bind)
	return nil
}

func tcpListen(bind string){
	var (
		conn *net.TCPConn
		err error
	)

	addr,err := net.ResolveTCPAddr("tcp",bind)
	if err != nil{
		log.Error("resolveTcpAddr:%s  error:%s",bind,err)
		panic(err)
	}
	l,err := net.ListenTCP("tcp",addr)
	if err != nil{
		log.Error("listen:%s error:%s",bind,err)
		panic(err)
	}
	defer func(){
		log.Info("tcp addr %s close",bind)
		if err := l.Close();err != nil{
			log.Error("listenner %s close err:%s",bind,err)

		}

	}()

	for{
		log.Debug("start accept")
		conn,err = l.AcceptTCP()
		if err != nil{
			log.Error("listenner.AcceptTCP err:%s",err)
			return
		}
		if err = conn.SetReadDeadline(time.Now().Add(time.Second * 60));err != nil{
			log.Error("conn.SetReadDeadline err:%s",err)
			conn.Close()
			return
		}
		if err := conn.SetKeepAlive(false);err != nil{
			log.Error("conn.SetKeepAlive err:%s",err)
			conn.Close()
			return
		}
		if err = conn.SetReadBuffer(1024);err != nil{
			log.Error("conn.SetReadBuffer() error(%v)",err)
			return
		}
		if err = conn.SetWriteBuffer(2048);err != nil{
			log.Error("conn.setWriteBuffer() error(%v)",err)
			return
		}
		go handleTcpConn(conn)
	}
}

func handleTcpConn(conn *net.TCPConn){
	p := &libs.Proto{}
	for{
		if err := p.ReadTcp(conn);err!=nil{
			log.Error("readtcp err:%s",err)
			break
		}
		log.Debug(p.SeqId)
		go write(conn)
	}
	conn.Close()
}

func write(conn *net.TCPConn){
	p := &libs.Proto{}
	p.SeqId = 101;
	p.Operation = 102;
	p.Ver = 103;
	p.Body = []byte("");
	p.WriteTcp(conn)
}

func dispatchTcp(conn *net.TCPConn){

}
