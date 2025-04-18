/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

func ConnectServer(port string) {
	url := "ws://localhost:" + port + "/ws"
	fmt.Println("Connecting to server on port:", port)
	con, _, er := websocket.DefaultDialer.Dial(url, nil)
	if er != nil {
		fmt.Println("Error connecting to server on port:", port)
		return
	}
	defer con.Close()
	fmt.Println("Connected to server on port:", port)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go WriteMessage(con)
	go ReadMessage(con)
	if _, ok := <-interrupt; ok {
		fmt.Println("Interrupt signal received, closing connection")
		con.Close()
		return
	}
	// Here you would add the logic to connect to the server
}

func WriteMessage(con *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		if scanner.Scan() {
			msg := scanner.Bytes()
			if err := con.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Println("Error writing message:", err)
				return
			}
		}
	}
}

func ReadMessage(con *websocket.Conn) {
	for {
		_, msg, err := con.ReadMessage()
		if err != nil {
			fmt.Println("Connection closed")
			return
		}
		fmt.Println("Received:- ", string(msg))
	}
}

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to a broadcast server on specified port",
	Long: `For example:
broadcast-cli connect 8080`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("connect called")
		var i int
		var er error
		if len(args) == 0 {
			i = 8080
		} else {
			i, er = strconv.Atoi(args[0])
			if er != nil || i < 1024 || i > 65535 {
				fmt.Println("Port number should be between 1024 and 65535")
				return
			}
		}
		if i == 0 {
			i = 8080
		}
		ConnectServer(strconv.Itoa(i))
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
