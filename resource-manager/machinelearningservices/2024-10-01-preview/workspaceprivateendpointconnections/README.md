
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/workspaceprivateendpointconnections` Documentation

The `workspaceprivateendpointconnections` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/workspaceprivateendpointconnections"
```


### Client Initialization

```go
client := workspaceprivateendpointconnections.NewWorkspacePrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspacePrivateEndpointConnectionsClient.PrivateEndpointConnectionsCreateOrUpdate`

```go
ctx := context.TODO()
id := workspaceprivateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "privateEndpointConnectionName")

payload := workspaceprivateendpointconnections.PrivateEndpointConnection{
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


### Example Usage: `WorkspacePrivateEndpointConnectionsClient.PrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := workspaceprivateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "privateEndpointConnectionName")

read, err := client.PrivateEndpointConnectionsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspacePrivateEndpointConnectionsClient.PrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := workspaceprivateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "privateEndpointConnectionName")

read, err := client.PrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspacePrivateEndpointConnectionsClient.PrivateEndpointConnectionsList`

```go
ctx := context.TODO()
id := workspaceprivateendpointconnections.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

read, err := client.PrivateEndpointConnectionsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
