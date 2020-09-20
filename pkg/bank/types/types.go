package types

// Money represents the amount of money in minimum units (cents, penny, dirams, etc.)
type Money int64

// Category represents the category in which the payment was made
type Category string

// Payment represents the info about payment
type Payment struct {
	ID int
	Amount Money
	Category Category
}
 