
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/dedicatedhosts` Documentation

The `dedicatedhosts` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2024-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/dedicatedhosts"
```


### Client Initialization

```go
client := dedicatedhosts.NewDedicatedHostsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DedicatedHostsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := commonids.NewDedicatedHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupName", "hostName")

payload := dedicatedhosts.DedicatedHost{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DedicatedHostsClient.Delete`

```go
ctx := context.TODO()
id := commonids.NewDedicatedHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupName", "hostName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DedicatedHostsClient.Get`

```go
ctx := context.TODO()
id := commonids.NewDedicatedHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupName", "hostName")

read, err := client.Get(ctx, id, dedicatedhosts.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DedicatedHostsClient.ListAvailableSizes`

```go
ctx := context.TODO()
id := commonids.NewDedicatedHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupName", "hostName")

// alternatively `client.ListAvailableSizes(ctx, id)` can be used to do batched pagination
items, err := client.ListAvailableSizesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DedicatedHostsClient.ListByHostGroup`

```go
ctx := context.TODO()
id := commonids.NewDedicatedHostGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupName")

// alternatively `client.ListByHostGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByHostGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DedicatedHostsClient.Redeploy`

```go
ctx := context.TODO()
id := commonids.NewDedicatedHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupName", "hostName")

if err := client.RedeployThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DedicatedHostsClient.Restart`

```go
ctx := context.TODO()
id := commonids.NewDedicatedHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupName", "hostName")

if err := client.RestartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DedicatedHostsClient.Update`

```go
ctx := context.TODO()
id := commonids.NewDedicatedHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupName", "hostName")

payload := dedicatedhosts.DedicatedHostUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
