package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type LoggerRPC struct{}

type LogMessage struct {
	ClientName     string
	DownloadedBits int64
	UploadedBits   int64
	TimeStamp      int
}

var clientMap map[string]LogMessage = make(map[string]LogMessage)
var startTime time.Time = time.Now()

func (this *LoggerRPC) Log(logMessage LogMessage) {
	clientMap[logMessage.ClientName] = logMessage
}

func ReportSwarmStatus() {
	fmt.Println("ClientInfo: Data")
	for key, value := range clientMap {
		fmt.Printf("%-15s: ", key)
		fmt.Printf("DownLoadedBits %d,", value.DownloadedBits)
		fmt.Printf("UploadedBits %d,", value.UploadedBits)
		fmt.Printf("TimeStamp %d,", time.Since(startTime).Seconds())
		fmt.Println()
	}
}

func ServeRPC() {
	logger := new(LoggerRPC)
	rpc.Register(logger)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":0")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

func main() {
	fmt.Printf("Beginning logger")
	ServeRPC()

	for {
		time.Sleep(time.Second * 2)
		ReportSwarmStatus()
	}
}
