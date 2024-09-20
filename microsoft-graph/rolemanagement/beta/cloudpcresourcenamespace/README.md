
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcresourcenamespace` Documentation

The `cloudpcresourcenamespace` SDK allows for interaction with Microsoft Graph `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcresourcenamespace"
```


### Client Initialization

```go
client := cloudpcresourcenamespace.NewCloudPCResourceNamespaceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudPCResourceNamespaceClient.CreateCloudPCResourceNamespace`

```go
ctx := context.TODO()

payload := cloudpcresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateCloudPCResourceNamespace(ctx, payload, cloudpcresourcenamespace.DefaultCreateCloudPCResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCResourceNamespaceClient.CreateCloudPCResourceNamespaceImportResourceAction`

```go
ctx := context.TODO()
id := cloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceId")

payload := cloudpcresourcenamespace.CreateCloudPCResourceNamespaceImportResourceActionRequest{
	// ...
}


read, err := client.CreateCloudPCResourceNamespaceImportResourceAction(ctx, id, payload, cloudpcresourcenamespace.DefaultCreateCloudPCResourceNamespaceImportResourceActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCResourceNamespaceClient.DeleteCloudPCResourceNamespace`

```go
ctx := context.TODO()
id := cloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceId")

read, err := client.DeleteCloudPCResourceNamespace(ctx, id, cloudpcresourcenamespace.DefaultDeleteCloudPCResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCResourceNamespaceClient.GetCloudPCResourceNamespace`

```go
ctx := context.TODO()
id := cloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceId")

read, err := client.GetCloudPCResourceNamespace(ctx, id, cloudpcresourcenamespace.DefaultGetCloudPCResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCResourceNamespaceClient.GetCloudPCResourceNamespacesCount`

```go
ctx := context.TODO()


read, err := client.GetCloudPCResourceNamespacesCount(ctx, cloudpcresourcenamespace.DefaultGetCloudPCResourceNamespacesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCResourceNamespaceClient.ListCloudPCResourceNamespaces`

```go
ctx := context.TODO()


// alternatively `client.ListCloudPCResourceNamespaces(ctx, cloudpcresourcenamespace.DefaultListCloudPCResourceNamespacesOperationOptions())` can be used to do batched pagination
items, err := client.ListCloudPCResourceNamespacesComplete(ctx, cloudpcresourcenamespace.DefaultListCloudPCResourceNamespacesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CloudPCResourceNamespaceClient.UpdateCloudPCResourceNamespace`

```go
ctx := context.TODO()
id := cloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceId")

payload := cloudpcresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateCloudPCResourceNamespace(ctx, id, payload, cloudpcresourcenamespace.DefaultUpdateCloudPCResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
