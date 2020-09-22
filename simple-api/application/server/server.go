package server

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
	""
)

type Room struct {
	Name string `json:"name"`
	Guid string `json:"guid"`
	Capacity int `json:"capacity"`
	HostUser string `json:"host_user"`
	Participants []string `json:"participants"`
}

var Rooms []Room

func handleRequests() {
	http.HandleFunc("/rooms", listRooms)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Rooms = []Room{
		Room{Name: "Games", Guid: "1", Capacity: 5, HostUser: "victoriaovilas", Participants: []string{
			"victoriaovilas", "pefreis",
		}},
		Room{Name: "Book Club", Guid: "2", Capacity: 10, HostUser: "helencunha", Participants: []string{
			"helencunha", "flaviog", "victoriaovilas", "pefreis",
		}},
	}
	handleRequests()
}

func listRooms(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: listRooms")
	json.NewEncoder(w).Encode(Rooms)
}

