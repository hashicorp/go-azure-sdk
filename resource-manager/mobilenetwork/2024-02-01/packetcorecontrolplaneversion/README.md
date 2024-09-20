
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-02-01/packetcorecontrolplaneversion` Documentation

The `packetcorecontrolplaneversion` SDK allows for interaction with Azure Resource Manager `mobilenetwork` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-02-01/packetcorecontrolplaneversion"
```


### Client Initialization

```go
client := packetcorecontrolplaneversion.NewPacketCoreControlPlaneVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PacketCoreControlPlaneVersionClient.Get`

```go
ctx := context.TODO()
id := packetcorecontrolplaneversion.NewPacketCoreControlPlaneVersionID("versionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PacketCoreControlPlaneVersionClient.GetBySubscription`

```go
ctx := context.TODO()
id := packetcorecontrolplaneversion.NewProviderPacketCoreControlPlaneVersionID("12345678-1234-9876-4563-123456789012", "versionName")

read, err := client.GetBySubscription(ctx, id)
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


### Example Usage: `PacketCoreControlPlaneVersionClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
