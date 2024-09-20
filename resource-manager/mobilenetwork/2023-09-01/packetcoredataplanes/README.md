
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcoredataplanes` Documentation

The `packetcoredataplanes` SDK allows for interaction with Azure Resource Manager `mobilenetwork` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcoredataplanes"
```


### Client Initialization

```go
client := packetcoredataplanes.NewPacketCoreDataPlanesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PacketCoreDataPlanesClient.ListByPacketCoreControlPlane`

```go
ctx := context.TODO()
id := packetcoredataplanes.NewPacketCoreControlPlaneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneName")

// alternatively `client.ListByPacketCoreControlPlane(ctx, id)` can be used to do batched pagination
items, err := client.ListByPacketCoreControlPlaneComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
