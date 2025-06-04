package main

import (
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var mu sync.Mutex
var message = []string{"You have to believe in yourself when no one else does.",
	"When you have a dream, you’ve got to grab it and never let go.",
	"The most important thing is to enjoy your life—to be happy—it's all that matters.",
	"Spread love everywhere you go. Let no one ever come without leaving happier.",
	"The biggest adventure you can take is to live the life of your dreams.",
	"The only thing we have to fear is fear itself.",
	"Some people want it to happen, some wish it would happen, others make it happen.",
	"You've got to be in it to win it.",
	"Success is not how high you have climbed, but how you make a positive difference to the world.",
	"The future belongs to those who believe in the beauty of their dreams.",
	"The best way to predict the future is to create it.",
	"You miss 100%% of the shots you don't take.",
	"Life is what happens when you're busy making other plans.",
	"The purpose of our lives is to be happy.",
	"Get busy living or get busy dying.",
	"You only live once, but if you do it right, once is enough.",
	"In the end, we only regret the chances we didn't take.",
	"Life is either a daring adventure or nothing at all.",
	"Time and tide wait for no one"}

func mess(conn *websocket.Conn, id int) {
	//log.Printf("Client %d: Sending message: %s\n", id, message)
	rand.NewSource(time.Now().UnixNano())

	randomIndex := rand.Intn(len(message))
	randomString := message[randomIndex]

	err := conn.WriteMessage(websocket.TextMessage, []byte(randomString+" from Client:- "+strconv.Itoa(id)))

	if err != nil {
		log.Println("Write Error:", err)
		return
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println("Read Error:", err)
		return
	}
	log.Printf("Client %d: Received: %s\n", id, msg)
	time.Sleep(1 * time.Second) // Simulate some processing time
}
func createClient(id int) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	log.Printf("Connecting to %s\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	for range 20 {
		mu.Lock()
		go mess(conn, id)
		mu.Unlock()
		time.Sleep(1 * time.Second) // Simulate some delay between messages

	}
	if err != nil {
		log.Fatal("Dial Error:", err)
	}

}

func main() {
	//	messages := "Hello, World!"
	//message := "Load Test Message: " + messages
	for i := 0; i < 1000; i++ {

		go createClient(i + 1)

		time.Sleep(1 * time.Second) // Add a delay to avoid connection issues
	}
	//time.Sleep(5 * time.Second) // Wait for all clients to finish
}
