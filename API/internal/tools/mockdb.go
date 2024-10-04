package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Smit": {
		AuthToken: "123ABC",
		Username:  "Smit",
	},
	"Pablo": {
		AuthToken: "456DEF",
		Username:  "Pablo",
	},
	"Juan": {
		AuthToken: "789GHI",
		Username:  "Juan",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"Smit": {
		Coins:    100,
		Username: "Smit",
	},
	"Pablo": {
		Coins:    200,
		Username: "Pablo",
	},
	"Juan": {
		Coins:    300,
		Username: "Juan",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
