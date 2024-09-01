package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/joaofilippe/go-websocket/models"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Chat Room!")
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
