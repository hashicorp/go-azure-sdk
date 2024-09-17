
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-03-01/dedicatedhost` Documentation

The `dedicatedhost` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-03-01/dedicatedhost"
```


### Client Initialization

```go
client := dedicatedhost.NewDedicatedHostClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DedicatedHostClient.ListAvailableSizes`

```go
ctx := context.TODO()
id := commonids.NewDedicatedHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupValue", "hostValue")

read, err := client.ListAvailableSizes(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DedicatedHostClient.ListByHostGroup`

```go
ctx := context.TODO()
id := commonids.NewDedicatedHostGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupValue")

// alternatively `client.ListByHostGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByHostGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DedicatedHostClient.Redeploy`

```go
ctx := context.TODO()
id := commonids.NewDedicatedHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupValue", "hostValue")

if err := client.RedeployThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DedicatedHostClient.Restart`

```go
ctx := context.TODO()
id := commonids.NewDedicatedHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupValue", "hostValue")

if err := client.RestartThenPoll(ctx, id); err != nil {
	// handle the error
}
```
