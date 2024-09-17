
## `github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2024-04-05/guestconfigurationassignmentreports` Documentation

The `guestconfigurationassignmentreports` SDK allows for interaction with Azure Resource Manager `guestconfiguration` (API Version `2024-04-05`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2024-04-05/guestconfigurationassignmentreports"
```


### Client Initialization

```go
client := guestconfigurationassignmentreports.NewGuestConfigurationAssignmentReportsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GuestConfigurationAssignmentReportsClient.Get`

```go
ctx := context.TODO()
id := guestconfigurationassignmentreports.NewProviders2GuestConfigurationAssignmentReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue", "reportIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationAssignmentReportsClient.List`

```go
ctx := context.TODO()
id := guestconfigurationassignmentreports.NewVirtualMachineProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationAssignmentReportsClient.VMSSGet`

```go
ctx := context.TODO()
id := guestconfigurationassignmentreports.NewReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetValue", "guestConfigurationAssignmentValue", "reportValue")

read, err := client.VMSSGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationAssignmentReportsClient.VMSSList`

```go
ctx := context.TODO()
id := guestconfigurationassignmentreports.NewGuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetValue", "guestConfigurationAssignmentValue")

read, err := client.VMSSList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
