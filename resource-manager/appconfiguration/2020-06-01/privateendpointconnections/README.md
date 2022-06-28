
## `github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2020-06-01/privateendpointconnections` Documentation

The `privateendpointconnections` SDK allows for interaction with the Azure Resource Manager Service `appconfiguration` (API Version `2020-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2020-06-01/privateendpointconnections"
```


### Client Initialization

```go
client := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configStoreValue", "privateEndpointConnectionValue")

payload := privateendpointconnections.PrivateEndpointConnection{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.Delete`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configStoreValue", "privateEndpointConnectionValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.Get`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configStoreValue", "privateEndpointConnectionValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ListByConfigurationStore`

```go
ctx := context.TODO()
id := privateendpointconnections.NewConfigurationStoreID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configStoreValue")
// alternatively `client.ListByConfigurationStore(ctx, id)` can be used to do batched pagination
items, err := client.ListByConfigurationStoreComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
