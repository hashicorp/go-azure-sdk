
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/l3networks` Documentation

The `l3networks` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/l3networks"
```


### Client Initialization

```go
client := l3networks.NewL3NetworksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `L3NetworksClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := l3networks.NewL3NetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "l3NetworkName")

payload := l3networks.L3Network{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, l3networks.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `L3NetworksClient.Delete`

```go
ctx := context.TODO()
id := l3networks.NewL3NetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "l3NetworkName")

if err := client.DeleteThenPoll(ctx, id, l3networks.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `L3NetworksClient.Get`

```go
ctx := context.TODO()
id := l3networks.NewL3NetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "l3NetworkName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `L3NetworksClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, l3networks.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, l3networks.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `L3NetworksClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, l3networks.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, l3networks.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `L3NetworksClient.Update`

```go
ctx := context.TODO()
id := l3networks.NewL3NetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "l3NetworkName")

payload := l3networks.L3NetworkPatchParameters{
	// ...
}


read, err := client.Update(ctx, id, payload, l3networks.DefaultUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
