
## `github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2022-01-25/guestconfigurationconnectedvmwarevsphereassignments` Documentation

The `guestconfigurationconnectedvmwarevsphereassignments` SDK allows for interaction with the Azure Resource Manager Service `guestconfiguration` (API Version `2022-01-25`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2022-01-25/guestconfigurationconnectedvmwarevsphereassignments"
```


### Client Initialization

```go
client := guestconfigurationconnectedvmwarevsphereassignments.NewGuestConfigurationConnectedVMwarevSphereAssignmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GuestConfigurationConnectedVMwarevSphereAssignmentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := guestconfigurationconnectedvmwarevsphereassignments.NewProviderVirtualMachineProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue")

payload := guestconfigurationconnectedvmwarevsphereassignments.GuestConfigurationAssignment{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationConnectedVMwarevSphereAssignmentsClient.Delete`

```go
ctx := context.TODO()
id := guestconfigurationconnectedvmwarevsphereassignments.NewProviderVirtualMachineProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationConnectedVMwarevSphereAssignmentsClient.Get`

```go
ctx := context.TODO()
id := guestconfigurationconnectedvmwarevsphereassignments.NewProviderVirtualMachineProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationConnectedVMwarevSphereAssignmentsClient.List`

```go
ctx := context.TODO()
id := guestconfigurationconnectedvmwarevsphereassignments.NewProviderVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
