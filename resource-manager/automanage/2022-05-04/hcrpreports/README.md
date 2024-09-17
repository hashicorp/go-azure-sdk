
## `github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/hcrpreports` Documentation

The `hcrpreports` SDK allows for interaction with Azure Resource Manager `automanage` (API Version `2022-05-04`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/hcrpreports"
```


### Client Initialization

```go
client := hcrpreports.NewHCRPReportsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HCRPReportsClient.Get`

```go
ctx := context.TODO()
id := hcrpreports.NewConfigurationProfileAssignmentReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineValue", "configurationProfileAssignmentValue", "reportValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HCRPReportsClient.ListByConfigurationProfileAssignments`

```go
ctx := context.TODO()
id := hcrpreports.NewProviders2ConfigurationProfileAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineValue", "configurationProfileAssignmentValue")

read, err := client.ListByConfigurationProfileAssignments(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
