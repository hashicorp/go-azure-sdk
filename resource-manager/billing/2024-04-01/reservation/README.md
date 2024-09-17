
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/reservation` Documentation

The `reservation` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/reservation"
```


### Client Initialization

```go
client := reservation.NewReservationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReservationClient.GetByReservationOrder`

```go
ctx := context.TODO()
id := reservation.NewReservationID("billingAccountValue", "reservationOrderIdValue", "reservationIdValue")

read, err := client.GetByReservationOrder(ctx, id, reservation.DefaultGetByReservationOrderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReservationClient.UpdateByBillingAccount`

```go
ctx := context.TODO()
id := reservation.NewReservationID("billingAccountValue", "reservationOrderIdValue", "reservationIdValue")

payload := reservation.Patch{
	// ...
}


if err := client.UpdateByBillingAccountThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
