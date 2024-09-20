
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowuserattributeassignment` Documentation

The `b2xuserflowuserattributeassignment` SDK allows for interaction with Microsoft Graph `identity` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowuserattributeassignment"
```


### Client Initialization

```go
client := b2xuserflowuserattributeassignment.NewB2xUserFlowUserAttributeAssignmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `B2xUserFlowUserAttributeAssignmentClient.CreateB2xUserFlowUserAttributeAssignment`

```go
ctx := context.TODO()
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowID("b2xIdentityUserFlowId")

payload := b2xuserflowuserattributeassignment.IdentityUserFlowAttributeAssignment{
	// ...
}


read, err := client.CreateB2xUserFlowUserAttributeAssignment(ctx, id, payload, b2xuserflowuserattributeassignment.DefaultCreateB2xUserFlowUserAttributeAssignmentOperationOptions())
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
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowIdUserAttributeAssignmentID("b2xIdentityUserFlowId", "identityUserFlowAttributeAssignmentId")

read, err := client.DeleteB2xUserFlowUserAttributeAssignment(ctx, id, b2xuserflowuserattributeassignment.DefaultDeleteB2xUserFlowUserAttributeAssignmentOperationOptions())
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
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowIdUserAttributeAssignmentID("b2xIdentityUserFlowId", "identityUserFlowAttributeAssignmentId")

read, err := client.GetB2xUserFlowUserAttributeAssignment(ctx, id, b2xuserflowuserattributeassignment.DefaultGetB2xUserFlowUserAttributeAssignmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowUserAttributeAssignmentClient.GetB2xUserFlowUserAttributeAssignmentsCount`

```go
ctx := context.TODO()
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowID("b2xIdentityUserFlowId")

read, err := client.GetB2xUserFlowUserAttributeAssignmentsCount(ctx, id, b2xuserflowuserattributeassignment.DefaultGetB2xUserFlowUserAttributeAssignmentsCountOperationOptions())
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
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowID("b2xIdentityUserFlowId")

// alternatively `client.ListB2xUserFlowUserAttributeAssignments(ctx, id, b2xuserflowuserattributeassignment.DefaultListB2xUserFlowUserAttributeAssignmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListB2xUserFlowUserAttributeAssignmentsComplete(ctx, id, b2xuserflowuserattributeassignment.DefaultListB2xUserFlowUserAttributeAssignmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `B2xUserFlowUserAttributeAssignmentClient.SetB2xUserFlowUserAttributeAssignmentsOrder`

```go
ctx := context.TODO()
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowID("b2xIdentityUserFlowId")

payload := b2xuserflowuserattributeassignment.SetB2xUserFlowUserAttributeAssignmentsOrderRequest{
	// ...
}


read, err := client.SetB2xUserFlowUserAttributeAssignmentsOrder(ctx, id, payload, b2xuserflowuserattributeassignment.DefaultSetB2xUserFlowUserAttributeAssignmentsOrderOperationOptions())
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
id := b2xuserflowuserattributeassignment.NewIdentityB2xUserFlowIdUserAttributeAssignmentID("b2xIdentityUserFlowId", "identityUserFlowAttributeAssignmentId")

payload := b2xuserflowuserattributeassignment.IdentityUserFlowAttributeAssignment{
	// ...
}


read, err := client.UpdateB2xUserFlowUserAttributeAssignment(ctx, id, payload, b2xuserflowuserattributeassignment.DefaultUpdateB2xUserFlowUserAttributeAssignmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
