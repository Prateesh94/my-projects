package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var wg sync.WaitGroup
var mu sync.Mutex
var msgcount int = 0
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
	"You miss 100% of the shots you don't take.",
	"Life is what happens when you're busy making other plans.",
	"The purpose of our lives is to be happy.",
	"Get busy living or get busy dying.",
	"You only live once, but if you do it right, once is enough.",
	"In the end, we only regret the chances we didn't take.",
	"Life is either a daring adventure or nothing at all.",
	"Time and tide wait for no one",
	"By being yourself, you put something wonderful in the world that was not there before.",
	"Believe you can and you're halfway there.",
	"Act as if what you do makes a difference. It does.",
	"Success is not final, failure is not fatal: It is the courage to continue that counts.",
	"Don't watch the clock; do what it does. Keep going.",
	"Success usually comes to those who are too busy to be looking for it.",
	"Opportunities don't happen, you create them.",
	"The only limit to our realization of tomorrow will be our doubts of today.",
	"Do what you can, with what you have, where you are.",
	"The only way to do great work is to love what you do.",
	"Success is walking from failure to failure with no loss of enthusiasm.",
	"Success is not the key to happiness. Happiness is the key to success.",
	"Success is not in what you have, but who you are.",
	"Success is not about how much money you make, but the difference you make in people's lives.",
	"Success is not about being the best. It's about always getting better."}

func mess(conn *websocket.Conn, id int) {
	//log.Printf("Client %d: Sending message: %s\n", id, message)
	rand.NewSource(time.Now().UnixNano())

	randomIndex := rand.Intn(len(message))
	randomString := message[randomIndex]
	time.Sleep(5 * time.Second)
resend:
	err := conn.WriteMessage(websocket.TextMessage, []byte(randomString+" from Client:- "+strconv.Itoa(id)))

	if err != nil {
		time.Sleep(5 * time.Second) // Wait before retrying
		goto resend
	} else {
		msgcount++
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println("Read Error:", err)
		return
	}
	log.Printf("Client %d: Received: %s\n", id, msg)
	//time.Sleep(1 * time.Second) // Simulate some processing time
}
func createClient(id int) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	//log.Printf("Connecting to %s\n", u.String())
redial:
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	defer conn.Close()
	defer wg.Done()
	if err != nil {
		red := "\033[31m"
		reset := "\033[0m"
		fmt.Printf("%sConnection failed,retrying in 5 sec%s", red, reset)
		time.Sleep(5 * time.Second)
		// Wait before retrying
		goto redial
	}
	//for range 50 {
	//time.Sleep(5 * time.Second) // Simulate some delay between messages
	time.Sleep(2 * time.Minute)
	for range 5 {
		//	mu.Lock()
		go mess(conn, id)
		//	mu.Unlock()
		time.Sleep(5 * time.Second)
	}
	//time.Sleep(3 * time.Second) // Simulate some delay between messages
	// Simulate some delay between messages

	//}
	fmt.Printf("Client with id:-%d has sent all messages, now closing connection\n", id)
	// if err != nil {
	// 	log.Fatal("Dial Error:", err)
	// }

}

func main() {
	//	messages := "Hello, World!"
	//message := "Load Test Message: " + messages
	i := 0
	for i < 10000 {
		wg.Add(1)
		go createClient(i + 1)
		i++

		time.Sleep(10 * time.Millisecond) // Add a delay to avoid connection issues
	}

	// Wait for all clients to finish
	fmt.Println("Waiting for all clients to finish...")
	//	time.Sleep(30 * time.Second) // Adjust this as needed
	wg.Wait()
	time.Sleep(10 * time.Second)
	fmt.Printf("Total Clients created: %d\n", i)
	fmt.Println("Total messages sent: ", msgcount)
	//time.Sleep(5 * time.Second) // Wait for all clients to finish
}
