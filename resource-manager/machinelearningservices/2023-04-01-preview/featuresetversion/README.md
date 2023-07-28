
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/featuresetversion` Documentation

The `featuresetversion` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2023-04-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/featuresetversion"
```


### Client Initialization

```go
client := featuresetversion.NewFeaturesetVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FeaturesetVersionClient.Backfill`

```go
ctx := context.TODO()
id := featuresetversion.NewFeatureSetVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureSetValue", "versionValue")

payload := featuresetversion.FeaturesetVersionBackfillRequest{
	// ...
}


if err := client.BackfillThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `FeaturesetVersionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := featuresetversion.NewFeatureSetVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureSetValue", "versionValue")

payload := featuresetversion.FeaturesetVersionResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `FeaturesetVersionClient.Delete`

```go
ctx := context.TODO()
id := featuresetversion.NewFeatureSetVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureSetValue", "versionValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `FeaturesetVersionClient.Get`

```go
ctx := context.TODO()
id := featuresetversion.NewFeatureSetVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureSetValue", "versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FeaturesetVersionClient.List`

```go
ctx := context.TODO()
id := featuresetversion.NewFeatureSetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureSetValue")

// alternatively `client.List(ctx, id, featuresetversion.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, featuresetversion.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FeaturesetVersionClient.ListMaterializationJobs`

```go
ctx := context.TODO()
id := featuresetversion.NewFeatureSetVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "featureSetValue", "versionValue")

// alternatively `client.ListMaterializationJobs(ctx, id, featuresetversion.DefaultListMaterializationJobsOperationOptions())` can be used to do batched pagination
items, err := client.ListMaterializationJobsComplete(ctx, id, featuresetversion.DefaultListMaterializationJobsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
