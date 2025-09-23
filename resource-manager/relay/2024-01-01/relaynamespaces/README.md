
## `github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/relaynamespaces` Documentation

The `relaynamespaces` SDK allows for interaction with Azure Resource Manager `relay` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/relaynamespaces"
```


### Client Initialization

```go
client := relaynamespaces.NewRelayNamespacesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RelayNamespacesClient.NamespacesCreateOrUpdate`

```go
ctx := context.TODO()
id := relaynamespaces.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

payload := relaynamespaces.RelayNamespace{
	// ...
}


if err := client.NamespacesCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `RelayNamespacesClient.NamespacesDelete`

```go
ctx := context.TODO()
id := relaynamespaces.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

if err := client.NamespacesDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `RelayNamespacesClient.NamespacesGet`

```go
ctx := context.TODO()
id := relaynamespaces.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

read, err := client.NamespacesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RelayNamespacesClient.NamespacesList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.NamespacesList(ctx, id)` can be used to do batched pagination
items, err := client.NamespacesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RelayNamespacesClient.NamespacesListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.NamespacesListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.NamespacesListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RelayNamespacesClient.NamespacesUpdate`

```go
ctx := context.TODO()
id := relaynamespaces.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

payload := relaynamespaces.RelayUpdateParameters{
	// ...
}


read, err := client.NamespacesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
