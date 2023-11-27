
## `github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2023-09-06/workspaceprivateendpointconnections` Documentation

The `workspaceprivateendpointconnections` SDK allows for interaction with the Azure Resource Manager Service `healthcareapis` (API Version `2023-09-06`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2023-09-06/workspaceprivateendpointconnections"
```


### Client Initialization

```go
client := workspaceprivateendpointconnections.NewWorkspacePrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspacePrivateEndpointConnectionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := workspaceprivateendpointconnections.NewWorkspacePrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "privateEndpointConnectionValue")

payload := workspaceprivateendpointconnections.PrivateEndpointConnectionDescription{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkspacePrivateEndpointConnectionsClient.Delete`

```go
ctx := context.TODO()
id := workspaceprivateendpointconnections.NewWorkspacePrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "privateEndpointConnectionValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkspacePrivateEndpointConnectionsClient.Get`

```go
ctx := context.TODO()
id := workspaceprivateendpointconnections.NewWorkspacePrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "privateEndpointConnectionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspacePrivateEndpointConnectionsClient.ListByWorkspace`

```go
ctx := context.TODO()
id := workspaceprivateendpointconnections.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

read, err := client.ListByWorkspace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
