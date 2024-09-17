
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-03-01-preview/workspacepolicy` Documentation

The `workspacepolicy` SDK allows for interaction with Azure Resource Manager `apimanagement` (API Version `2023-03-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-03-01-preview/workspacepolicy"
```


### Client Initialization

```go
client := workspacepolicy.NewWorkspacePolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspacePolicyClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := workspacepolicy.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue")

payload := workspacepolicy.PolicyContract{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, workspacepolicy.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspacePolicyClient.Delete`

```go
ctx := context.TODO()
id := workspacepolicy.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue")

read, err := client.Delete(ctx, id, workspacepolicy.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspacePolicyClient.Get`

```go
ctx := context.TODO()
id := workspacepolicy.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue")

read, err := client.Get(ctx, id, workspacepolicy.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspacePolicyClient.GetEntityTag`

```go
ctx := context.TODO()
id := workspacepolicy.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue")

read, err := client.GetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspacePolicyClient.ListByApi`

```go
ctx := context.TODO()
id := workspacepolicy.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue")

// alternatively `client.ListByApi(ctx, id)` can be used to do batched pagination
items, err := client.ListByApiComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
