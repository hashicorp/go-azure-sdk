
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-04-01/routinginfo` Documentation

The `routinginfo` SDK allows for interaction with the Azure Resource Manager Service `mobilenetwork` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-04-01/routinginfo"
```


### Client Initialization

```go
client := routinginfo.NewRoutingInfoClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RoutingInfoClient.Get`

```go
ctx := context.TODO()
id := routinginfo.NewPacketCoreControlPlaneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoutingInfoClient.List`

```go
ctx := context.TODO()
id := routinginfo.NewPacketCoreControlPlaneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
