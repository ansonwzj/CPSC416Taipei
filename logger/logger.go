package main

import (
	"container/list"
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
var clientList *list.List = list.New()
var startTime time.Time = time.Now()

func (this *LoggerRPC) Log(logMessage *loggerdata.LogMessage, logReply *loggerdata.LogReply) error {
	//log.Println("Received message from client")
	if _, ok := clientMap[logMessage.ClientName]; !ok {
		clientList.PushBack(logMessage.ClientName)
	}
	logMessage.TimeStamp = time.Since(startTime).Seconds()
	clientMap[logMessage.ClientName] = *logMessage

	return nil
}

func ReportSwarmStatus() {
	format := "%-20s %-20s %-20s %-20s\n"
	fmt.Printf(format, "ClientID:", "DownloadedBits", "UploadedBits", "Timestamp")

	for e := clientList.Front(); e != nil; e = e.Next() {
		clientID := e.Value.(string)
		logMessage := clientMap[clientID]

		fmt.Printf("%-20s %-20d %-20d %-20f\n",
			clientID,
			logMessage.DownloadedBits,
			logMessage.UploadedBits,
			logMessage.TimeStamp)
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
	log.Println("listening on: " + ip + ":" + port)
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
