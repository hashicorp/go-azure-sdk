
## `github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2024-04-05/guestconfigurations` Documentation

The `guestconfigurations` SDK allows for interaction with Azure Resource Manager `guestconfiguration` (API Version `2024-04-05`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2024-04-05/guestconfigurations"
```


### Client Initialization

```go
client := guestconfigurations.NewGuestconfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GuestconfigurationsClient.GuestConfigurationAssignmentReportsVMSSGet`

```go
ctx := context.TODO()
id := guestconfigurations.NewReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "guestConfigurationAssignmentName", "reportName")

read, err := client.GuestConfigurationAssignmentReportsVMSSGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestconfigurationsClient.GuestConfigurationAssignmentReportsVMSSList`

```go
ctx := context.TODO()
id := guestconfigurations.NewGuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "guestConfigurationAssignmentName")

// alternatively `client.GuestConfigurationAssignmentReportsVMSSList(ctx, id)` can be used to do batched pagination
items, err := client.GuestConfigurationAssignmentReportsVMSSListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GuestconfigurationsClient.GuestConfigurationAssignmentsRGList`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.GuestConfigurationAssignmentsRGList(ctx, id)` can be used to do batched pagination
items, err := client.GuestConfigurationAssignmentsRGListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GuestconfigurationsClient.GuestConfigurationAssignmentsVMSSCreateOrUpdate`

```go
ctx := context.TODO()
id := guestconfigurations.NewGuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "guestConfigurationAssignmentName")

payload := guestconfigurations.GuestConfigurationAssignment{
	// ...
}


read, err := client.GuestConfigurationAssignmentsVMSSCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestconfigurationsClient.GuestConfigurationAssignmentsVMSSDelete`

```go
ctx := context.TODO()
id := guestconfigurations.NewGuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "guestConfigurationAssignmentName")

read, err := client.GuestConfigurationAssignmentsVMSSDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestconfigurationsClient.GuestConfigurationAssignmentsVMSSGet`

```go
ctx := context.TODO()
id := guestconfigurations.NewGuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "guestConfigurationAssignmentName")

read, err := client.GuestConfigurationAssignmentsVMSSGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestconfigurationsClient.GuestConfigurationAssignmentsVMSSList`

```go
ctx := context.TODO()
id := guestconfigurations.NewVirtualMachineScaleSetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName")

// alternatively `client.GuestConfigurationAssignmentsVMSSList(ctx, id)` can be used to do batched pagination
items, err := client.GuestConfigurationAssignmentsVMSSListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
