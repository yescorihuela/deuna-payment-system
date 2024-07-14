package entities

import (
	"time"

	"github.com/jaevor/go-nanoid"
)

func NewNanoId() string {
	newNanoId, err := nanoid.CustomASCII("0123456789", 8)
	if err != nil {
		panic(err)
	}
	return newNanoId()
}

type Merchant struct {
	Id                int
	Name              string
	Balance           float64
	NotificationEmail string
	MerchantCode      string
	Enabled           bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
