
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/privatelinkresources` Documentation

The `privatelinkresources` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2026-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/privatelinkresources"
```


### Client Initialization

```go
client := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkResourcesClient.ContainerAppPrivateLinkResourcesGet`

```go
ctx := context.TODO()
id := privatelinkresources.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "privateLinkResourceName")

read, err := client.ContainerAppPrivateLinkResourcesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkResourcesClient.ContainerAppPrivateLinkResourcesList`

```go
ctx := context.TODO()
id := privatelinkresources.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName")

// alternatively `client.ContainerAppPrivateLinkResourcesList(ctx, id)` can be used to do batched pagination
items, err := client.ContainerAppPrivateLinkResourcesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateLinkResourcesClient.ManagedEnvironmentPrivateLinkResourcesGet`

```go
ctx := context.TODO()
id := privatelinkresources.NewManagedEnvironmentPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentName", "privateLinkResourceName")

read, err := client.ManagedEnvironmentPrivateLinkResourcesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkResourcesClient.ManagedEnvironmentPrivateLinkResourcesList`

```go
ctx := context.TODO()
id := privatelinkresources.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentName")

// alternatively `client.ManagedEnvironmentPrivateLinkResourcesList(ctx, id)` can be used to do batched pagination
items, err := client.ManagedEnvironmentPrivateLinkResourcesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
