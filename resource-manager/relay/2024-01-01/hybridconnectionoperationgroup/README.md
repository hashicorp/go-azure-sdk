
## `github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/hybridconnectionoperationgroup` Documentation

The `hybridconnectionoperationgroup` SDK allows for interaction with Azure Resource Manager `relay` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/hybridconnectionoperationgroup"
```


### Client Initialization

```go
client := hybridconnectionoperationgroup.NewHybridConnectionOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HybridConnectionOperationGroupClient.HybridConnectionsCreateOrUpdate`

```go
ctx := context.TODO()
id := hybridconnectionoperationgroup.NewHybridConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "hybridConnectionName")

payload := hybridconnectionoperationgroup.HybridConnection{
	// ...
}


read, err := client.HybridConnectionsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionOperationGroupClient.HybridConnectionsDelete`

```go
ctx := context.TODO()
id := hybridconnectionoperationgroup.NewHybridConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "hybridConnectionName")

read, err := client.HybridConnectionsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionOperationGroupClient.HybridConnectionsGet`

```go
ctx := context.TODO()
id := hybridconnectionoperationgroup.NewHybridConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "hybridConnectionName")

read, err := client.HybridConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionOperationGroupClient.HybridConnectionsListByNamespace`

```go
ctx := context.TODO()
id := hybridconnectionoperationgroup.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

// alternatively `client.HybridConnectionsListByNamespace(ctx, id)` can be used to do batched pagination
items, err := client.HybridConnectionsListByNamespaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
