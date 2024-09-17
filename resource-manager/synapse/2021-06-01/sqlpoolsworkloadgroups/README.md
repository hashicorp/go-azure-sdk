
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsworkloadgroups` Documentation

The `sqlpoolsworkloadgroups` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsworkloadgroups"
```


### Client Initialization

```go
client := sqlpoolsworkloadgroups.NewSqlPoolsWorkloadGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsWorkloadGroupsClient.SqlPoolWorkloadGroupCreateOrUpdate`

```go
ctx := context.TODO()
id := sqlpoolsworkloadgroups.NewWorkloadGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "workloadGroupValue")

payload := sqlpoolsworkloadgroups.WorkloadGroup{
	// ...
}


if err := client.SqlPoolWorkloadGroupCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SqlPoolsWorkloadGroupsClient.SqlPoolWorkloadGroupDelete`

```go
ctx := context.TODO()
id := sqlpoolsworkloadgroups.NewWorkloadGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "workloadGroupValue")

if err := client.SqlPoolWorkloadGroupDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SqlPoolsWorkloadGroupsClient.SqlPoolWorkloadGroupGet`

```go
ctx := context.TODO()
id := sqlpoolsworkloadgroups.NewWorkloadGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "workloadGroupValue")

read, err := client.SqlPoolWorkloadGroupGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsWorkloadGroupsClient.SqlPoolWorkloadGroupList`

```go
ctx := context.TODO()
id := sqlpoolsworkloadgroups.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

// alternatively `client.SqlPoolWorkloadGroupList(ctx, id)` can be used to do batched pagination
items, err := client.SqlPoolWorkloadGroupListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
