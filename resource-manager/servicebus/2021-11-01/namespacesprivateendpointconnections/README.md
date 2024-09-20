
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/namespacesprivateendpointconnections` Documentation

The `namespacesprivateendpointconnections` SDK allows for interaction with Azure Resource Manager `servicebus` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/namespacesprivateendpointconnections"
```


### Client Initialization

```go
client := namespacesprivateendpointconnections.NewNamespacesPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NamespacesPrivateEndpointConnectionsClient.PrivateEndpointConnectionsCreateOrUpdate`

```go
ctx := context.TODO()
id := namespacesprivateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "privateEndpointConnectionName")

payload := namespacesprivateendpointconnections.PrivateEndpointConnection{
	// ...
}


read, err := client.PrivateEndpointConnectionsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespacesPrivateEndpointConnectionsClient.PrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := namespacesprivateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "privateEndpointConnectionName")

if err := client.PrivateEndpointConnectionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NamespacesPrivateEndpointConnectionsClient.PrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := namespacesprivateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "privateEndpointConnectionName")

read, err := client.PrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespacesPrivateEndpointConnectionsClient.PrivateEndpointConnectionsList`

```go
ctx := context.TODO()
id := namespacesprivateendpointconnections.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

// alternatively `client.PrivateEndpointConnectionsList(ctx, id)` can be used to do batched pagination
items, err := client.PrivateEndpointConnectionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
