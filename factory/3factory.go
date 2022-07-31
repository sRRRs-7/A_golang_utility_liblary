package factory

import (
	"fmt"
	"log"
)

// struct -> interface -> constructor -> implement

type payment interface {
	pay(amount float32) string
}

type cashPM struct {}
type creditPM struct {}
type cryptCurrencyPM struct {}

const (
	cash = iota
	credit
	cryptCurrency
)

func Factory3() {
	init, err := constructor3(credit)
	if err != nil {
		log.Fatalf("Error constructor: %v", err)
	}
	str := init.pay(100)
	fmt.Println(str)

	init, err = constructor3(cryptCurrency)
	if err != nil {
		log.Fatalf("Error constructor: %v", err)
	}
	str = init.pay(50)
	fmt.Println(str)

	init, err = constructor3(cash)
	if err != nil {
		log.Fatalf("Error constructor: %v", err)
	}
	str = init.pay(10000)
	fmt.Println(str)
}

func constructor3(payId int) (payment, error) {
	switch payId {
	case cash:
		return new(cashPM), nil
	case credit:
		return new(creditPM), nil
	case cryptCurrency:
		return new(cryptCurrencyPM), nil
	default:
		return nil, fmt.Errorf("unsupported payment: %d", payId)
	}
}

func (ca *cashPM) pay(amount float32) string {
	return fmt.Sprintf("successfully Cash payment: %v dollars", amount)
}

func (cr *creditPM) pay(amount float32) string {
	return fmt.Sprintf("successfully Credit payment: %v dollars", amount)
}

func (cc *cryptCurrencyPM) pay(amount float32) string {
	return fmt.Sprintf("successfully Crypt Currency payment: %v BTC", amount)
}