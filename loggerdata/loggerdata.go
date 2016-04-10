package loggerdata

type LogMessage struct {
	ClientName     string
	DownloadedBits uint64
	UploadedBits   uint64
	Percentage     float32
	TimeStamp      float64
}

type LogReply struct {
}
