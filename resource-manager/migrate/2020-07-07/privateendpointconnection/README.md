
## `github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/privateendpointconnection` Documentation

The `privateendpointconnection` SDK allows for interaction with the Azure Resource Manager Service `migrate` (API Version `2020-07-07`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/privateendpointconnection"
```


### Client Initialization

```go
client := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionClient.DeletePrivateEndpointConnection`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "masterSiteValue", "privateEndpointConnectionValue")

read, err := client.DeletePrivateEndpointConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionClient.GetPrivateEndpointConnection`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "masterSiteValue", "privateEndpointConnectionValue")

read, err := client.GetPrivateEndpointConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionClient.GetPrivateEndpointConnections`

```go
ctx := context.TODO()
id := privateendpointconnection.NewMasterSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "masterSiteValue")

// alternatively `client.GetPrivateEndpointConnections(ctx, id)` can be used to do batched pagination
items, err := client.GetPrivateEndpointConnectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateEndpointConnectionClient.PutPrivateEndpointConnection`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "masterSiteValue", "privateEndpointConnectionValue")

payload := privateendpointconnection.PrivateEndpointConnection{
	// ...
}


read, err := client.PutPrivateEndpointConnection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
