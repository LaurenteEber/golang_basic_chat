package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)
	// sockets
	server.OnConnect("connection", func(so socketio.Conn) error {
		so.SetContext("")
		fmt.Println("A new user connected", so.ID())
		return nil
	})

	//http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
