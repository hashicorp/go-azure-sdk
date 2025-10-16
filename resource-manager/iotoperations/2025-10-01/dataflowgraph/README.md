
## `github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/dataflowgraph` Documentation

The `dataflowgraph` SDK allows for interaction with Azure Resource Manager `iotoperations` (API Version `2025-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/dataflowgraph"
```


### Client Initialization

```go
client := dataflowgraph.NewDataflowGraphClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataflowGraphClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := dataflowgraph.NewDataflowGraphID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "dataflowProfileName", "dataflowGraphName")

payload := dataflowgraph.DataflowGraphResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataflowGraphClient.Delete`

```go
ctx := context.TODO()
id := dataflowgraph.NewDataflowGraphID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "dataflowProfileName", "dataflowGraphName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DataflowGraphClient.Get`

```go
ctx := context.TODO()
id := dataflowgraph.NewDataflowGraphID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "dataflowProfileName", "dataflowGraphName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataflowGraphClient.ListByDataflowProfile`

```go
ctx := context.TODO()
id := dataflowgraph.NewDataflowProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "dataflowProfileName")

// alternatively `client.ListByDataflowProfile(ctx, id)` can be used to do batched pagination
items, err := client.ListByDataflowProfileComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
