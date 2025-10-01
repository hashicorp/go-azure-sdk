
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/privateendpointconnectionresources` Documentation

The `privateendpointconnectionresources` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/privateendpointconnectionresources"
```


### Client Initialization

```go
client := privateendpointconnectionresources.NewPrivateEndpointConnectionResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionResourcesClient.PrivateEndpointConnectionDelete`

```go
ctx := context.TODO()
id := privateendpointconnectionresources.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "privateEndpointConnectionName")

if err := client.PrivateEndpointConnectionDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionResourcesClient.PrivateEndpointConnectionGet`

```go
ctx := context.TODO()
id := privateendpointconnectionresources.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "privateEndpointConnectionName")

read, err := client.PrivateEndpointConnectionGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionResourcesClient.PrivateEndpointConnectionPut`

```go
ctx := context.TODO()
id := privateendpointconnectionresources.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "privateEndpointConnectionName")

payload := privateendpointconnectionresources.PrivateEndpointConnectionResource{
	// ...
}


if err := client.PrivateEndpointConnectionPutThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
