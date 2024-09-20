
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/reservations` Documentation

The `reservations` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/reservations"
```


### Client Initialization

```go
client := reservations.NewReservationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReservationsClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := reservations.NewBillingAccountID("billingAccountName")

// alternatively `client.ListByBillingAccount(ctx, id, reservations.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, reservations.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReservationsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := reservations.NewBillingProfileID("billingAccountName", "billingProfileName")

// alternatively `client.ListByBillingProfile(ctx, id, reservations.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, reservations.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReservationsClient.ListByReservationOrder`

```go
ctx := context.TODO()
id := reservations.NewReservationOrderID("billingAccountName", "reservationOrderId")

// alternatively `client.ListByReservationOrder(ctx, id)` can be used to do batched pagination
items, err := client.ListByReservationOrderComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
