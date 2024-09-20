
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementresourcenamespace` Documentation

The `devicemanagementresourcenamespace` SDK allows for interaction with Microsoft Graph `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementresourcenamespace"
```


### Client Initialization

```go
client := devicemanagementresourcenamespace.NewDeviceManagementResourceNamespaceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceManagementResourceNamespaceClient.CreateDeviceManagementResourceNamespace`

```go
ctx := context.TODO()

payload := devicemanagementresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateDeviceManagementResourceNamespace(ctx, payload, devicemanagementresourcenamespace.DefaultCreateDeviceManagementResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementResourceNamespaceClient.CreateDeviceManagementResourceNamespaceImportResourceAction`

```go
ctx := context.TODO()
id := devicemanagementresourcenamespace.NewRoleManagementDeviceManagementResourceNamespaceID("unifiedRbacResourceNamespaceId")

payload := devicemanagementresourcenamespace.CreateDeviceManagementResourceNamespaceImportResourceActionRequest{
	// ...
}


read, err := client.CreateDeviceManagementResourceNamespaceImportResourceAction(ctx, id, payload, devicemanagementresourcenamespace.DefaultCreateDeviceManagementResourceNamespaceImportResourceActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementResourceNamespaceClient.DeleteDeviceManagementResourceNamespace`

```go
ctx := context.TODO()
id := devicemanagementresourcenamespace.NewRoleManagementDeviceManagementResourceNamespaceID("unifiedRbacResourceNamespaceId")

read, err := client.DeleteDeviceManagementResourceNamespace(ctx, id, devicemanagementresourcenamespace.DefaultDeleteDeviceManagementResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementResourceNamespaceClient.GetDeviceManagementResourceNamespace`

```go
ctx := context.TODO()
id := devicemanagementresourcenamespace.NewRoleManagementDeviceManagementResourceNamespaceID("unifiedRbacResourceNamespaceId")

read, err := client.GetDeviceManagementResourceNamespace(ctx, id, devicemanagementresourcenamespace.DefaultGetDeviceManagementResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementResourceNamespaceClient.GetDeviceManagementResourceNamespacesCount`

```go
ctx := context.TODO()


read, err := client.GetDeviceManagementResourceNamespacesCount(ctx, devicemanagementresourcenamespace.DefaultGetDeviceManagementResourceNamespacesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementResourceNamespaceClient.ListDeviceManagementResourceNamespaces`

```go
ctx := context.TODO()


// alternatively `client.ListDeviceManagementResourceNamespaces(ctx, devicemanagementresourcenamespace.DefaultListDeviceManagementResourceNamespacesOperationOptions())` can be used to do batched pagination
items, err := client.ListDeviceManagementResourceNamespacesComplete(ctx, devicemanagementresourcenamespace.DefaultListDeviceManagementResourceNamespacesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceManagementResourceNamespaceClient.UpdateDeviceManagementResourceNamespace`

```go
ctx := context.TODO()
id := devicemanagementresourcenamespace.NewRoleManagementDeviceManagementResourceNamespaceID("unifiedRbacResourceNamespaceId")

payload := devicemanagementresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateDeviceManagementResourceNamespace(ctx, id, payload, devicemanagementresourcenamespace.DefaultUpdateDeviceManagementResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
