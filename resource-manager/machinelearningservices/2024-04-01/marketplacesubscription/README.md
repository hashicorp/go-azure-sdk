
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/marketplacesubscription` Documentation

The `marketplacesubscription` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/marketplacesubscription"
```


### Client Initialization

```go
client := marketplacesubscription.NewMarketplaceSubscriptionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MarketplaceSubscriptionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := marketplacesubscription.NewMarketplaceSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "marketplaceSubscriptionName")

payload := marketplacesubscription.MarketplaceSubscriptionResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `MarketplaceSubscriptionClient.Delete`

```go
ctx := context.TODO()
id := marketplacesubscription.NewMarketplaceSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "marketplaceSubscriptionName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MarketplaceSubscriptionClient.Get`

```go
ctx := context.TODO()
id := marketplacesubscription.NewMarketplaceSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "marketplaceSubscriptionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MarketplaceSubscriptionClient.List`

```go
ctx := context.TODO()
id := marketplacesubscription.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, marketplacesubscription.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, marketplacesubscription.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
