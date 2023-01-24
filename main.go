package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func inputNumber() (int, int, error) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	var n int
	_, e := fmt.Scanln(&n)
	if e != nil {
		stdin.ReadString('\n')
	}
	return n, r, e
}

func main() {
	seedMoney := 1000

L:
	for {
		fmt.Print("please you enter number: ")
		result, randomNumber, err := inputNumber()
		if err != nil {
			fmt.Println("input type only number")
		} else {
			if result != randomNumber {
				switch {
				case seedMoney < 100:
					fmt.Println("not enough seed money < 100")
					break L
				case seedMoney > 100:
					seedMoney -= 100
					fmt.Println("lose seed money! your seed money:", seedMoney)
				case seedMoney == 100:
					seedMoney = 0
					fmt.Println("lose seed money. it is zero! game over")
					break L
				}
			} else {
				if seedMoney >= 5000 {
					fmt.Println("enough seed money >= 5000!")
					break
				} else {
					seedMoney += 500
					fmt.Println("collect! increase seed money!", seedMoney)
				}
			}
		}
	}
}
