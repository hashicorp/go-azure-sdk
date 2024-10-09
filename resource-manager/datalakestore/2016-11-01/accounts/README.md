
## `github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/accounts` Documentation

The `accounts` SDK allows for interaction with Azure Resource Manager `datalakestore` (API Version `2016-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/accounts"
```


### Client Initialization

```go
client := accounts.NewAccountsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccountsClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := accounts.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := accounts.CheckNameAvailabilityParameters{
	// ...
}


read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccountsClient.Create`

```go
ctx := context.TODO()
id := accounts.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

payload := accounts.CreateDataLakeStoreAccountParameters{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AccountsClient.Delete`

```go
ctx := context.TODO()
id := accounts.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AccountsClient.EnableKeyVault`

```go
ctx := context.TODO()
id := accounts.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

read, err := client.EnableKeyVault(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccountsClient.Get`

```go
ctx := context.TODO()
id := accounts.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccountsClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, accounts.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, accounts.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AccountsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, accounts.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, accounts.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AccountsClient.Update`

```go
ctx := context.TODO()
id := accounts.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

payload := accounts.UpdateDataLakeStoreAccountParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
