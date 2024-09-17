
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/reservationorders` Documentation

The `reservationorders` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/reservationorders"
```


### Client Initialization

```go
client := reservationorders.NewReservationOrdersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReservationOrdersClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := reservationorders.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, reservationorders.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, reservationorders.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
