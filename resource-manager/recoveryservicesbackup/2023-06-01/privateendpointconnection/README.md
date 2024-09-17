
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-06-01/privateendpointconnection` Documentation

The `privateendpointconnection` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-06-01/privateendpointconnection"
```


### Client Initialization

```go
client := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionClient.Delete`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "privateEndpointConnectionValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionClient.Get`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "privateEndpointConnectionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionClient.Put`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "privateEndpointConnectionValue")

payload := privateendpointconnection.PrivateEndpointConnectionResource{
	// ...
}


if err := client.PutThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
