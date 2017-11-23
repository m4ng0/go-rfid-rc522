package main

import (
	"fmt"
	"time"

	rfid "github.com/m4ng0/go-rfid-rc522/rfid"
	rc522 "github.com/m4ng0/go-rfid-rc522/rfid/rc522"
)

func main() {
	reader, err := rc522.NewRfidReader()
	if err != nil {
		fmt.Println(err)
		return
	}
	readerChan, err := rfid.NewReaderChan(reader, 1000*time.Millisecond)
	if err != nil {
		fmt.Println(err)
		return
	}
	rfidChan := readerChan.GetChan()
	for {
		select {
		case id := <-rfidChan:
			fmt.Println(id)
		}
	}
}
