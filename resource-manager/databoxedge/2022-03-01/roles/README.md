
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/roles` Documentation

The `roles` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2022-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/roles"
```


### Client Initialization

```go
client := roles.NewRolesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RolesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := roles.NewRoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deviceName", "name")

payload := roles.Role{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `RolesClient.Delete`

```go
ctx := context.TODO()
id := roles.NewRoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deviceName", "name")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `RolesClient.Get`

```go
ctx := context.TODO()
id := roles.NewRoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deviceName", "name")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RolesClient.ListByDataBoxEdgeDevice`

```go
ctx := context.TODO()
id := roles.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deviceName")

// alternatively `client.ListByDataBoxEdgeDevice(ctx, id)` can be used to do batched pagination
items, err := client.ListByDataBoxEdgeDeviceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
