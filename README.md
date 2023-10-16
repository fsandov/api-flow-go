# Go Client for Flow.cl Commerce API

[![GoDoc](https://godoc.org/github.com/fsandov/api-flow-go?status.svg)](https://godoc.org/github.com/fsandov/api-flow-go)
![License](https://img.shields.io/badge/License-MIT-blue.svg)

Go client (unofficial) for [Flow.cl Commerce API](https://www.flow.cl/docs/api.html), a payment gateway for Chile and Latin America.

## Installation

```bash
go get github.com/fsandov/api-flow-go
```

## Documentation
All the documentation of Flow API Rest can be found [here](https://www.flow.cl/docs/api.html)

## Usage

### Create a client

1. Create a client with your API key and secret key. You can get your keys [here](https://www.flow.cl/app/web/misDatos.php) (you need to be logged in). Its recommended to use environment variables to store your keys.
2. Create a context, you can use the `context.Background()` function. Also you can use the `context.WithTimeout()` function to set a timeout for the request.
3. Create a struct with data to send to Flow API.

Example:

```go
import "github.com/fsandov/api-flow-go"

func main() {
    // Create a context
    ctx := context.Background()
    // Set api key and secret key
    apiKey := ""
    secretKey := ""

    // Define data to send to Flow API
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

    // Create a client with your api key and secret key
    client, err := payment.NewClient(apiKey, secretKey)
    // Send data to Flow API
    order, err := client.CreateOrder(ctx, orderData)
}
```
You can see more examples in the examples folder of this repository.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](LICENSE)
