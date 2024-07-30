
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcresourcenamespace` Documentation

The `cloudpcresourcenamespace` SDK allows for interaction with the Azure Resource Manager Service `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcresourcenamespace"
```


### Client Initialization

```go
client := cloudpcresourcenamespace.NewCloudPCResourceNamespaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudPCResourceNamespaceClient.CreateCloudPCResourceNamespace`

```go
ctx := context.TODO()

payload := cloudpcresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateCloudPCResourceNamespace(ctx, payload)
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
id := cloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := cloudpcresourcenamespace.CreateCloudPCResourceNamespaceImportResourceActionRequest{
	// ...
}


read, err := client.CreateCloudPCResourceNamespaceImportResourceAction(ctx, id, payload)
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
id := cloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.DeleteCloudPCResourceNamespace(ctx, id)
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
id := cloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.GetCloudPCResourceNamespace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudPCResourceNamespaceClient.GetCloudPCResourceNamespaceCount`

```go
ctx := context.TODO()


read, err := client.GetCloudPCResourceNamespaceCount(ctx)
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


// alternatively `client.ListCloudPCResourceNamespaces(ctx)` can be used to do batched pagination
items, err := client.ListCloudPCResourceNamespacesComplete(ctx)
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
id := cloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := cloudpcresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateCloudPCResourceNamespace(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
