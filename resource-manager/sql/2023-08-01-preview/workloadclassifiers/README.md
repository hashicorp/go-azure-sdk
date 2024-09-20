
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01-preview/workloadclassifiers` Documentation

The `workloadclassifiers` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01-preview/workloadclassifiers"
```


### Client Initialization

```go
client := workloadclassifiers.NewWorkloadClassifiersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkloadClassifiersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := workloadclassifiers.NewWorkloadClassifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "workloadGroupName", "workloadClassifierName")

payload := workloadclassifiers.WorkloadClassifier{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadClassifiersClient.Delete`

```go
ctx := context.TODO()
id := workloadclassifiers.NewWorkloadClassifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "workloadGroupName", "workloadClassifierName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadClassifiersClient.Get`

```go
ctx := context.TODO()
id := workloadclassifiers.NewWorkloadClassifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "workloadGroupName", "workloadClassifierName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadClassifiersClient.ListByWorkloadGroup`

```go
ctx := context.TODO()
id := workloadclassifiers.NewWorkloadGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "workloadGroupName")

// alternatively `client.ListByWorkloadGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByWorkloadGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
