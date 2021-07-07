package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	//declaration of all the variiables that are used
	var choice, choiceStart bool
	var name1, name2, Userstring, guessword, activeperson, inactiveperson string
	var active, inactive int
	var comparestring int

	//printing the welcome statement and starting statements
	fmt.Println("-------WELCOME TO STRING GAME-------")
	fmt.Println("Note: This is a 2 player game, only 2 can play!")
	fmt.Println("Do you want to start the game? enter 1 for YES and 0 for NO")
	fmt.Scan(&choice)

	if choice == true {

		//details of the players
		fmt.Println("Enter the name of player 1: ")
		fmt.Scan(&name1)

		fmt.Println("Enter the name of player 2: ")
		fmt.Scan(&name2)

		fmt.Println("Who wants to start first?  choose 1 for", name1, "and choose 2 for", name2)
		fmt.Scan(&choiceStart)

		if choiceStart == true {
			//storing a copy of entered names in respective required variables according to the condition
			active = 1
			activeperson = name1
			inactiveperson = name2
			fmt.Println("Player 1", name1, "has started the game")

		} else {

			active = 2
			activeperson = name2
			inactiveperson = name1
			fmt.Println("Player 2", name2, "has started the game")

		}

		fmt.Println("Enter the input string: ")

		fmt.Scan(&Userstring)

		//the entered input string will be split into each chracter and will be stored as an array in newArray
		newArray := strings.Split(Userstring, "")

		//fmt.Println(newArray)

		//fmt.Println(len(newArray))

		//to sort the elements in an array in ascending order
		//this is displayed as the rearranged string
		//opponent should find out the actual string from this rearranged string

		sort.Slice(newArray, func(i, j int) bool {

			return newArray[i] < newArray[j]

		})

		fmt.Println("The string entered by player", active, " is rearranged as: ")

		//printing the rearranged string
		for _, v := range newArray {

			fmt.Printf(v)
		}

		fmt.Println("\n=========================================================")

		if active == 1 {
			inactive = 2
			fmt.Println("Now, It's player 2 turn to guess the correct word")
			fmt.Println("Player 2 has 3 chances to guess the correct word")
		} else {
			inactive = 1
			fmt.Println("\nNow, It's player 1 turn to guess the correct word")
			fmt.Println("Player 1 has 3 chances to guess the correct word")
		}

		//while loop for giving the opponent 3 chances and printing who is the winner
		i := 1
		for {
			if i >= 4 {
				break
			}
			fmt.Println("=========================================================")
			fmt.Println("Player ", inactive, ": Attempt---", i)
			fmt.Scan(&guessword)

			comparestring = strings.Compare(guessword, Userstring)

			if comparestring == 0 {
				i = i + 10
				fmt.Println(inactiveperson, " has won the game")

			} else {
				i++
				fmt.Println("The word entered doesn't match")
				fmt.Println("Player ", inactive, "has only", (4 - i), "attempts left")
			}

		}
		if (4 - i) == 0 {
			fmt.Println(activeperson, " has won the game")
		}

	} else {
		fmt.Println("You can start the game when you are ready to play")
	}

}
