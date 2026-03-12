
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hybridconnections` Documentation

The `hybridconnections` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hybridconnections"
```


### Client Initialization

```go
client := hybridconnections.NewHybridConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HybridConnectionsClient.AppServicePlansDeleteHybridConnection`

```go
ctx := context.TODO()
id := hybridconnections.NewHybridConnectionNamespaceRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "hybridConnectionNamespaceName", "relayName")

read, err := client.AppServicePlansDeleteHybridConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionsClient.AppServicePlansGetHybridConnection`

```go
ctx := context.TODO()
id := hybridconnections.NewHybridConnectionNamespaceRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "hybridConnectionNamespaceName", "relayName")

read, err := client.AppServicePlansGetHybridConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionsClient.AppServicePlansListHybridConnectionKeys`

```go
ctx := context.TODO()
id := hybridconnections.NewHybridConnectionNamespaceRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "hybridConnectionNamespaceName", "relayName")

read, err := client.AppServicePlansListHybridConnectionKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionsClient.AppServicePlansListWebAppsByHybridConnection`

```go
ctx := context.TODO()
id := hybridconnections.NewHybridConnectionNamespaceRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "hybridConnectionNamespaceName", "relayName")

// alternatively `client.AppServicePlansListWebAppsByHybridConnection(ctx, id)` can be used to do batched pagination
items, err := client.AppServicePlansListWebAppsByHybridConnectionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
