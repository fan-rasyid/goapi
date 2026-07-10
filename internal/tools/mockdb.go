package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"irpan" : {
		AuthToken: "123ABC",
		Username: "irpan",
	},
	"adel" : {
		AuthToken: "456DEF",
		Username: "adel",
	},
	"holy" : {
		AuthToken: "789GHI",
		Username: "holy",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"irpan" : {
		Coins: 100,
		Username: "irpan",
	},
	"adel" : {
		Coins: 200,
		Username: "adel",
	},
	"holy" : {
		Coins: 300,
		Username: "holy",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second*1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok{
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails{
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