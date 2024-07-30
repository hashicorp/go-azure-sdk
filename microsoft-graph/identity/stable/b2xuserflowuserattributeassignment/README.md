
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowuserattributeassignment` Documentation

The `b2xuserflowuserattributeassignment` SDK allows for interaction with the Azure Resource Manager Service `identity` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowuserattributeassignment"
```


### Client Initialization

```go
client := b2xuserflowuserattributeassignment.NewB2xUserFlowUserAttributeAssignmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `B2xUserFlowUserAttributeAssignmentClient.CreateB2xUserFlowUserAttributeAssignment`

```go
ctx := context.TODO()
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := b2xuserflowuserattributeassignment.IdentityUserFlowAttributeAssignment{
	// ...
}


read, err := client.CreateB2xUserFlowUserAttributeAssignment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowUserAttributeAssignmentClient.DeleteB2xUserFlowUserAttributeAssignment`

```go
ctx := context.TODO()
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowIdUserAttributeAssignmentID("b2xIdentityUserFlowIdValue", "identityUserFlowAttributeAssignmentIdValue")

read, err := client.DeleteB2xUserFlowUserAttributeAssignment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowUserAttributeAssignmentClient.GetB2xUserFlowUserAttributeAssignment`

```go
ctx := context.TODO()
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowIdUserAttributeAssignmentID("b2xIdentityUserFlowIdValue", "identityUserFlowAttributeAssignmentIdValue")

read, err := client.GetB2xUserFlowUserAttributeAssignment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowUserAttributeAssignmentClient.GetB2xUserFlowUserAttributeAssignmentCount`

```go
ctx := context.TODO()
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.GetB2xUserFlowUserAttributeAssignmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowUserAttributeAssignmentClient.ListB2xUserFlowUserAttributeAssignments`

```go
ctx := context.TODO()
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

// alternatively `client.ListB2xUserFlowUserAttributeAssignments(ctx, id)` can be used to do batched pagination
items, err := client.ListB2xUserFlowUserAttributeAssignmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `B2xUserFlowUserAttributeAssignmentClient.SetIdentityB2xUserFlowUserAttributeAssignmentsOrder`

```go
ctx := context.TODO()
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := b2xuserflowuserattributeassignment.SetIdentityB2xUserFlowUserAttributeAssignmentsOrderRequest{
	// ...
}


read, err := client.SetIdentityB2xUserFlowUserAttributeAssignmentsOrder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowUserAttributeAssignmentClient.UpdateB2xUserFlowUserAttributeAssignment`

```go
ctx := context.TODO()
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowIdUserAttributeAssignmentID("b2xIdentityUserFlowIdValue", "identityUserFlowAttributeAssignmentIdValue")

payload := b2xuserflowuserattributeassignment.IdentityUserFlowAttributeAssignment{
	// ...
}


read, err := client.UpdateB2xUserFlowUserAttributeAssignment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
