
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcorecontrolplanerollback` Documentation

The `packetcorecontrolplanerollback` SDK allows for interaction with Azure Resource Manager `mobilenetwork` (API Version `2022-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcorecontrolplanerollback"
```


### Client Initialization

```go
client := packetcorecontrolplanerollback.NewPacketCoreControlPlaneRollbackClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PacketCoreControlPlaneRollbackClient.PacketCoreControlPlanesRollback`

```go
ctx := context.TODO()
id := packetcorecontrolplanerollback.NewPacketCoreControlPlaneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneValue")

if err := client.PacketCoreControlPlanesRollbackThenPoll(ctx, id); err != nil {
	// handle the error
}
```
