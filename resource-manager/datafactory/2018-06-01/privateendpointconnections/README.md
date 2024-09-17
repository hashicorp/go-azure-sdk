
## `github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/privateendpointconnections` Documentation

The `privateendpointconnections` SDK allows for interaction with Azure Resource Manager `datafactory` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/privateendpointconnections"
```


### Client Initialization

```go
client := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionsClient.ListByFactory`

```go
ctx := context.TODO()
id := privateendpointconnections.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue")

// alternatively `client.ListByFactory(ctx, id)` can be used to do batched pagination
items, err := client.ListByFactoryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateEndpointConnectionsClient.PrivateEndpointConnectionCreateOrUpdate`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "privateEndpointConnectionValue")

payload := privateendpointconnections.PrivateLinkConnectionApprovalRequestResource{
	// ...
}


read, err := client.PrivateEndpointConnectionCreateOrUpdate(ctx, id, payload, privateendpointconnections.DefaultPrivateEndpointConnectionCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionsClient.PrivateEndpointConnectionDelete`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "privateEndpointConnectionValue")

read, err := client.PrivateEndpointConnectionDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionsClient.PrivateEndpointConnectionGet`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "privateEndpointConnectionValue")

read, err := client.PrivateEndpointConnectionGet(ctx, id, privateendpointconnections.DefaultPrivateEndpointConnectionGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
