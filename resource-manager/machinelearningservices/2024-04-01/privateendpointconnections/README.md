
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/privateendpointconnections` Documentation

The `privateendpointconnections` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/privateendpointconnections"
```


### Client Initialization

```go
client := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionsClient.List`

```go
ctx := context.TODO()
id := privateendpointconnections.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
