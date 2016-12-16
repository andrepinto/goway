package proxy

import (
	"log"
	"encoding/json"
)

type BasicLog struct {

}

//noinspection GoUnusedExportedFunction
func NewBasicLog() *BasicLog{
	return &BasicLog{}
}

func(lg *BasicLog) Log(record *LogRecord){
	b, err := json.Marshal(record)
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(b))
}
