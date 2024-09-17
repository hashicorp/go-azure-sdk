
## `github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2020-06-25/guestconfigurationassignmentreports` Documentation

The `guestconfigurationassignmentreports` SDK allows for interaction with Azure Resource Manager `guestconfiguration` (API Version `2020-06-25`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2020-06-25/guestconfigurationassignmentreports"
```


### Client Initialization

```go
client := guestconfigurationassignmentreports.NewGuestConfigurationAssignmentReportsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GuestConfigurationAssignmentReportsClient.Get`

```go
ctx := context.TODO()
id := guestconfigurationassignmentreports.NewGuestConfigurationAssignmentReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue", "reportIdValue")

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
id := guestconfigurationassignmentreports.NewProviders2GuestConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "guestConfigurationAssignmentValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
