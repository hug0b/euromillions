package main

import (
	"fmt"
	"reflect"

	"github.com/hug0b/lottery/utils/set"
)

const MAX_ATTEMPTS int = 1_000_000
const JACKPOT int = 230_000_000
const SCND_PRIZE int = 200_738
const TICKET_PRICE = 2.5

func main() {
	fmt.Println("🤑 Each EuroMillions lottery ticket costs 2,50€. The jackpot for this game is 230 million €!")
	fmt.Println("\n🤷 It doesn't matter what the jackpot is though, because the probability of winning is 0,000 000 72 %")
	fmt.Println("\n😎 This simulation gives you the thrill of playing without wasting money.")
	fmt.Println("\nEnter 5 different numbers from 1 to 50, with spaces between each number. (For example: 5 17 23 42 50)")

	var gridA [5]int
	var gridB [2]int
	var attempts int

	playerNumbers := set.NewSet[int]()
	playerLuckyStarNumbers := set.NewSet[int]()

	for i := 0; i < 5; i++ {
		_, err := fmt.Scan(&gridA[i])

		if err != nil || gridA[i] < 1 || gridA[i] > 50 {
			fmt.Println("Invalid input: please enter 5 positive integers under 50 separated by spaces")
			return
		}

		if playerNumbers.Has(gridA[i]) {
			fmt.Println("Invalid input: each number must be unique")
			return
		}

		playerNumbers.Add(gridA[i])
	}

	fmt.Println("Enter 2 \"lucky star\" numbers from 1 to 12.")

	for i := 0; i < 2; i++ {
		_, err := fmt.Scan(&gridB[i])

		if err != nil || gridB[i] < 1 || gridB[i] > 12 {
			fmt.Println("Invalid input: please enter 2 positive integers under 12 separated by spaces")
			return
		}

		if playerLuckyStarNumbers.Has(gridB[i]) {
			fmt.Println("Invalid input: each number must be unique")
			return
		}

		playerLuckyStarNumbers.Add(gridB[i])
	}

	fmt.Printf("How many times do you want to play? (Max: %d)\n", MAX_ATTEMPTS)

	_, err := fmt.Scan(&attempts)

	if err != nil || attempts < 1 || attempts > MAX_ATTEMPTS {
		fmt.Printf("Invalid input: please enter a positive integer under %d\n", MAX_ATTEMPTS)
	}

	totalCost := float32(attempts) * TICKET_PRICE

	fmt.Printf("It costs %.0f€ to play %d times, but don't worry. I'm sure you'll win it all back.\n", totalCost, attempts)
	fmt.Println("Press Enter to start...")
	fmt.Scanln()

	for i := 0; i < attempts; i++ {
		winningNumbers := set.GetRandIntSet(5, 50)
		winningLuckyStarNumbers := set.GetRandIntSet(2, 12)

		fmt.Printf("The winning numbers are: %-14s and %-5s", winningNumbers.ToElementsString(), winningLuckyStarNumbers.ToElementsString())

		if reflect.DeepEqual(playerNumbers, winningNumbers) {

			correctLuckyStarNumbers := len(winningLuckyStarNumbers.Intersection(playerLuckyStarNumbers))

			// 2 lucky star numbers: won the jackpot
			if correctLuckyStarNumbers == 2 {
				currentCost := float32(i) * TICKET_PRICE
				totalEarned := float32(JACKPOT) - currentCost

				fmt.Println("You have won the EuroMillions Lottery Jackpot! Congratulations!")

				if totalEarned > 0 {
					fmt.Printf("Your total earnings would be %.0f€ if this were real!\n", totalEarned)
				} else {
					fmt.Printf("Sadly by buying all those tickets your total earnings are %.0f€...\n", totalEarned)
				}

				return

				// 1 lucky star number: won the 2nd prize
			} else if correctLuckyStarNumbers == 1 {
				currentCost := float32(i) * TICKET_PRICE
				totalEarned := float32(SCND_PRIZE) - currentCost

				fmt.Println("\nYou have won the EuroMillions second prize! Congratulations!")

				if totalEarned > 0 {
					fmt.Printf("Your total earnings would be %.0f€ if this were real!\n", totalEarned)
				} else {
					fmt.Printf("Sadly by buying all those tickets your total earnings are %.0f€...\n", totalEarned)
				}

				return
			}

		} else {
			fmt.Println(" You lost.")
		}
	}

	fmt.Printf("You have wasted %.0f€.\n", totalCost)
	fmt.Println("Thanks for playing!")
}
