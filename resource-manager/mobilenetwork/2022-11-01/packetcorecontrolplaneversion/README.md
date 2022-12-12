
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcorecontrolplaneversion` Documentation

The `packetcorecontrolplaneversion` SDK allows for interaction with the Azure Resource Manager Service `mobilenetwork` (API Version `2022-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcorecontrolplaneversion"
```


### Client Initialization

```go
client := packetcorecontrolplaneversion.NewPacketCoreControlPlaneVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PacketCoreControlPlaneVersionClient.Get`

```go
ctx := context.TODO()
id := packetcorecontrolplaneversion.NewPacketCoreControlPlaneVersionID("versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PacketCoreControlPlaneVersionClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx)` can be used to do batched pagination
items, err := client.ListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
