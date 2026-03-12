
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/relayserviceconnectionentityoperationgroup` Documentation

The `relayserviceconnectionentityoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/relayserviceconnectionentityoperationgroup"
```


### Client Initialization

```go
client := relayserviceconnectionentityoperationgroup.NewRelayServiceConnectionEntityOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RelayServiceConnectionEntityOperationGroupClient.WebAppsCreateOrUpdateRelayServiceConnectionSlot`

```go
ctx := context.TODO()
id := relayserviceconnectionentityoperationgroup.NewSlotHybridConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "hybridConnectionName")

payload := relayserviceconnectionentityoperationgroup.RelayServiceConnectionEntity{
	// ...
}


read, err := client.WebAppsCreateOrUpdateRelayServiceConnectionSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RelayServiceConnectionEntityOperationGroupClient.WebAppsDeleteRelayServiceConnectionSlot`

```go
ctx := context.TODO()
id := relayserviceconnectionentityoperationgroup.NewSlotHybridConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "hybridConnectionName")

read, err := client.WebAppsDeleteRelayServiceConnectionSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RelayServiceConnectionEntityOperationGroupClient.WebAppsGetRelayServiceConnectionSlot`

```go
ctx := context.TODO()
id := relayserviceconnectionentityoperationgroup.NewSlotHybridConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "hybridConnectionName")

read, err := client.WebAppsGetRelayServiceConnectionSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RelayServiceConnectionEntityOperationGroupClient.WebAppsUpdateRelayServiceConnectionSlot`

```go
ctx := context.TODO()
id := relayserviceconnectionentityoperationgroup.NewSlotHybridConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "hybridConnectionName")

payload := relayserviceconnectionentityoperationgroup.RelayServiceConnectionEntity{
	// ...
}


read, err := client.WebAppsUpdateRelayServiceConnectionSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
