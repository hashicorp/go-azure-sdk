
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/packetcorecontrolplanereinstall` Documentation

The `packetcorecontrolplanereinstall` SDK allows for interaction with Azure Resource Manager `mobilenetwork` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/packetcorecontrolplanereinstall"
```


### Client Initialization

```go
client := packetcorecontrolplanereinstall.NewPacketCoreControlPlaneReinstallClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PacketCoreControlPlaneReinstallClient.PacketCoreControlPlanesReinstall`

```go
ctx := context.TODO()
id := packetcorecontrolplanereinstall.NewPacketCoreControlPlaneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneName")

if err := client.PacketCoreControlPlanesReinstallThenPoll(ctx, id); err != nil {
	// handle the error
}
```
