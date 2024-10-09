
## `github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/blueprintassignments` Documentation

The `blueprintassignments` SDK allows for interaction with Azure Resource Manager `blueprints` (API Version `2018-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/blueprintassignments"
```


### Client Initialization

```go
client := blueprintassignments.NewBlueprintAssignmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BlueprintAssignmentsClient.AssignmentsWhoIsBlueprint`

```go
ctx := context.TODO()
id := blueprintassignments.NewScopedBlueprintAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "blueprintAssignmentName")

read, err := client.AssignmentsWhoIsBlueprint(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
