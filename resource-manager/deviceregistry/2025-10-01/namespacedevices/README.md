
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/namespacedevices` Documentation

The `namespacedevices` SDK allows for interaction with Azure Resource Manager `deviceregistry` (API Version `2025-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/namespacedevices"
```


### Client Initialization

```go
client := namespacedevices.NewNamespaceDevicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NamespaceDevicesClient.CreateOrReplace`

```go
ctx := context.TODO()
id := namespacedevices.NewDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "deviceName")

payload := namespacedevices.NamespaceDevice{
	// ...
}


if err := client.CreateOrReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NamespaceDevicesClient.Delete`

```go
ctx := context.TODO()
id := namespacedevices.NewDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "deviceName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NamespaceDevicesClient.Get`

```go
ctx := context.TODO()
id := namespacedevices.NewDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "deviceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespaceDevicesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := namespacedevices.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NamespaceDevicesClient.Update`

```go
ctx := context.TODO()
id := namespacedevices.NewDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "deviceName")

payload := namespacedevices.NamespaceDeviceUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
