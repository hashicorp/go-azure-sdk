
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/workloadgroups` Documentation

The `workloadgroups` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2023-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/workloadgroups"
```


### Client Initialization

```go
client := workloadgroups.NewWorkloadGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkloadGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := workloadgroups.NewWorkloadGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "workloadGroupValue")

payload := workloadgroups.WorkloadGroup{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadGroupsClient.Delete`

```go
ctx := context.TODO()
id := workloadgroups.NewWorkloadGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "workloadGroupValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadGroupsClient.Get`

```go
ctx := context.TODO()
id := workloadgroups.NewWorkloadGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "workloadGroupValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadGroupsClient.ListByDatabase`

```go
ctx := context.TODO()
id := commonids.NewSqlDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

// alternatively `client.ListByDatabase(ctx, id)` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
