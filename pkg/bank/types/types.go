package types

// Money представляет собой в минимальных единицах (центы, копейки, дирам и т.д. )
type Money int64

// Currency представляет код валют
type Currency string

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN  представляет номер карты
type PAN string

// Card  представляет информацию о платёжной карте
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}
