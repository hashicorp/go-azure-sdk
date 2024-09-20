
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcaptures` Documentation

The `packetcaptures` SDK allows for interaction with Azure Resource Manager `mobilenetwork` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-09-01/packetcaptures"
```


### Client Initialization

```go
client := packetcaptures.NewPacketCapturesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PacketCapturesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := packetcaptures.NewPacketCaptureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneName", "packetCaptureName")

payload := packetcaptures.PacketCapture{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PacketCapturesClient.Delete`

```go
ctx := context.TODO()
id := packetcaptures.NewPacketCaptureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneName", "packetCaptureName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PacketCapturesClient.Get`

```go
ctx := context.TODO()
id := packetcaptures.NewPacketCaptureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneName", "packetCaptureName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PacketCapturesClient.ListByPacketCoreControlPlane`

```go
ctx := context.TODO()
id := packetcaptures.NewPacketCoreControlPlaneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneName")

// alternatively `client.ListByPacketCoreControlPlane(ctx, id)` can be used to do batched pagination
items, err := client.ListByPacketCoreControlPlaneComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PacketCapturesClient.Stop`

```go
ctx := context.TODO()
id := packetcaptures.NewPacketCaptureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneName", "packetCaptureName")

if err := client.StopThenPoll(ctx, id); err != nil {
	// handle the error
}
```
