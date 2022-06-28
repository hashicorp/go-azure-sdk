
## `github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/accounts` Documentation

The `accounts` SDK allows for interaction with the Azure Resource Manager Service `datalakeanalytics` (API Version `2016-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/accounts"
```


### Client Initialization

```go
client := accounts.NewAccountsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccountsClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := accounts.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

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
id := accounts.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := accounts.CreateDataLakeAnalyticsAccountParameters{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AccountsClient.Delete`

```go
ctx := context.TODO()
id := accounts.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AccountsClient.Get`

```go
ctx := context.TODO()
id := accounts.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

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
id := accounts.NewSubscriptionID()

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
id := accounts.NewResourceGroupID()

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
id := accounts.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := accounts.UpdateDataLakeAnalyticsAccountParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
