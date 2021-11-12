package card

import "github.com/mrGreatProgrammer/bank/pkg/bank/types"

// Withdraw снимаает деньги с карты
func Withdraw(card *types.Card, amount types.Money)  {
	const withdrawLimit = 20_000_00
	if amount < 0 {
		return
	}

	if amount > withdrawLimit {
		return
	}

	if !card.Active {
		return
	}

	if card.Balance < amount {
		return
	}

	card.Balance -= amount
}