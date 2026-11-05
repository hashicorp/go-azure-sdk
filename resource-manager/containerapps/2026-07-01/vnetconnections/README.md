
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/vnetconnections` Documentation

The `vnetconnections` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2026-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/vnetconnections"
```


### Client Initialization

```go
client := vnetconnections.NewVnetConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VnetConnectionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := vnetconnections.NewVnetConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sandboxGroupName", "vnetConnectionName")

payload := vnetconnections.VnetConnection{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetConnectionsClient.Delete`

```go
ctx := context.TODO()
id := vnetconnections.NewVnetConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sandboxGroupName", "vnetConnectionName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VnetConnectionsClient.Get`

```go
ctx := context.TODO()
id := vnetconnections.NewVnetConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sandboxGroupName", "vnetConnectionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetConnectionsClient.ListBySandboxGroup`

```go
ctx := context.TODO()
id := vnetconnections.NewSandboxGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sandboxGroupName")

// alternatively `client.ListBySandboxGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListBySandboxGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
