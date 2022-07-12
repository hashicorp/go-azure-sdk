
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/dedicatedhost` Documentation

The `dedicatedhost` SDK allows for interaction with the Azure Resource Manager Service `compute` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/dedicatedhost"
```


### Client Initialization

```go
client := dedicatedhost.NewDedicatedHostClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DedicatedHostClient.ListByHostGroup`

```go
ctx := context.TODO()
id := dedicatedhost.NewHostGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupValue")

// alternatively `client.ListByHostGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByHostGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DedicatedHostClient.Restart`

```go
ctx := context.TODO()
id := dedicatedhost.NewHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostGroupValue", "hostValue")

if err := client.RestartThenPoll(ctx, id); err != nil {
	// handle the error
}
```
