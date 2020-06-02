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

	var board *components.Board
	for {
		flag := 0
		fmt.Println("\n\t\t\tEnter the size of the board: ")
		sizeNow, _ := reader.ReadString('\n')
		sizeNow = strings.TrimSpace(sizeNow)
		size, err := strconv.Atoi(sizeNow)
		if err != nil || (size < 3 || size > 5) {
			fmt.Println("\t\t\tPlease enter a valid number")
			flag = 1
		}
		if flag == 0 {
			board = components.NewBoard(uint8(size))
			break
		}
	}

	var player1, player2 *components.Player

	fmt.Println("\t\t\tEnter the Name of player 1: ")
	name1, _ := reader.ReadString('\n')
	name1 = strings.TrimSpace(name1)

	fmt.Println("\t\t\tEnter the Name of player 2: ")
	name2, _ := reader.ReadString('\n')
	name2 = strings.TrimSpace(name2)

	rand.Seed(time.Now().UTC().UnixNano())
	randomNumber := rand.Intn(10)

	cmd := exec.Command("cmd", "/c", "cls")

	if randomNumber >= 5 {
		player1 = components.NewPlayer(name1, "X")
		player2 = components.NewPlayer(name2, "O")
		time.Sleep(time.Second)
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("\t\t\tComputer chooses", name1, "to play first!")
	} else {
		player1 = components.NewPlayer(name2, "X")
		player2 = components.NewPlayer(name1, "O")
		time.Sleep(time.Second)
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("\t\t\tComputer chooses", name2, "to play first!")
	}
	ourBoardService := service.NewBoardService(board)
	ourResultService := service.NewResultService(ourBoardService)
	ourGameService := service.NewGameService(ourResultService, [2]*components.Player{player1, player2})
	fmt.Println(ourGameService.PrintBoard())

	for {
		var result service.Result
		for {
			flag := 0
			fmt.Println("\t\t\tEnter your position", player1.Name)
			indexNow, _ := reader.ReadString('\n')
			indexNow = strings.TrimSpace(indexNow)
			index, err := strconv.Atoi(indexNow)
			if err != nil {
				fmt.Println(err)
				flag = 1
				continue
			} else {
				time.Sleep(time.Second)
				cmd1 := exec.Command("cmd", "/c", "cls")
				cmd1.Stdout = os.Stdout
				cmd1.Run()
			}

			result, err = ourGameService.Play(uint8(index))
			if err != nil {
				fmt.Println(err)
				flag = 1
				continue
			}
			if flag == 0 {
				break
			}
		}

		fmt.Println(ourGameService.PrintBoard())
		fmt.Println("")
		if result.Win == true {
			fmt.Println("\t\t\tWinner is ", player1.Name)
			break
		}
		if result.Draw == true {
			fmt.Println("\t\t\tIt is a draw")
			break
		}
		for {
			flag := 0
			fmt.Println("\t\t\tEnter your position", player2.Name)
			indexNow, err := reader.ReadString('\n')
			indexNow = strings.TrimSpace(indexNow)
			index, err := strconv.Atoi(indexNow)
			if err != nil {
				fmt.Println(err)
				flag = 1
				continue
			} else {
				time.Sleep(time.Second)
				cmd1 := exec.Command("cmd", "/c", "cls")
				cmd1.Stdout = os.Stdout
				cmd1.Run()
			}

			result, err = ourGameService.Play(uint8(index))
			if err != nil {
				fmt.Println(err)
				flag = 1
				continue
			}
			if flag == 0 {
				break
			}
		}
		fmt.Println(ourGameService.PrintBoard())
		if result.Win == true {
			fmt.Println("\t\t\tWinner is", player2.Name)
			break
		}
		if result.Draw == true {
			fmt.Println("\t\t\tIt's a draw")
			break
		}
	}
}
