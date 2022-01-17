package bully

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	prefix   = "[leaderelection]"
	timepath = "/tmp/time.log"
)

type Event string

const (
	SERVICE_START          Event = "t0"
	ELECTION_START         Event = "t1"
	LEADER_ELECTED         Event = "t2"
	READY                  Event = "t3"
	STABLE_LEADER          Event = "t4"
	SERVICE_STOP           Event = "t5"
	T2_START               Event = "t6"
	T2_ELECTION_START      Event = "t7"
	T2_LEADER_ELECTED      Event = "t8"
	T2_READY               Event = "t9"
	T2_STABLE              Event = "tA"
	T2_STOPPED             Event = "tB"
	T3_START               Event = "tC"
	T3_ELECTION_START      Event = "tD"
	T3_LEADER_ELECTED      Event = "tE"
	T3_READY               Event = "tF"
	T3_STABLE              Event = "tG"
	T3_STOPPED             Event = "tH"
	UNSTABLE               Event = "UNSTABLE"
	LEADER_STOPPED         Event = "LEADER_STOPPED"
	CONNECTION_ESTABLISHED Event = "tX"
)

//Fatal logs with [leaderelection] prefix
func Fatal(msg string) {
	log.Fatalf("%s%s", prefix, msg)
}

//Print logs with [leaderelection] prefix
func PrintTiming(event Event) {
	time := time.Now().UnixMilli()
	str := fmt.Sprintf("%s:%d", event, time)
	if event == ELECTION_START {
		str += "\n" + fmt.Sprintf("%s:%d", T2_ELECTION_START, time)
		str += "\n" + fmt.Sprintf("%s:%d", T3_ELECTION_START, time)
	}
	if event == LEADER_ELECTED {
		str += "\n" + fmt.Sprintf("%s:%d", T2_LEADER_ELECTED, time)
		str += "\n" + fmt.Sprintf("%s:%d", T3_LEADER_ELECTED, time)
	}
	log.Print(str)
	writeToFile(str)
}

//Print logs with [leaderelection] prefix
func Print(msg string) {
	log.Printf("%s%s", prefix, msg)
}

//Print logs with [leaderelection] prefix
func Debug(msg string) {
	writeToFile("DEB:" + msg)
}

func writeToFile(msg string) {
	msg = msg + "\n"
	// If the file doesn't exist, create it, or append to the file
	file, err := os.OpenFile(timepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	b := []byte(msg)
	if _, err := file.Write(b); err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
