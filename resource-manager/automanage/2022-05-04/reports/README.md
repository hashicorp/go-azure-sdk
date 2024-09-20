
## `github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/reports` Documentation

The `reports` SDK allows for interaction with Azure Resource Manager `automanage` (API Version `2022-05-04`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/reports"
```


### Client Initialization

```go
client := reports.NewReportsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReportsClient.Get`

```go
ctx := context.TODO()
id := reports.NewProviders2ConfigurationProfileAssignmentReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "configurationProfileAssignmentName", "reportName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportsClient.ListByConfigurationProfileAssignments`

```go
ctx := context.TODO()
id := reports.NewVirtualMachineProviders2ConfigurationProfileAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "configurationProfileAssignmentName")

read, err := client.ListByConfigurationProfileAssignments(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
