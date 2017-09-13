package main


var R *RoomBucket
func main(){
	R = NewBucket()

	startTcp("localhost:5333")
	InitHttp()

	signalCH := InitSignal()
	HandleSignal(signalCH)
	// exit
}
