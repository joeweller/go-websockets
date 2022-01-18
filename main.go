package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

// create websocket HTTP upgrader using defaults
var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/", indexHandler)
	fmt.Println("starting server on 127.0.0.1:5000")
	http.ListenAndServe("127.0.0.1:5000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("assets/index.html")
	if err == nil {
		w.Write(file)
	} else {
		fmt.Println(err)
	}
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket error: ", err)
		return
	}

	// close connection when event loop ends
	defer ws.Close()

	// event loop
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("websocket read error: ", err)
			break
		}

		fmt.Println("message type:", messageType)
		if messageType == 1 {
			fmt.Println("text: ", string(message))
			message = []byte(sayHello(string((message))))
		} else {
			fmt.Println("binary: ", message)
		}

		err = ws.WriteMessage(messageType, message)

		if err != nil {
			fmt.Println("websocket response error: ", err)
			break
		}
	}
}

func sayHello(text string) string {
	if strings.HasPrefix(text, "my name is") {
		text = strings.Replace(text, "my name is", "hello, ", 1) + "!"
	}
	return text
}
