
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/relayserviceconnectionentities` Documentation

The `relayserviceconnectionentities` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/relayserviceconnectionentities"
```


### Client Initialization

```go
client := relayserviceconnectionentities.NewRelayServiceConnectionEntitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RelayServiceConnectionEntitiesClient.WebAppsCreateOrUpdateRelayServiceConnection`

```go
ctx := context.TODO()
id := relayserviceconnectionentities.NewHybridConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "hybridConnectionName")

payload := relayserviceconnectionentities.RelayServiceConnectionEntity{
	// ...
}


read, err := client.WebAppsCreateOrUpdateRelayServiceConnection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RelayServiceConnectionEntitiesClient.WebAppsDeleteRelayServiceConnection`

```go
ctx := context.TODO()
id := relayserviceconnectionentities.NewHybridConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "hybridConnectionName")

read, err := client.WebAppsDeleteRelayServiceConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RelayServiceConnectionEntitiesClient.WebAppsGetRelayServiceConnection`

```go
ctx := context.TODO()
id := relayserviceconnectionentities.NewHybridConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "hybridConnectionName")

read, err := client.WebAppsGetRelayServiceConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RelayServiceConnectionEntitiesClient.WebAppsUpdateRelayServiceConnection`

```go
ctx := context.TODO()
id := relayserviceconnectionentities.NewHybridConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "hybridConnectionName")

payload := relayserviceconnectionentities.RelayServiceConnectionEntity{
	// ...
}


read, err := client.WebAppsUpdateRelayServiceConnection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
