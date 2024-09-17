
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/featurestoreentityversion` Documentation

The `featurestoreentityversion` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/featurestoreentityversion"
```


### Client Initialization

```go
client := featurestoreentityversion.NewFeaturestoreEntityVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FeaturestoreEntityVersionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := featurestoreentityversion.NewFeatureStoreEntityVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureStoreEntityValue", "versionValue")

payload := featurestoreentityversion.FeaturestoreEntityVersionResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `FeaturestoreEntityVersionClient.Delete`

```go
ctx := context.TODO()
id := featurestoreentityversion.NewFeatureStoreEntityVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureStoreEntityValue", "versionValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `FeaturestoreEntityVersionClient.Get`

```go
ctx := context.TODO()
id := featurestoreentityversion.NewFeatureStoreEntityVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureStoreEntityValue", "versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FeaturestoreEntityVersionClient.List`

```go
ctx := context.TODO()
id := featurestoreentityversion.NewFeatureStoreEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureStoreEntityValue")

// alternatively `client.List(ctx, id, featurestoreentityversion.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, featurestoreentityversion.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
