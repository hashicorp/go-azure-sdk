
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01/feature` Documentation

The `feature` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01/feature"
```


### Client Initialization

```go
client := feature.NewFeatureClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FeatureClient.Get`

```go
ctx := context.TODO()
id := feature.NewFeatureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "featureSetName", "versionName", "featureName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FeatureClient.List`

```go
ctx := context.TODO()
id := feature.NewFeatureSetVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "featureSetName", "versionName")

// alternatively `client.List(ctx, id, feature.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, feature.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
