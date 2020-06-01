package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"tic_tac_toe_app/components"
	"tic_tac_toe_app/service"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n\t\t\tWelcome to 2 player Tic Tac Toe Game")
	fmt.Println("\n\t\t\tInstructions :")
	fmt.Println("\t\t\t1. You have to enter the number of rows.")
	fmt.Println("\t\t\t2. It should be an Integer greater than 2 and less than 6.")
	fmt.Println("\t\t\t3. After entering the dimension, type the names of the 2 players.")
	fmt.Println("\t\t\t4. The computer will decide at random who goes first.")
	fmt.Println("\t\t\t5. Player who goes first will be displayed and get assigned X.")
	fmt.Println("\t\t\t6. The positions start with 0 at first row and first column.")

FLAG:
	fmt.Println("\n\t\t\tEnter the number of rows to start")
	dimensions, _ := reader.ReadString('\n')
	dimensions = strings.TrimSpace(dimensions)
	dimension, err := strconv.Atoi(dimensions)

	if err != nil || (dimension < 3 || dimension > 5) {
		fmt.Println("\t\t\tYou were supposed to enter an Integer greater than 2 and lesser than 6.")
		goto FLAG
	}

	fmt.Println("\t\t\tType the name of Player 1: ")

	player1, _ := reader.ReadString('\n')

	player1 = strings.TrimSpace(player1)

	fmt.Println("\t\t\tType the name of Player 2: ")

	player2, _ := reader.ReadString('\n')

	player2 = strings.TrimSpace(player2)

	rand.Seed(time.Now().UTC().UnixNano())

	randomNumber := rand.Intn(10)

	var first, second *components.Player
	cmd := exec.Command("cmd", "/c", "cls")

	if randomNumber >= 5 {
		first = components.NewPlayer(player1, "X")
		second = components.NewPlayer(player2, "O")
		time.Sleep(time.Second)
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("\t\t\tComputer chooses", player1, "to play first!")
	} else {
		first = components.NewPlayer(player2, "X")
		second = components.NewPlayer(player1, "O")
		time.Sleep(time.Second)
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("\t\t\tComputer chooses", player2, "to play first!")
	}

	newGame := service.NewGameService(first, second, dimension)

	var result string
	start := true

	fmt.Println("\t\t\tEnter your desired position number to mark")
	for {

		fmt.Print(newGame.PrintBoard())

		fmt.Println("Enter position number : ")

		pos, _ := reader.ReadString('\n')
		pos = strings.TrimSpace(pos)
		position, err := strconv.Atoi(pos)

		if err == nil {
			time.Sleep(time.Second)
			cmd1 := exec.Command("cmd", "/c", "cls")
			cmd1.Stdout = os.Stdout
			cmd1.Run()
		} else {
			fmt.Println("\t\t\tPosition number should be an integer")
			continue
		}

		if start {
			result, err = newGame.Play(uint8(position), first)
		} else {
			result, err = newGame.Play(uint8(position), second)
		}

		if err != nil {
			fmt.Println(err)
			continue
		} else if result == "win" {
			if start {
				fmt.Print(newGame.PrintBoard())
				fmt.Println(first.Name, "is the winner! ")
				break
			} else {
				fmt.Print(newGame.PrintBoard())
				fmt.Println(second.Name, "is the winner! ")
				break
			}
		} else if result == "draw" {
			fmt.Print("\t\t\tThe match is draw")
			break
		}

		if !start {
			fmt.Print("\t\t\t", first.Name, " your chance \n")
		} else {
			fmt.Print("\t\t\t", second.Name, " your chance \n")
		}
		start = !start
	}
}
