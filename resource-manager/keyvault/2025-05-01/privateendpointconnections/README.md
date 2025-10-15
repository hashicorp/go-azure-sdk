
## `github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/privateendpointconnections` Documentation

The `privateendpointconnections` SDK allows for interaction with Azure Resource Manager `keyvault` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/privateendpointconnections"
```


### Client Initialization

```go
client := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionsClient.Delete`

```go
ctx := context.TODO()
id := commonids.NewKeyVaultPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "privateEndpointConnectionName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.Get`

```go
ctx := context.TODO()
id := commonids.NewKeyVaultPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "privateEndpointConnectionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ListByResource`

```go
ctx := context.TODO()
id := commonids.NewKeyVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.ListByResource(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateEndpointConnectionsClient.Put`

```go
ctx := context.TODO()
id := commonids.NewKeyVaultPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "privateEndpointConnectionName")

payload := privateendpointconnections.PrivateEndpointConnection{
	// ...
}


read, err := client.Put(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
