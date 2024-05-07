
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-04-01/attacheddatanetworks` Documentation

The `attacheddatanetworks` SDK allows for interaction with the Azure Resource Manager Service `mobilenetwork` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-04-01/attacheddatanetworks"
```


### Client Initialization

```go
client := attacheddatanetworks.NewAttachedDataNetworksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AttachedDataNetworksClient.ListByPacketCoreDataPlane`

```go
ctx := context.TODO()
id := attacheddatanetworks.NewPacketCoreDataPlaneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneValue", "packetCoreDataPlaneValue")

// alternatively `client.ListByPacketCoreDataPlane(ctx, id)` can be used to do batched pagination
items, err := client.ListByPacketCoreDataPlaneComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
