
## `github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2021-04-01-preview/vnetpeering` Documentation

The `vnetpeering` SDK allows for interaction with the Azure Resource Manager Service `databricks` (API Version `2021-04-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2021-04-01-preview/vnetpeering"
```


### Client Initialization

```go
client := vnetpeering.NewVNetPeeringClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `VNetPeeringClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := vnetpeering.NewVirtualNetworkPeeringID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "peeringValue")

payload := vnetpeering.VirtualNetworkPeering{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `VNetPeeringClient.Delete`

```go
ctx := context.TODO()
id := vnetpeering.NewVirtualNetworkPeeringID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "peeringValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `VNetPeeringClient.Get`

```go
ctx := context.TODO()
id := vnetpeering.NewVirtualNetworkPeeringID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "peeringValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VNetPeeringClient.ListByWorkspace`

```go
ctx := context.TODO()
id := vnetpeering.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")
// alternatively `client.ListByWorkspace(ctx, id)` can be used to do batched pagination
items, err := client.ListByWorkspaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
