package main

import (
	"context"
	"fmt"

	"github.com/fsandov/api-flow-go/pkg/payment"
)

func main() {
	ctx := context.Background()
	apiKey := ""
	secretKey := ""
	orderData := payment.CreateOrderRequest{
		CommerceOrder:   "",
		Subject:         "",
		Currency:        "",
		Amount:          1,
		Email:           "",
		PaymentMethod:   payment.AllMethods,
		URLConfirmation: "",
		URLReturn:       "",
	}

	client, err := payment.NewClient(apiKey, secretKey)
	if err != nil {
		fmt.Println("error: ", err)
	}
	order, err := client.CreateOrder(ctx, orderData)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("order: ", order)

}
