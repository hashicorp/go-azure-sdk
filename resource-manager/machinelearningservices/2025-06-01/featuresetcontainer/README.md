
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2025-06-01/featuresetcontainer` Documentation

The `featuresetcontainer` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2025-06-01/featuresetcontainer"
```


### Client Initialization

```go
client := featuresetcontainer.NewFeaturesetContainerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FeaturesetContainerClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := featuresetcontainer.NewFeatureSetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "featureSetName")

payload := featuresetcontainer.FeaturesetContainerResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `FeaturesetContainerClient.Delete`

```go
ctx := context.TODO()
id := featuresetcontainer.NewFeatureSetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "featureSetName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `FeaturesetContainerClient.GetEntity`

```go
ctx := context.TODO()
id := featuresetcontainer.NewFeatureSetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "featureSetName")

read, err := client.GetEntity(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FeaturesetContainerClient.List`

```go
ctx := context.TODO()
id := featuresetcontainer.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, featuresetcontainer.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, featuresetcontainer.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
