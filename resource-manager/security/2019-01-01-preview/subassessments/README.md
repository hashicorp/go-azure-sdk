
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/subassessments` Documentation

The `subassessments` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2019-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/subassessments"
```


### Client Initialization

```go
client := subassessments.NewSubAssessmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SubAssessmentsClient.Get`

```go
ctx := context.TODO()
id := subassessments.NewScopedSubAssessmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "assessmentValue", "subAssessmentValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubAssessmentsClient.List`

```go
ctx := context.TODO()
id := subassessments.NewScopedAssessmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "assessmentValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SubAssessmentsClient.ListAll`

```go
ctx := context.TODO()
id := subassessments.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ListAll(ctx, id)` can be used to do batched pagination
items, err := client.ListAllComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
