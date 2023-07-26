
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/diagnosticsettings` Documentation

The `diagnosticsettings` SDK allows for interaction with the Azure Resource Manager Service `databoxedge` (API Version `2023-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/diagnosticsettings"
```


### Client Initialization

```go
client := diagnosticsettings.NewDiagnosticSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiagnosticSettingsClient.GetDiagnosticProactiveLogCollectionSettings`

```go
ctx := context.TODO()
id := diagnosticsettings.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue")

read, err := client.GetDiagnosticProactiveLogCollectionSettings(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticSettingsClient.GetDiagnosticRemoteSupportSettings`

```go
ctx := context.TODO()
id := diagnosticsettings.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue")

read, err := client.GetDiagnosticRemoteSupportSettings(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticSettingsClient.UpdateDiagnosticProactiveLogCollectionSettings`

```go
ctx := context.TODO()
id := diagnosticsettings.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue")

payload := diagnosticsettings.DiagnosticProactiveLogCollectionSettings{
	// ...
}


if err := client.UpdateDiagnosticProactiveLogCollectionSettingsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DiagnosticSettingsClient.UpdateDiagnosticRemoteSupportSettings`

```go
ctx := context.TODO()
id := diagnosticsettings.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue")

payload := diagnosticsettings.DiagnosticRemoteSupportSettings{
	// ...
}


if err := client.UpdateDiagnosticRemoteSupportSettingsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
