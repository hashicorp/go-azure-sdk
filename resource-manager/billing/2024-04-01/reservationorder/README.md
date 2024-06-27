
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/reservationorder` Documentation

The `reservationorder` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/reservationorder"
```


### Client Initialization

```go
client := reservationorder.NewReservationOrderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReservationOrderClient.GetByBillingAccount`

```go
ctx := context.TODO()
id := reservationorder.NewReservationOrderID("billingAccountValue", "reservationOrderIdValue")

read, err := client.GetByBillingAccount(ctx, id, reservationorder.DefaultGetByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
