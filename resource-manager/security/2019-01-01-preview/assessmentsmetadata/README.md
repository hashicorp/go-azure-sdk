
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/assessmentsmetadata` Documentation

The `assessmentsmetadata` SDK allows for interaction with Azure Resource Manager `security` (API Version `2019-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/assessmentsmetadata"
```


### Client Initialization

```go
client := assessmentsmetadata.NewAssessmentsMetadataClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AssessmentsMetadataClient.Get`

```go
ctx := context.TODO()
id := assessmentsmetadata.NewAssessmentMetadataID("assessmentMetadataValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssessmentsMetadataClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx)` can be used to do batched pagination
items, err := client.ListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AssessmentsMetadataClient.SubscriptionCreate`

```go
ctx := context.TODO()
id := assessmentsmetadata.NewProviderAssessmentMetadataID("12345678-1234-9876-4563-123456789012", "assessmentMetadataValue")

payload := assessmentsmetadata.SecurityAssessmentMetadata{
	// ...
}


read, err := client.SubscriptionCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssessmentsMetadataClient.SubscriptionDelete`

```go
ctx := context.TODO()
id := assessmentsmetadata.NewProviderAssessmentMetadataID("12345678-1234-9876-4563-123456789012", "assessmentMetadataValue")

read, err := client.SubscriptionDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssessmentsMetadataClient.SubscriptionGet`

```go
ctx := context.TODO()
id := assessmentsmetadata.NewProviderAssessmentMetadataID("12345678-1234-9876-4563-123456789012", "assessmentMetadataValue")

read, err := client.SubscriptionGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssessmentsMetadataClient.SubscriptionList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.SubscriptionList(ctx, id)` can be used to do batched pagination
items, err := client.SubscriptionListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
