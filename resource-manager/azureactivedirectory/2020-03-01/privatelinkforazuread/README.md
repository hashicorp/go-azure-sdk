
## `github.com/hashicorp/go-azure-sdk/resource-manager/azureactivedirectory/2020-03-01/privatelinkforazuread` Documentation

The `privatelinkforazuread` SDK allows for interaction with the Azure Resource Manager Service `azureactivedirectory` (API Version `2020-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/azureactivedirectory/2020-03-01/privatelinkforazuread"
```


### Client Initialization

```go
client := privatelinkforazuread.NewPrivateLinkForAzureAdClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkForAzureAdClient.Create`

```go
ctx := context.TODO()
id := privatelinkforazuread.NewPrivateLinkForAzureAdID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkForAzureAdValue")

payload := privatelinkforazuread.PrivateLinkPolicy{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateLinkForAzureAdClient.Delete`

```go
ctx := context.TODO()
id := privatelinkforazuread.NewPrivateLinkForAzureAdID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkForAzureAdValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkForAzureAdClient.Get`

```go
ctx := context.TODO()
id := privatelinkforazuread.NewPrivateLinkForAzureAdID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkForAzureAdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkForAzureAdClient.List`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateLinkForAzureAdClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateLinkForAzureAdClient.Update`

```go
ctx := context.TODO()
id := privatelinkforazuread.NewPrivateLinkForAzureAdID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkForAzureAdValue")

payload := privatelinkforazuread.PrivateLinkPolicyUpdateParameter{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
