
## `github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2024-04-05/guestconfigurationassignmenthcrpreports` Documentation

The `guestconfigurationassignmenthcrpreports` SDK allows for interaction with Azure Resource Manager `guestconfiguration` (API Version `2024-04-05`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2024-04-05/guestconfigurationassignmenthcrpreports"
```


### Client Initialization

```go
client := guestconfigurationassignmenthcrpreports.NewGuestConfigurationAssignmentHCRPReportsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GuestConfigurationAssignmentHCRPReportsClient.GuestConfigurationHCRPAssignmentReportsGet`

```go
ctx := context.TODO()
id := guestconfigurationassignmenthcrpreports.NewGuestConfigurationAssignmentReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineName", "guestConfigurationAssignmentName", "reportId")

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
id := guestconfigurationassignmenthcrpreports.NewProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineName", "guestConfigurationAssignmentName")

read, err := client.GuestConfigurationHCRPAssignmentReportsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
