
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/nodes` Documentation

The `nodes` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2022-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/nodes"
```


### Client Initialization

```go
client := nodes.NewNodesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NodesClient.ListByDataBoxEdgeDevice`

```go
ctx := context.TODO()
id := nodes.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

// alternatively `client.ListByDataBoxEdgeDevice(ctx, id)` can be used to do batched pagination
items, err := client.ListByDataBoxEdgeDeviceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
