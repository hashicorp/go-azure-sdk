
## `github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/sessionhostmanagement` Documentation

The `sessionhostmanagement` SDK allows for interaction with Azure Resource Manager `desktopvirtualization` (API Version `2024-04-08-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/sessionhostmanagement"
```


### Client Initialization

```go
client := sessionhostmanagement.NewSessionHostManagementClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SessionHostManagementClient.ControlSessionHostUpdatePost`

```go
ctx := context.TODO()
id := sessionhostmanagement.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName")

payload := sessionhostmanagement.HostPoolUpdateControlParameter{
	// ...
}


if err := client.ControlSessionHostUpdatePostThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SessionHostManagementClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := sessionhostmanagement.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName")

payload := sessionhostmanagement.SessionHostManagement{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SessionHostManagementClient.InitiateSessionHostUpdatePost`

```go
ctx := context.TODO()
id := sessionhostmanagement.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName")

payload := sessionhostmanagement.UpdateSessionHostsRequestBody{
	// ...
}


read, err := client.InitiateSessionHostUpdatePost(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SessionHostManagementClient.Update`

```go
ctx := context.TODO()
id := sessionhostmanagement.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName")

payload := sessionhostmanagement.SessionHostManagementPatch{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SessionHostManagementClient.UpdateStatusGet`

```go
ctx := context.TODO()
id := sessionhostmanagement.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName")

read, err := client.UpdateStatusGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
