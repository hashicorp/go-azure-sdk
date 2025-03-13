
## `github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2024-10-01/privateendpointconnection` Documentation

The `privateendpointconnection` SDK allows for interaction with Azure Resource Manager `dashboard` (API Version `2024-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2024-10-01/privateendpointconnection"
```


### Client Initialization

```go
client := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionClient.Approve`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName", "privateEndpointConnectionName")

payload := privateendpointconnection.PrivateEndpointConnection{
	// ...
}


if err := client.ApproveThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionClient.Delete`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName", "privateEndpointConnectionName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionClient.Get`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName", "privateEndpointConnectionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionClient.List`

```go
ctx := context.TODO()
id := privateendpointconnection.NewGrafanaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
