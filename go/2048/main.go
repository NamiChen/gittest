package main

import (
	"fmt"
	"game/main.go/game"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	fmt.Println("Use {W A S D} or Arrow keys to move the board")
	fmt.Printf("Press and key to start\n")
	_, _, err := keyboard.GetSingleKey()
	if err != nil {
		log.Fatalln("error while taking input to start the game ")
	}

	fmt.Println("Getting Started")
	b := game.NewBoard()
	b.AddElement()
	b.AddElement()
	for true {
		if b.IsOver() {
			break
		}
		b.AddElement()
		b.Display()
		res := b.TakeInput()
		if !res {
			return
		}
		fmt.Println("game over")
	}
	// for i := 0; i < 10; i++ {
	// 	b.Display()
	// 	b.AddElement()
	// }
	// fmt.Println("Game over!")
}
