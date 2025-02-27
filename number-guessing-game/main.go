package main

import (
	"fmt"
	"main/levels"
	"os"
)

func main() {
	var i, _ int
	for {
		fmt.Printf("\nWelcome to the Number Guessing Game!\n\n")
		fmt.Printf("I'm thinking of a number between 1 and 100.\n\n")
		fmt.Printf("Please Select Difficulty to begin----\n\n")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium(5 chances)")
		fmt.Println("3. Hard(3 chances)")
		fmt.Println("4. To view High-Score")
		fmt.Println("5. To Exit Game")
		fmt.Scanf("%d\n", &i)

		switch i {
		case 1:
			fmt.Println("Great! You have selected Easy Difficulty Level.")
			fmt.Printf("Lets Start the Game!\n\n")
			j, tim, er := levels.Easy(10)
			if er != nil {
				fmt.Printf("Too Bad you couldnt guess the number,better luck! next time\n")
				fmt.Println("The no. was ", j)
				fmt.Printf("Time taken:- %.2f secs\n", tim.Seconds())
			} else {
				fmt.Printf("Congratulations you won the game in %d attempts\n", j)
				fmt.Printf("Time taken:- %.2f secs\n", tim.Seconds())
			}
			break
		case 2:
			fmt.Println("Great! You have selected Medium Difficulty Level.")
			fmt.Printf("Lets Start the Game!\n\n")
			j, tim, er := levels.Easy(5)
			if er != nil {
				fmt.Printf("Too Bad you couldnt guess the number,better luck! next time\n")
				fmt.Println("The no. was ", j)
				fmt.Printf("Time taken:- %.2f secs\n", tim.Seconds())
			} else {
				fmt.Printf("Congratulations you won the game in %d attempts\n", j)
				fmt.Printf("Time taken:- %.2f secs\n", tim.Seconds())
			}
			break
		case 3:
			fmt.Println("Great! You have selected Hard Difficulty Level.")
			fmt.Printf("Lets Start the Game!\n\n")
			j, tim, er := levels.Easy(3)
			if er != nil {
				fmt.Printf("Too Bad you couldnt guess the number,better luck! next time\n")
				fmt.Println("The no. was ", j)
				fmt.Printf("Time taken:- %.2f secs\n", tim.Seconds())
			} else {
				fmt.Printf("Congratulations you won the game in %d attempts\n", j)
				fmt.Printf("Time taken:- %.2f secs\n", tim.Seconds())
			}
			break
		case 4:
			levels.Read()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Please choose one of the above options")
		}
	}
}
