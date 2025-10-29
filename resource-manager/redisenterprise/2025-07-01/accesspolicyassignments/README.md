
## `github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2025-07-01/accesspolicyassignments` Documentation

The `accesspolicyassignments` SDK allows for interaction with Azure Resource Manager `redisenterprise` (API Version `2025-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2025-07-01/accesspolicyassignments"
```


### Client Initialization

```go
client := accesspolicyassignments.NewAccessPolicyAssignmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessPolicyAssignmentsClient.AccessPolicyAssignmentCreateUpdate`

```go
ctx := context.TODO()
id := accesspolicyassignments.NewAccessPolicyAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "redisEnterpriseName", "databaseName", "accessPolicyAssignmentName")

payload := accesspolicyassignments.AccessPolicyAssignment{
	// ...
}


if err := client.AccessPolicyAssignmentCreateUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AccessPolicyAssignmentsClient.AccessPolicyAssignmentDelete`

```go
ctx := context.TODO()
id := accesspolicyassignments.NewAccessPolicyAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "redisEnterpriseName", "databaseName", "accessPolicyAssignmentName")

if err := client.AccessPolicyAssignmentDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AccessPolicyAssignmentsClient.AccessPolicyAssignmentGet`

```go
ctx := context.TODO()
id := accesspolicyassignments.NewAccessPolicyAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "redisEnterpriseName", "databaseName", "accessPolicyAssignmentName")

read, err := client.AccessPolicyAssignmentGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessPolicyAssignmentsClient.AccessPolicyAssignmentList`

```go
ctx := context.TODO()
id := accesspolicyassignments.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "redisEnterpriseName", "databaseName")

// alternatively `client.AccessPolicyAssignmentList(ctx, id)` can be used to do batched pagination
items, err := client.AccessPolicyAssignmentListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
