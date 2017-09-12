package main


func main(){
	startTcp("localhost:5333")

	signalCH := InitSignal()
	HandleSignal(signalCH)
	// exit
}
