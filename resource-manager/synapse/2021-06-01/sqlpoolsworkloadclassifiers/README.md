
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsworkloadclassifiers` Documentation

The `sqlpoolsworkloadclassifiers` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsworkloadclassifiers"
```


### Client Initialization

```go
client := sqlpoolsworkloadclassifiers.NewSqlPoolsWorkloadClassifiersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsWorkloadClassifiersClient.SqlPoolWorkloadClassifierCreateOrUpdate`

```go
ctx := context.TODO()
id := sqlpoolsworkloadclassifiers.NewWorkloadClassifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "workloadGroupValue", "workloadClassifierValue")

payload := sqlpoolsworkloadclassifiers.WorkloadClassifier{
	// ...
}


if err := client.SqlPoolWorkloadClassifierCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SqlPoolsWorkloadClassifiersClient.SqlPoolWorkloadClassifierDelete`

```go
ctx := context.TODO()
id := sqlpoolsworkloadclassifiers.NewWorkloadClassifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "workloadGroupValue", "workloadClassifierValue")

if err := client.SqlPoolWorkloadClassifierDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SqlPoolsWorkloadClassifiersClient.SqlPoolWorkloadClassifierGet`

```go
ctx := context.TODO()
id := sqlpoolsworkloadclassifiers.NewWorkloadClassifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "workloadGroupValue", "workloadClassifierValue")

read, err := client.SqlPoolWorkloadClassifierGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsWorkloadClassifiersClient.SqlPoolWorkloadClassifierList`

```go
ctx := context.TODO()
id := sqlpoolsworkloadclassifiers.NewWorkloadGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "workloadGroupValue")

// alternatively `client.SqlPoolWorkloadClassifierList(ctx, id)` can be used to do batched pagination
items, err := client.SqlPoolWorkloadClassifierListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
