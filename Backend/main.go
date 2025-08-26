package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// define an upgrader for Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 2024,

	//we'll need to check origin of our server
	//this will allow us to make request from react development servr to here

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// define a reader which will listen fornew msgs being sent to our websocket endpoint
func reader(conn *websocket.Conn) {
	for {
		//read in message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		//print that msg
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// define our websocket endpoint
func serveWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	//upgrade this connection to websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	//listen indefinitely for new messages coming through on our websocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	//map our '/ws' endpoint to 'serverWS' function
	http.HandleFunc("/ws", serveWS)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
