package main

import (
	"net/http"
	"time"
	"net"
	log "github.com/alecthomas/log4go"
	"encoding/json"
	"comet/define"
	"comet/libs"
)

func InitHttp(){
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("/v1/push/room",PushRoom)

	httpServer := &http.Server{Handler:httpServeMux,ReadTimeout:1*time.Second,WriteTimeout:1*time.Second}
	httpServer.SetKeepAlivesEnabled(true)
	l,_ := net.Listen("tcp",":5999")
	if err:=httpServer.Serve(l);err != nil{
		log.Error("http serve err:%s",err)
		panic(err)
	}
}


func PushRoom(w http.ResponseWriter,r *http.Request){
	var res = map[string]interface{}{"ret":"abc"}
	var body string = "hello"
	p := libs.Proto{SeqId:1}
	arg := &define.BoardcastRoomArg{Rid:"123",P:p}
	R.pusharg <- arg
	defer retPWrite(w,r,res,&body,time.Now())
}

func retPWrite(w http.ResponseWriter, r *http.Request, res map[string]interface{}, body *string, start time.Time) {
	data, err := json.Marshal(res)
	if err != nil {
		log.Error("json.Marshal(\"%v\") error(%v)", res, err)
		return
	}
	dataStr := string(data)
	if _, err := w.Write([]byte(dataStr)); err != nil {
		log.Error("w.Write(\"%s\") error(%v)", dataStr, err)
	}
	log.Info("req: \"%s\", post: \"%s\", res:\"%s\", ip:\"%s\", time:\"%fs\"", r.URL.String(), *body, dataStr, r.RemoteAddr, time.Now().Sub(start).Seconds())
}