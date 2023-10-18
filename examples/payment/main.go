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

	client, err := payment.NewClient(apiKey, secretKey)
	if err != nil {
		fmt.Println("error: ", err)
	}

	order, err := createOrder(ctx, client)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(order)

	status, err := getStatus(ctx, client, order.Token)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(status)

	statusByID, err := getStatusByCommerceID(ctx, client, "")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(statusByID)

	statusByFlowOrder, err := getStatusByFlowOrder(ctx, client, "")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(statusByFlowOrder)

}

func createOrder(ctx context.Context, client *payment.Client) (*payment.CreateOrderResponse, error) {
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
	order, err := client.CreateOrder(ctx, orderData)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func getStatus(ctx context.Context, client *payment.Client, token string) (*payment.Status, error) {
	status, err := client.GetStatus(ctx, token)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func getStatusByCommerceID(ctx context.Context, client *payment.Client, commerceID string) (*payment.Status, error) {
	status, err := client.GetStatusByCommerceID(ctx, commerceID)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func getStatusByFlowOrder(ctx context.Context, client *payment.Client, flowOrder string) (*payment.Status, error) {
	status, err := client.GetStatusFlowOrder(ctx, flowOrder)
	if err != nil {
		return nil, err
	}
	return status, nil
}
