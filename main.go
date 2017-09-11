package main

import (
	log "github.com/alecthomas/log4go"
	"net"
)
func main(){
	startTcp("localhost:5333")
	signalCH := InitSignal()
	HandleSignal(signalCH)
	// exit
	log.Info("comet stop")
}

func startTcp(bind string)error{
	log.Info("start listen tcp addr:%s",bind)
	go tcpListen(bind)
	return nil
}

func tcpListen(bind string){
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
		conn,err := l.AcceptTCP()
		if err != nil{
			log.Error("listenner.AcceptTCP err:%s",err)
			continue
		}
		if err := conn.SetKeepAlive(false);err != nil{
			log.Error("conn.SetKeepAlive err:%s",err)
			conn.Close()
			continue
		}
		go handleTcpConn(conn)
	}
	select {

	}
}

func handleTcpConn(conn net.Conn){
	log.Info("hello")
	for{
		var msg [100]byte
		n,err := conn.Read(msg[0:])
		if err != nil{
			log.Info("conn.Read err:%s",err)
			break
		}
		log.Debug("msg:%s",string(msg[0:n]))
	}
	conn.Close()
}

func write(conn net.TCPConn){

}