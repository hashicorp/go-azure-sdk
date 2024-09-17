
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-02-01/diagnosticspackages` Documentation

The `diagnosticspackages` SDK allows for interaction with Azure Resource Manager `mobilenetwork` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-02-01/diagnosticspackages"
```


### Client Initialization

```go
client := diagnosticspackages.NewDiagnosticsPackagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiagnosticsPackagesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := diagnosticspackages.NewDiagnosticsPackageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneValue", "diagnosticsPackageValue")

if err := client.CreateOrUpdateThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DiagnosticsPackagesClient.Delete`

```go
ctx := context.TODO()
id := diagnosticspackages.NewDiagnosticsPackageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneValue", "diagnosticsPackageValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DiagnosticsPackagesClient.Get`

```go
ctx := context.TODO()
id := diagnosticspackages.NewDiagnosticsPackageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneValue", "diagnosticsPackageValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsPackagesClient.ListByPacketCoreControlPlane`

```go
ctx := context.TODO()
id := diagnosticspackages.NewPacketCoreControlPlaneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneValue")

// alternatively `client.ListByPacketCoreControlPlane(ctx, id)` can be used to do batched pagination
items, err := client.ListByPacketCoreControlPlaneComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
