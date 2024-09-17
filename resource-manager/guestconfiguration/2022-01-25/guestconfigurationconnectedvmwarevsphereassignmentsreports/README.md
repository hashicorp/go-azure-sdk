
## `github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2022-01-25/guestconfigurationconnectedvmwarevsphereassignmentsreports` Documentation

The `guestconfigurationconnectedvmwarevsphereassignmentsreports` SDK allows for interaction with Azure Resource Manager `guestconfiguration` (API Version `2022-01-25`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2022-01-25/guestconfigurationconnectedvmwarevsphereassignmentsreports"
```


### Client Initialization

```go
client := guestconfigurationconnectedvmwarevsphereassignmentsreports.NewGuestConfigurationConnectedVMwarevSphereAssignmentsReportsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GuestConfigurationConnectedVMwarevSphereAssignmentsReportsClient.Get`

```go
ctx := context.TODO()
id := guestconfigurationconnectedvmwarevsphereassignmentsreports.NewVirtualMachineProviders2GuestConfigurationAssignmentReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue", "reportIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationConnectedVMwarevSphereAssignmentsReportsClient.List`

```go
ctx := context.TODO()
id := guestconfigurationconnectedvmwarevsphereassignmentsreports.NewProviderVirtualMachineProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
