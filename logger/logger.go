package main

import (
	"fmt"
	"github.com/ansonwzj/CPSC416Taipei/loggerdata"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
)

type LoggerRPC struct{}

var clientMap map[string]loggerdata.LogMessage = make(map[string]loggerdata.LogMessage)
var startTime time.Time = time.Now()

func (this *LoggerRPC) Log(logMessage *loggerdata.LogMessage, logReply *loggerdata.LogReply) error {
	log.Println("Received message from client")
	clientMap[logMessage.ClientName] = *logMessage
	return nil
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

	if len(os.Args) < 3 {
		fmt.Println("Usage: logger.exe ip port")
	}

	ip := os.Args[1]
	port := os.Args[2]

	logger := new(LoggerRPC)
	rpc.Register(logger)
	rpc.HandleHTTP()
	log.Println("listening on" + ip + ":" + port)
	l, e := net.Listen("tcp", ip+":"+port)
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
