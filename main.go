package main

import "fmt"

func getName() string {
	name := ""
	fmt.Println("Welcome to the Casino")
	fmt.Printf("Enter you name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return "Error"
	}
	// Fixed the Printf statement
	fmt.Printf("Hello %s! Let's play!\n", name)
	return name
}

func getBet(balance uint) uint {
	var bet uint

	for true {
		fmt.Printf("Enter your bet (balance = $%d) or press 0 to exit :", balance)
		fmt.Scanln(&bet)

		if bet > balance {
			fmt.Println("You dont have enough balance")
		} else {
			break
		}
	}

	return bet
}

func generateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol, count := range symbols {
		for i := 0; i < int(count); i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}

	return symbolArr
}

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	// multipliers := map[string]uint{
	// 	"A" : 20,
	// 	"B" : 10,
	// 	"C" : 5,
	// 	"D" : 2
	// };
	balance := uint(200)
	getName()

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}

		balance -= bet
	}

	symbolArr := generateSymbolArray(symbols)
	fmt.Println(symbolArr)
	fmt.Printf("Registered the Bet , Now you are left with (balance = $%d) \n", balance)
}
