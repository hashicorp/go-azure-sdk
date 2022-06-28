
## `github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-07-01/privateendpointconnection` Documentation

The `privateendpointconnection` SDK allows for interaction with the Azure Resource Manager Service `purview` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-07-01/privateendpointconnection"
```


### Client Initialization

```go
client := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "privateEndpointConnectionValue")

payload := privateendpointconnection.PrivateEndpointConnection{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionClient.Delete`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "privateEndpointConnectionValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionClient.Get`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "privateEndpointConnectionValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionClient.ListByAccount`

```go
ctx := context.TODO()
id := privateendpointconnection.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")
// alternatively `client.ListByAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
