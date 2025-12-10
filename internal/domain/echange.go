package domain

import "time"

type Exchange struct {
	Email     string    `bson:"email"`
	From      string    `bson:"from"`
	To        string    `bson:"to"`
	Amount    string    `bson:"amount"`
	CreatedAt time.Time `bson:"created_at"`
}
