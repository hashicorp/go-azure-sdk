
## `github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2024-04-05/guestconfigurationhcrpassignments` Documentation

The `guestconfigurationhcrpassignments` SDK allows for interaction with Azure Resource Manager `guestconfiguration` (API Version `2024-04-05`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2024-04-05/guestconfigurationhcrpassignments"
```


### Client Initialization

```go
client := guestconfigurationhcrpassignments.NewGuestConfigurationHCRPAssignmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GuestConfigurationHCRPAssignmentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := guestconfigurationhcrpassignments.NewProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineValue", "guestConfigurationAssignmentValue")

payload := guestconfigurationhcrpassignments.GuestConfigurationAssignment{
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


### Example Usage: `GuestConfigurationHCRPAssignmentsClient.Delete`

```go
ctx := context.TODO()
id := guestconfigurationhcrpassignments.NewProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineValue", "guestConfigurationAssignmentValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationHCRPAssignmentsClient.Get`

```go
ctx := context.TODO()
id := guestconfigurationhcrpassignments.NewProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineValue", "guestConfigurationAssignmentValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationHCRPAssignmentsClient.List`

```go
ctx := context.TODO()
id := guestconfigurationhcrpassignments.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
