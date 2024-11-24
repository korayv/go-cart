package main

import "fmt"

type GiftCard struct {
	CheckBalance func() string
	DeductAmount func(int) int
}

func giftcardactivate() GiftCard {
	balance := 100
	return GiftCard{
		CheckBalance: func() string {
			return fmt.Sprintf("Your balance is %d", balance)
		},
		DeductAmount: func(amount int) int {
			balance -= amount
			return balance
		},
	}
}
func main() {
	cardHolder1 := giftcardactivate()
	cardHolder2 := giftcardactivate()
	cardHolder1.CheckBalance()
	cardHolder2.CheckBalance()
	cardHolder1.DeductAmount(30)
	cardHolder2.DeductAmount(10)
	cardHolder1.CheckBalance()
	cardHolder2.CheckBalance()
}
