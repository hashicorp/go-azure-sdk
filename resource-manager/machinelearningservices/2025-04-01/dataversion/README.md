
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2025-04-01/dataversion` Documentation

The `dataversion` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2025-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2025-04-01/dataversion"
```


### Client Initialization

```go
client := dataversion.NewDataVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataVersionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := dataversion.NewDataVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "dataName", "versionName")

payload := dataversion.DataVersionBaseResource{
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


### Example Usage: `DataVersionClient.Delete`

```go
ctx := context.TODO()
id := dataversion.NewDataVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "dataName", "versionName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataVersionClient.Get`

```go
ctx := context.TODO()
id := dataversion.NewDataVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "dataName", "versionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataVersionClient.List`

```go
ctx := context.TODO()
id := dataversion.NewWorkspaceDataID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "dataName")

// alternatively `client.List(ctx, id, dataversion.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, dataversion.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataVersionClient.Publish`

```go
ctx := context.TODO()
id := dataversion.NewDataVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "dataName", "versionName")

payload := dataversion.DestinationAsset{
	// ...
}


if err := client.PublishThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
