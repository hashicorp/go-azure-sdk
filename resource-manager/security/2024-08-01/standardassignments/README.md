
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2024-08-01/standardassignments` Documentation

The `standardassignments` SDK allows for interaction with Azure Resource Manager `security` (API Version `2024-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2024-08-01/standardassignments"
```


### Client Initialization

```go
client := standardassignments.NewStandardAssignmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StandardAssignmentsClient.Create`

```go
ctx := context.TODO()
id := standardassignments.NewScopedStandardAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "standardAssignmentName")

payload := standardassignments.StandardAssignment{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardAssignmentsClient.Delete`

```go
ctx := context.TODO()
id := standardassignments.NewScopedStandardAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "standardAssignmentName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardAssignmentsClient.Get`

```go
ctx := context.TODO()
id := standardassignments.NewScopedStandardAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "standardAssignmentName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardAssignmentsClient.List`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
