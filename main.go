package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	gtfs "github.com/goridethebus/com.google.transit.realtime"
	"google.golang.org/protobuf/proto"
)

func main() {
	resp, err := http.Get("https://tmgtfsprd.sorttrpcloud.com/TMGTFSRealTimeWebService/tripupdate/tripupdates.pb")
	if err != nil {
		fmt.Print("Error encountered")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tripupdate gtfs.TripUpdate
	if err := proto.Unmarshal(body, &tripupdate); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println(tripupdate.Trip)
}
