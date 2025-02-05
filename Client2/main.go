package main

import (
	"bufio"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func readMsg(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error while recieving messages")
			return
		} else {
			log.Printf("Recieved: %v", string(message))
		}
	}
}

func inputMsg(msg chan string) {
	reader := bufio.NewReader(os.Stdin)
	msgg, _ := reader.ReadString('\n')
	msg <- msgg
}

var msg chan string

func main() {
	socketURL := "ws://localhost:8000" + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(socketURL, nil)
	if err != nil {
		log.Println("Error while dial:", err)
		return
	}
	msg := make(chan string)
	go readMsg(conn)
	for {
		go inputMsg(msg)
		select {
		case text := <-msg:
			if text == "" {
				log.Println("Sorry")
				continue
			}
			err := conn.WriteMessage(websocket.TextMessage, []byte(text))
			if err != nil {
				log.Println("Error during writing to websocket:", err)
			} else {
				log.Println("Message sent")
			}
		}
	}
}
