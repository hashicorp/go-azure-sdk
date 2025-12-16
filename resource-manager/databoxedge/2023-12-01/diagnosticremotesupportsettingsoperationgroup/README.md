
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/diagnosticremotesupportsettingsoperationgroup` Documentation

The `diagnosticremotesupportsettingsoperationgroup` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/diagnosticremotesupportsettingsoperationgroup"
```


### Client Initialization

```go
client := diagnosticremotesupportsettingsoperationgroup.NewDiagnosticRemoteSupportSettingsOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiagnosticRemoteSupportSettingsOperationGroupClient.DiagnosticSettingsGetDiagnosticRemoteSupportSettings`

```go
ctx := context.TODO()
id := diagnosticremotesupportsettingsoperationgroup.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

read, err := client.DiagnosticSettingsGetDiagnosticRemoteSupportSettings(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticRemoteSupportSettingsOperationGroupClient.DiagnosticSettingsUpdateDiagnosticRemoteSupportSettings`

```go
ctx := context.TODO()
id := diagnosticremotesupportsettingsoperationgroup.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

payload := diagnosticremotesupportsettingsoperationgroup.DiagnosticRemoteSupportSettings{
	// ...
}


if err := client.DiagnosticSettingsUpdateDiagnosticRemoteSupportSettingsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
