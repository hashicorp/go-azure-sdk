
## `github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/privateendpointconnections` Documentation

The `privateendpointconnections` SDK allows for interaction with Azure Resource Manager `hardwaresecuritymodules` (API Version `2025-03-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/privateendpointconnections"
```


### Client Initialization

```go
client := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionsClient.CloudHsmClusterPrivateEndpointConnectionsCreate`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudHsmClusterName", "privateEndpointConnectionName")

payload := privateendpointconnections.PrivateEndpointConnection{
	// ...
}


read, err := client.CloudHsmClusterPrivateEndpointConnectionsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionsClient.CloudHsmClusterPrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudHsmClusterName", "privateEndpointConnectionName")

if err := client.CloudHsmClusterPrivateEndpointConnectionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.CloudHsmClusterPrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudHsmClusterName", "privateEndpointConnectionName")

read, err := client.CloudHsmClusterPrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ListByCloudHsmCluster`

```go
ctx := context.TODO()
id := privateendpointconnections.NewCloudHsmClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudHsmClusterName")

// alternatively `client.ListByCloudHsmCluster(ctx, id)` can be used to do batched pagination
items, err := client.ListByCloudHsmClusterComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
