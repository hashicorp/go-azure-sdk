
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/containerappslabelhistory` Documentation

The `containerappslabelhistory` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2026-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/containerappslabelhistory"
```


### Client Initialization

```go
client := containerappslabelhistory.NewContainerAppsLabelHistoryClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContainerAppsLabelHistoryClient.DeleteLabelHistory`

```go
ctx := context.TODO()
id := containerappslabelhistory.NewLabelHistoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "labelHistoryName")

read, err := client.DeleteLabelHistory(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContainerAppsLabelHistoryClient.GetLabelHistory`

```go
ctx := context.TODO()
id := containerappslabelhistory.NewLabelHistoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "labelHistoryName")

read, err := client.GetLabelHistory(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContainerAppsLabelHistoryClient.ListLabelHistory`

```go
ctx := context.TODO()
id := containerappslabelhistory.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName")

// alternatively `client.ListLabelHistory(ctx, id, containerappslabelhistory.DefaultListLabelHistoryOperationOptions())` can be used to do batched pagination
items, err := client.ListLabelHistoryComplete(ctx, id, containerappslabelhistory.DefaultListLabelHistoryOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
