package main

import (
    "fmt"
    "math/rand"
    "time"
)

func getName() string {
    var name string
    fmt.Println("Welcome to the Casino")
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    if name == "" {
        name = "Player"
    }
    fmt.Printf("Hello %s! Let's play!\n", name)
    return name
}

func getBet(balance uint) uint {
    var bet uint
    for {
        fmt.Printf("Enter your bet (balance = $%d) or press 0 to exit: ", balance)
        _, err := fmt.Scanln(&bet)
        if err != nil {
            fmt.Println("Please enter a valid number")
            continue
        }
        if bet > balance {
            fmt.Println("You don't have enough balance")
            continue
        }
        break
    }
    return bet
}

func generateSymbolArray(symbols map[string]uint) []string {
    var symbolArr []string
    for symbol, count := range symbols {
        for i := uint(0); i < count; i++ {
            symbolArr = append(symbolArr, symbol)
        }
    }
    return symbolArr
}

func getRandomNumber(min, max int) int {
    return min + rand.Intn(max-min+1)
}

func printSpinResult(spin [][]string) {
    fmt.Println("\nYour spin result:")
    for _, row := range spin {
        for _, symbol := range row {
            fmt.Printf("[%s] ", symbol)
        }
        fmt.Println()
    }
    fmt.Println()
}

func getSpin(reel []string, rows, cols int) [][]string {
    result := make([][]string, rows)
    for i := range result {
        result[i] = make([]string, cols)
    }

    for col := 0; col < cols; col++ {
        selected := make(map[int]bool)
        for row := 0; row < rows; row++ {
            for {
                random := getRandomNumber(0, len(reel)-1)
                if !selected[random] {
                    selected[random] = true
                    result[row][col] = reel[random]
                    break
                }
            }
        }
    }
    return result
}

func checkWin(spin [][]string, multipliers map[string]uint) uint {
    // Simple win check: look for three in a row
    winAmount := uint(0)

    // Check rows
    for i := 0; i < len(spin); i++ {
        if spin[i][0] == spin[i][1] && spin[i][1] == spin[i][2] {
            if mult, exists := multipliers[spin[i][0]]; exists {
                winAmount += mult
            }
        }
    }
    return winAmount
}

func main() {
    // Initialize random seed
    rand.Seed(time.Now().UnixNano())

    symbols := map[string]uint{
        "A": 4,
        "B": 7,
        "C": 12,
        "D": 20,
    }

    multipliers := map[string]uint{
        "A": 20,
        "B": 10,
        "C": 5,
        "D": 2,
    }

    balance := uint(200)
    playerName := getName()
    symbolArr := generateSymbolArray(symbols)

    fmt.Printf("%s starts with balance: $%d\n", playerName, balance)

    for balance > 0 {
        bet := getBet(balance)
        if bet == 0 {
            fmt.Println("Thank you for playing!")
            break
        }

        balance -= bet
        spin := getSpin(symbolArr, 3, 3)
        printSpinResult(spin)

        winAmount := checkWin(spin, multipliers) * bet
        if winAmount > 0 {
            fmt.Printf("Congratulations! You won $%d!\n", winAmount)
            balance += winAmount
        }

        fmt.Printf("Current balance: $%d\n", balance)
    }

    if balance == 0 {
        fmt.Println("Game over! You've run out of money!")
    }

    fmt.Printf("Final balance: $%d\n", balance)
}
