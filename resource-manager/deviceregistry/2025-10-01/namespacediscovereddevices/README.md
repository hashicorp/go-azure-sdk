
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/namespacediscovereddevices` Documentation

The `namespacediscovereddevices` SDK allows for interaction with Azure Resource Manager `deviceregistry` (API Version `2025-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/namespacediscovereddevices"
```


### Client Initialization

```go
client := namespacediscovereddevices.NewNamespaceDiscoveredDevicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NamespaceDiscoveredDevicesClient.CreateOrReplace`

```go
ctx := context.TODO()
id := namespacediscovereddevices.NewDiscoveredDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "discoveredDeviceName")

payload := namespacediscovereddevices.NamespaceDiscoveredDevice{
	// ...
}


if err := client.CreateOrReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NamespaceDiscoveredDevicesClient.Delete`

```go
ctx := context.TODO()
id := namespacediscovereddevices.NewDiscoveredDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "discoveredDeviceName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NamespaceDiscoveredDevicesClient.Get`

```go
ctx := context.TODO()
id := namespacediscovereddevices.NewDiscoveredDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "discoveredDeviceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespaceDiscoveredDevicesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := namespacediscovereddevices.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NamespaceDiscoveredDevicesClient.Update`

```go
ctx := context.TODO()
id := namespacediscovereddevices.NewDiscoveredDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "discoveredDeviceName")

payload := namespacediscovereddevices.NamespaceDiscoveredDeviceUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
