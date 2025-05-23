
## `github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/gateways` Documentation

The `gateways` SDK allows for interaction with Azure Resource Manager `hybridcompute` (API Version `2025-01-13`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/gateways"
```


### Client Initialization

```go
client := gateways.NewGatewaysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GatewaysClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := gateways.NewGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "gatewayName")

payload := gateways.Gateway{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `GatewaysClient.Delete`

```go
ctx := context.TODO()
id := gateways.NewGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "gatewayName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `GatewaysClient.Get`

```go
ctx := context.TODO()
id := gateways.NewGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "gatewayName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GatewaysClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GatewaysClient.ListBySubscription`

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


### Example Usage: `GatewaysClient.Update`

```go
ctx := context.TODO()
id := gateways.NewGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "gatewayName")

payload := gateways.GatewayUpdate{
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
