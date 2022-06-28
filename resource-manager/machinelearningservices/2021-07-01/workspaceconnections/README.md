
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/workspaceconnections` Documentation

The `workspaceconnections` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/workspaceconnections"
```


### Client Initialization

```go
client := workspaceconnections.NewWorkspaceConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceConnectionsClient.Create`

```go
ctx := context.TODO()
id := workspaceconnections.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "connectionValue")

payload := workspaceconnections.WorkspaceConnection{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceConnectionsClient.Delete`

```go
ctx := context.TODO()
id := workspaceconnections.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "connectionValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceConnectionsClient.Get`

```go
ctx := context.TODO()
id := workspaceconnections.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "connectionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceConnectionsClient.List`

```go
ctx := context.TODO()
id := workspaceconnections.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

read, err := client.List(ctx, id, workspaceconnections.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
