
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2026-07-15/aigateways` Documentation

The `aigateways` SDK allows for interaction with Azure Resource Manager `web` (API Version `2026-07-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2026-07-15/aigateways"
```


### Client Initialization

```go
client := aigateways.NewAiGatewaysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AiGatewaysClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := aigateways.NewAigatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "aigatewayName")

payload := aigateways.AiGateway{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AiGatewaysClient.Delete`

```go
ctx := context.TODO()
id := aigateways.NewAigatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "aigatewayName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AiGatewaysClient.Get`

```go
ctx := context.TODO()
id := aigateways.NewAigatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "aigatewayName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AiGatewaysClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, aigateways.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, aigateways.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AiGatewaysClient.ListBySubscription`

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


### Example Usage: `AiGatewaysClient.Patch`

```go
ctx := context.TODO()
id := aigateways.NewAigatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "aigatewayName")

payload := aigateways.AiGatewayTagsUpdate{
	// ...
}


read, err := client.Patch(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
