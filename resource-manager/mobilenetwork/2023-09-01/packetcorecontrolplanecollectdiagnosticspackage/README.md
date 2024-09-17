
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcorecontrolplanecollectdiagnosticspackage` Documentation

The `packetcorecontrolplanecollectdiagnosticspackage` SDK allows for interaction with Azure Resource Manager `mobilenetwork` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcorecontrolplanecollectdiagnosticspackage"
```


### Client Initialization

```go
client := packetcorecontrolplanecollectdiagnosticspackage.NewPacketCoreControlPlaneCollectDiagnosticsPackageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PacketCoreControlPlaneCollectDiagnosticsPackageClient.PacketCoreControlPlanesCollectDiagnosticsPackage`

```go
ctx := context.TODO()
id := packetcorecontrolplanecollectdiagnosticspackage.NewPacketCoreControlPlaneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneValue")

payload := packetcorecontrolplanecollectdiagnosticspackage.PacketCoreControlPlaneCollectDiagnosticsPackage{
	// ...
}


if err := client.PacketCoreControlPlanesCollectDiagnosticsPackageThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
