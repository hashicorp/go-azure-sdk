
## `github.com/hashicorp/go-azure-sdk/resource-manager/identity/v1.0/identityb2xuserflowuserattributeassignment` Documentation

The `identityb2xuserflowuserattributeassignment` SDK allows for interaction with the Azure Resource Manager Service `identity` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/identity/v1.0/identityb2xuserflowuserattributeassignment"
```


### Client Initialization

```go
client := identityb2xuserflowuserattributeassignment.NewIdentityB2xUserFlowUserAttributeAssignmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IdentityB2xUserFlowUserAttributeAssignmentClient.CreateIdentityB2xUserFlowByIdUserAttributeAssignment`

```go
ctx := context.TODO()
id := identityb2xuserflowuserattributeassignment.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := identityb2xuserflowuserattributeassignment.IdentityUserFlowAttributeAssignment{
	// ...
}


read, err := client.CreateIdentityB2xUserFlowByIdUserAttributeAssignment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowUserAttributeAssignmentClient.DeleteIdentityB2xUserFlowByIdUserAttributeAssignmentById`

```go
ctx := context.TODO()
id := identityb2xuserflowuserattributeassignment.NewIdentityB2xUserFlowUserAttributeAssignmentID("b2xIdentityUserFlowIdValue", "identityUserFlowAttributeAssignmentIdValue")

read, err := client.DeleteIdentityB2xUserFlowByIdUserAttributeAssignmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowUserAttributeAssignmentClient.GetIdentityB2xUserFlowByIdUserAttributeAssignmentById`

```go
ctx := context.TODO()
id := identityb2xuserflowuserattributeassignment.NewIdentityB2xUserFlowUserAttributeAssignmentID("b2xIdentityUserFlowIdValue", "identityUserFlowAttributeAssignmentIdValue")

read, err := client.GetIdentityB2xUserFlowByIdUserAttributeAssignmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowUserAttributeAssignmentClient.GetIdentityB2xUserFlowByIdUserAttributeAssignmentCount`

```go
ctx := context.TODO()
id := identityb2xuserflowuserattributeassignment.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.GetIdentityB2xUserFlowByIdUserAttributeAssignmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowUserAttributeAssignmentClient.ListIdentityB2xUserFlowByIdUserAttributeAssignments`

```go
ctx := context.TODO()
id := identityb2xuserflowuserattributeassignment.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

// alternatively `client.ListIdentityB2xUserFlowByIdUserAttributeAssignments(ctx, id)` can be used to do batched pagination
items, err := client.ListIdentityB2xUserFlowByIdUserAttributeAssignmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IdentityB2xUserFlowUserAttributeAssignmentClient.SetIdentityB2xUserFlowByIdUserAttributeAssignmentsOrder`

```go
ctx := context.TODO()
id := identityb2xuserflowuserattributeassignment.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := identityb2xuserflowuserattributeassignment.SetIdentityB2xUserFlowByIdUserAttributeAssignmentsOrderRequest{
	// ...
}


read, err := client.SetIdentityB2xUserFlowByIdUserAttributeAssignmentsOrder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowUserAttributeAssignmentClient.UpdateIdentityB2xUserFlowByIdUserAttributeAssignmentById`

```go
ctx := context.TODO()
id := identityb2xuserflowuserattributeassignment.NewIdentityB2xUserFlowUserAttributeAssignmentID("b2xIdentityUserFlowIdValue", "identityUserFlowAttributeAssignmentIdValue")

payload := identityb2xuserflowuserattributeassignment.IdentityUserFlowAttributeAssignment{
	// ...
}


read, err := client.UpdateIdentityB2xUserFlowByIdUserAttributeAssignmentById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
