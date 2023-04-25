
## `github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2020-06-25/guestconfigurationassignmenthcrpreports` Documentation

The `guestconfigurationassignmenthcrpreports` SDK allows for interaction with the Azure Resource Manager Service `guestconfiguration` (API Version `2020-06-25`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2020-06-25/guestconfigurationassignmenthcrpreports"
```


### Client Initialization

```go
client := guestconfigurationassignmenthcrpreports.NewGuestConfigurationAssignmentHCRPReportsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GuestConfigurationAssignmentHCRPReportsClient.GuestConfigurationHCRPAssignmentReportsGet`

```go
ctx := context.TODO()
id := guestconfigurationassignmenthcrpreports.NewReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineValue", "guestConfigurationAssignmentValue", "reportIdValue")

read, err := client.GuestConfigurationHCRPAssignmentReportsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestConfigurationAssignmentHCRPReportsClient.GuestConfigurationHCRPAssignmentReportsList`

```go
ctx := context.TODO()
id := guestconfigurationassignmenthcrpreports.NewGuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineValue", "guestConfigurationAssignmentValue")

read, err := client.GuestConfigurationHCRPAssignmentReportsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
