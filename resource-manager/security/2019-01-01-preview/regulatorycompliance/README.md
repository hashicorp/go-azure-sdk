
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/regulatorycompliance` Documentation

The `regulatorycompliance` SDK allows for interaction with Azure Resource Manager `security` (API Version `2019-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/regulatorycompliance"
```


### Client Initialization

```go
client := regulatorycompliance.NewRegulatoryComplianceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RegulatoryComplianceClient.AssessmentsGet`

```go
ctx := context.TODO()
id := regulatorycompliance.NewRegulatoryComplianceAssessmentID("12345678-1234-9876-4563-123456789012", "regulatoryComplianceStandardName", "regulatoryComplianceControlName", "regulatoryComplianceAssessmentName")

read, err := client.AssessmentsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RegulatoryComplianceClient.AssessmentsList`

```go
ctx := context.TODO()
id := regulatorycompliance.NewRegulatoryComplianceControlID("12345678-1234-9876-4563-123456789012", "regulatoryComplianceStandardName", "regulatoryComplianceControlName")

// alternatively `client.AssessmentsList(ctx, id, regulatorycompliance.DefaultAssessmentsListOperationOptions())` can be used to do batched pagination
items, err := client.AssessmentsListComplete(ctx, id, regulatorycompliance.DefaultAssessmentsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RegulatoryComplianceClient.ControlsGet`

```go
ctx := context.TODO()
id := regulatorycompliance.NewRegulatoryComplianceControlID("12345678-1234-9876-4563-123456789012", "regulatoryComplianceStandardName", "regulatoryComplianceControlName")

read, err := client.ControlsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RegulatoryComplianceClient.ControlsList`

```go
ctx := context.TODO()
id := regulatorycompliance.NewRegulatoryComplianceStandardID("12345678-1234-9876-4563-123456789012", "regulatoryComplianceStandardName")

// alternatively `client.ControlsList(ctx, id, regulatorycompliance.DefaultControlsListOperationOptions())` can be used to do batched pagination
items, err := client.ControlsListComplete(ctx, id, regulatorycompliance.DefaultControlsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RegulatoryComplianceClient.StandardsGet`

```go
ctx := context.TODO()
id := regulatorycompliance.NewRegulatoryComplianceStandardID("12345678-1234-9876-4563-123456789012", "regulatoryComplianceStandardName")

read, err := client.StandardsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RegulatoryComplianceClient.StandardsList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.StandardsList(ctx, id, regulatorycompliance.DefaultStandardsListOperationOptions())` can be used to do batched pagination
items, err := client.StandardsListComplete(ctx, id, regulatorycompliance.DefaultStandardsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
