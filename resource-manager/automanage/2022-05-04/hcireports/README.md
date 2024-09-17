
## `github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/hcireports` Documentation

The `hcireports` SDK allows for interaction with Azure Resource Manager `automanage` (API Version `2022-05-04`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/hcireports"
```


### Client Initialization

```go
client := hcireports.NewHCIReportsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HCIReportsClient.Get`

```go
ctx := context.TODO()
id := hcireports.NewReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "configurationProfileAssignmentValue", "reportValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HCIReportsClient.ListByConfigurationProfileAssignments`

```go
ctx := context.TODO()
id := hcireports.NewConfigurationProfileAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "configurationProfileAssignmentValue")

read, err := client.ListByConfigurationProfileAssignments(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
