package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func inputNumber() (int, error) {
	var n int
	_, e := fmt.Scanln(&n)
	if e != nil {
		stdin.ReadString('\n')
	}
	return n, e
}

func main() {
	rand.Seed(time.Now().UnixNano())
	seedMoney := 1000
	randomNumber := rand.Intn(5) + 1

	for {
		fmt.Print("please you enter number: ")
		result, err := inputNumber()
		if err != nil {
			fmt.Println("input type only number")
		} else {
			if result != randomNumber {
				switch {
				case seedMoney < 100:
					fmt.Println("not enough seed money < 100")
					break
				case seedMoney > 100:
					seedMoney -= 100
					fmt.Println("lose seed money! your seed money:", seedMoney)
				case seedMoney == 100:
					seedMoney = 0
					fmt.Println("lose seed money. it is zero! game over")
					break
				}
			} else {
				if seedMoney >= 5000 {
					fmt.Println("enough seed money >= 5000!")
					break
				} else {
					seedMoney += 500
					fmt.Println("collect! increase seed money!", seedMoney)
					break
				}
			}
		}
	}
}
