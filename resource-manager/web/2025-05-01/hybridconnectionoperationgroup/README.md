
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hybridconnectionoperationgroup` Documentation

The `hybridconnectionoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hybridconnectionoperationgroup"
```


### Client Initialization

```go
client := hybridconnectionoperationgroup.NewHybridConnectionOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HybridConnectionOperationGroupClient.WebAppsCreateOrUpdateHybridConnection`

```go
ctx := context.TODO()
id := hybridconnectionoperationgroup.NewRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "hybridConnectionNamespaceName", "relayName")

payload := hybridconnectionoperationgroup.HybridConnection{
	// ...
}


read, err := client.WebAppsCreateOrUpdateHybridConnection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionOperationGroupClient.WebAppsDeleteHybridConnection`

```go
ctx := context.TODO()
id := hybridconnectionoperationgroup.NewRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "hybridConnectionNamespaceName", "relayName")

read, err := client.WebAppsDeleteHybridConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionOperationGroupClient.WebAppsGetHybridConnection`

```go
ctx := context.TODO()
id := hybridconnectionoperationgroup.NewRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "hybridConnectionNamespaceName", "relayName")

read, err := client.WebAppsGetHybridConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionOperationGroupClient.WebAppsUpdateHybridConnection`

```go
ctx := context.TODO()
id := hybridconnectionoperationgroup.NewRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "hybridConnectionNamespaceName", "relayName")

payload := hybridconnectionoperationgroup.HybridConnection{
	// ...
}


read, err := client.WebAppsUpdateHybridConnection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
