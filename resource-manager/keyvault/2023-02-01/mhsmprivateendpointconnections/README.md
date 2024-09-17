
## `github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-02-01/mhsmprivateendpointconnections` Documentation

The `mhsmprivateendpointconnections` SDK allows for interaction with Azure Resource Manager `keyvault` (API Version `2023-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-02-01/mhsmprivateendpointconnections"
```


### Client Initialization

```go
client := mhsmprivateendpointconnections.NewMHSMPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MHSMPrivateEndpointConnectionsClient.Delete`

```go
ctx := context.TODO()
id := mhsmprivateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMValue", "privateEndpointConnectionValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MHSMPrivateEndpointConnectionsClient.Get`

```go
ctx := context.TODO()
id := mhsmprivateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMValue", "privateEndpointConnectionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MHSMPrivateEndpointConnectionsClient.Put`

```go
ctx := context.TODO()
id := mhsmprivateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMValue", "privateEndpointConnectionValue")

payload := mhsmprivateendpointconnections.MHSMPrivateEndpointConnection{
	// ...
}


read, err := client.Put(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
