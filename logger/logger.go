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
var downloadf *os.File
var uploadf *os.File
var ratiof *os.File

func (this *LoggerRPC) Log(logMessage *loggerdata.LogMessage, logReply *loggerdata.LogReply) error {
	//log.Println("Received message from client")
	if _, ok := clientMap[logMessage.ClientName]; !ok {
		clientList.PushBack(logMessage.ClientName)
	}
	logMessage.TimeStamp = time.Since(startTime).Seconds()
	clientMap[logMessage.ClientName] = *logMessage

	return nil
}

// Structure of csv:
// DownLoadedBitsTable
// | TimeStamp | ClientID# | ClientID2... |
// | 2         | 2003      | 20132...     |
// | 2         | 2003      | 20132...     |
// | 2         | 2003      | 20132...     |

func ReportSwarmStatus() {

	downloadAcc := fmt.Sprintf("%f", time.Since(startTime).Seconds())
	uploadAcc := fmt.Sprintf("%f", time.Since(startTime).Seconds())
	ratioAcc := fmt.Sprintf("%f", time.Since(startTime).Seconds())

	for e := clientList.Front(); e != nil; e = e.Next() {
		clientID := e.Value.(string)
		logMessage := clientMap[clientID]

		downloadAcc = fmt.Sprintf("%s,%d", downloadAcc, logMessage.DownloadedBits)
		uploadAcc = fmt.Sprintf("%s,%d", uploadAcc, logMessage.UploadedBits)
		ratioAcc = fmt.Sprintf("%s,%f", ratioAcc, logMessage.Percentage)
	}

	downloadf.WriteString(downloadAcc + "\n")
	uploadf.WriteString(uploadAcc + "\n")
	ratiof.WriteString(ratioAcc + "\n")

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
	fmt.Println("Beginning logger")
	format := "%-20s %-20s %-20s %-20s\n"
	fmt.Printf(format, "ClientID:", "DownloadedBits", "UploadedBits", "Timestamp")

	downloadf, _ = os.OpenFile("download.csv", os.O_RDWR|os.O_CREATE, 0660)
	uploadf, _ = os.OpenFile("upload.csv", os.O_RDWR|os.O_CREATE, 0660)
	ratiof, _ = os.OpenFile("ratiof.csv", os.O_RDWR|os.O_CREATE, 0660)
	defer downloadf.Close()
	defer ratiof.Close()
	defer uploadf.Close()

	ServeRPC()

	for {
		time.Sleep(time.Second / 2)
		ReportSwarmStatus()
	}
}
