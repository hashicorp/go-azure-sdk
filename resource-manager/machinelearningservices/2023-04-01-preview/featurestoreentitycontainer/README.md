
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/featurestoreentitycontainer` Documentation

The `featurestoreentitycontainer` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2023-04-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/featurestoreentitycontainer"
```


### Client Initialization

```go
client := featurestoreentitycontainer.NewFeaturestoreEntityContainerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FeaturestoreEntityContainerClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := featurestoreentitycontainer.NewFeatureStoreEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureStoreEntityValue")

payload := featurestoreentitycontainer.FeaturestoreEntityContainerResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `FeaturestoreEntityContainerClient.Delete`

```go
ctx := context.TODO()
id := featurestoreentitycontainer.NewFeatureStoreEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureStoreEntityValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `FeaturestoreEntityContainerClient.GetEntity`

```go
ctx := context.TODO()
id := featurestoreentitycontainer.NewFeatureStoreEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureStoreEntityValue")

read, err := client.GetEntity(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FeaturestoreEntityContainerClient.List`

```go
ctx := context.TODO()
id := featurestoreentitycontainer.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.List(ctx, id, featurestoreentitycontainer.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, featurestoreentitycontainer.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
