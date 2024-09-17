
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/shares` Documentation

The `shares` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/shares"
```


### Client Initialization

```go
client := shares.NewSharesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SharesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := shares.NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "shareValue")

payload := shares.Share{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SharesClient.Delete`

```go
ctx := context.TODO()
id := shares.NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "shareValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SharesClient.Get`

```go
ctx := context.TODO()
id := shares.NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "shareValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SharesClient.ListByDataBoxEdgeDevice`

```go
ctx := context.TODO()
id := shares.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue")

// alternatively `client.ListByDataBoxEdgeDevice(ctx, id)` can be used to do batched pagination
items, err := client.ListByDataBoxEdgeDeviceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SharesClient.Refresh`

```go
ctx := context.TODO()
id := shares.NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "shareValue")

if err := client.RefreshThenPoll(ctx, id); err != nil {
	// handle the error
}
```
