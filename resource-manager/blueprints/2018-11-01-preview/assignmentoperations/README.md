
## `github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/assignmentoperations` Documentation

The `assignmentoperations` SDK allows for interaction with Azure Resource Manager `blueprints` (API Version `2018-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/assignmentoperations"
```


### Client Initialization

```go
client := assignmentoperations.NewAssignmentOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AssignmentOperationsClient.Get`

```go
ctx := context.TODO()
id := assignmentoperations.NewScopedAssignmentOperationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "assignmentName", "assignmentOperationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentOperationsClient.List`

```go
ctx := context.TODO()
id := assignmentoperations.NewScopedBlueprintAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "assignmentName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
