
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2023-11-01/lots` Documentation

The `lots` SDK allows for interaction with Azure Resource Manager `consumption` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2023-11-01/lots"
```


### Client Initialization

```go
client := lots.NewLotsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LotsClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := lots.NewBillingAccountID("billingAccountIdValue")

// alternatively `client.ListByBillingAccount(ctx, id, lots.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, lots.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LotsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := lots.NewBillingProfileID("billingAccountIdValue", "billingProfileIdValue")

// alternatively `client.ListByBillingProfile(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LotsClient.ListByCustomer`

```go
ctx := context.TODO()
id := lots.NewCustomerID("billingAccountIdValue", "customerIdValue")

// alternatively `client.ListByCustomer(ctx, id, lots.DefaultListByCustomerOperationOptions())` can be used to do batched pagination
items, err := client.ListByCustomerComplete(ctx, id, lots.DefaultListByCustomerOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
