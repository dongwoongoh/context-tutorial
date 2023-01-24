package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	startMoney = 1000
	loseMoney  = 100
	winMoney   = 500
	fullMoney  = 5000
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
	seedMoney := startMoney

L:
	for {
		fmt.Print("please you enter number: ")
		result, randomNumber, err := inputNumber()
		if err != nil {
			fmt.Println("input type only number")
		} else {
			if result != randomNumber {
				switch {
				case seedMoney < loseMoney:
					fmt.Println("not enough seed money < 100")
					break L
				case seedMoney > loseMoney:
					seedMoney -= loseMoney
					fmt.Println("lose seed money! your seed money:", seedMoney)
				case seedMoney == loseMoney:
					seedMoney = 0
					fmt.Println("lose seed money. it is zero! game over")
					break L
				}
			} else {
				if seedMoney >= fullMoney {
					fmt.Println("enough seed money >= 5000!")
					break
				} else {
					seedMoney += winMoney
					fmt.Println("collect! increase seed money!", seedMoney)
				}
			}
		}
	}
}
