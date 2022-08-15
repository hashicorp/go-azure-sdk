
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/reservationtransactions` Documentation

The `reservationtransactions` SDK allows for interaction with the Azure Resource Manager Service `consumption` (API Version `2019-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/reservationtransactions"
```


### Client Initialization

```go
client := reservationtransactions.NewReservationTransactionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReservationTransactionsClient.List`

```go
ctx := context.TODO()
id := reservationtransactions.NewBillingAccountID("billingAccountIdValue")

// alternatively `client.List(ctx, id, reservationtransactions.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, reservationtransactions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReservationTransactionsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := reservationtransactions.NewBillingProfileID("billingAccountIdValue", "billingProfileIdValue")

// alternatively `client.ListByBillingProfile(ctx, id, reservationtransactions.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, reservationtransactions.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
