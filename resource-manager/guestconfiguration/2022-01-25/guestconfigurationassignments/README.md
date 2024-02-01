
## `github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2022-01-25/guestconfigurationassignments` Documentation

The `guestconfigurationassignments` SDK allows for interaction with the Azure Resource Manager Service `guestconfiguration` (API Version `2022-01-25`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2022-01-25/guestconfigurationassignments"
```


### Client Initialization

```go
client := guestconfigurationassignments.NewGuestConfigurationAssignmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GuestConfigurationAssignmentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := guestconfigurationassignments.NewVirtualMachineProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue")

payload := guestconfigurationassignments.GuestConfigurationAssignment{
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


### Example Usage: `GuestConfigurationAssignmentsClient.Delete`

```go
ctx := context.TODO()
id := guestconfigurationassignments.NewVirtualMachineProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationAssignmentsClient.Get`

```go
ctx := context.TODO()
id := guestconfigurationassignments.NewVirtualMachineProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationAssignmentsClient.List`

```go
ctx := context.TODO()
id := guestconfigurationassignments.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationAssignmentsClient.RGList`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.RGList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationAssignmentsClient.SubscriptionList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.SubscriptionList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationAssignmentsClient.VMSSCreateOrUpdate`

```go
ctx := context.TODO()
id := guestconfigurationassignments.NewGuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetValue", "guestConfigurationAssignmentValue")

payload := guestconfigurationassignments.GuestConfigurationAssignment{
	// ...
}


read, err := client.VMSSCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationAssignmentsClient.VMSSDelete`

```go
ctx := context.TODO()
id := guestconfigurationassignments.NewGuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetValue", "guestConfigurationAssignmentValue")

read, err := client.VMSSDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationAssignmentsClient.VMSSGet`

```go
ctx := context.TODO()
id := guestconfigurationassignments.NewGuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetValue", "guestConfigurationAssignmentValue")

read, err := client.VMSSGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationAssignmentsClient.VMSSList`

```go
ctx := context.TODO()
id := guestconfigurationassignments.NewVirtualMachineScaleSetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetValue")

read, err := client.VMSSList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
