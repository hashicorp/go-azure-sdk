
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-03-01-preview/apirevision` Documentation

The `apirevision` SDK allows for interaction with Azure Resource Manager `apimanagement` (API Version `2023-03-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-03-01-preview/apirevision"
```


### Client Initialization

```go
client := apirevision.NewApiRevisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApiRevisionClient.ListByService`

```go
ctx := context.TODO()
id := apirevision.NewApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue")

// alternatively `client.ListByService(ctx, id, apirevision.DefaultListByServiceOperationOptions())` can be used to do batched pagination
items, err := client.ListByServiceComplete(ctx, id, apirevision.DefaultListByServiceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApiRevisionClient.WorkspaceApiRevisionListByService`

```go
ctx := context.TODO()
id := apirevision.NewWorkspaceApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "apiIdValue")

// alternatively `client.WorkspaceApiRevisionListByService(ctx, id, apirevision.DefaultWorkspaceApiRevisionListByServiceOperationOptions())` can be used to do batched pagination
items, err := client.WorkspaceApiRevisionListByServiceComplete(ctx, id, apirevision.DefaultWorkspaceApiRevisionListByServiceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
