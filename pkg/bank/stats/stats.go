package stats

import (
	"github.com/KarrenAeris/bank/pkg/bank/types"
)

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	n := types.Money(0)
	for _, payment := range payments {
		sum += payment.Amount
		n++
	}
	return sum / n
}