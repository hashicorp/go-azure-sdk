
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-10-01-preview/cognitiveservicescommitmentplans` Documentation

The `cognitiveservicescommitmentplans` SDK allows for interaction with Azure Resource Manager `cognitive` (API Version `2023-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-10-01-preview/cognitiveservicescommitmentplans"
```


### Client Initialization

```go
client := cognitiveservicescommitmentplans.NewCognitiveServicesCommitmentPlansClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CognitiveServicesCommitmentPlansClient.CommitmentPlansCreateOrUpdatePlan`

```go
ctx := context.TODO()
id := cognitiveservicescommitmentplans.NewCommitmentPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "commitmentPlanValue")

payload := cognitiveservicescommitmentplans.CommitmentPlan{
	// ...
}


if err := client.CommitmentPlansCreateOrUpdatePlanThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CognitiveServicesCommitmentPlansClient.CommitmentPlansDeletePlan`

```go
ctx := context.TODO()
id := cognitiveservicescommitmentplans.NewCommitmentPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "commitmentPlanValue")

if err := client.CommitmentPlansDeletePlanThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `CognitiveServicesCommitmentPlansClient.CommitmentPlansGetPlan`

```go
ctx := context.TODO()
id := cognitiveservicescommitmentplans.NewCommitmentPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "commitmentPlanValue")

read, err := client.CommitmentPlansGetPlan(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CognitiveServicesCommitmentPlansClient.CommitmentPlansListPlansByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.CommitmentPlansListPlansByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.CommitmentPlansListPlansByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CognitiveServicesCommitmentPlansClient.CommitmentPlansListPlansBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.CommitmentPlansListPlansBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.CommitmentPlansListPlansBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CognitiveServicesCommitmentPlansClient.CommitmentPlansUpdatePlan`

```go
ctx := context.TODO()
id := cognitiveservicescommitmentplans.NewCommitmentPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "commitmentPlanValue")

payload := cognitiveservicescommitmentplans.PatchResourceTagsAndSku{
	// ...
}


if err := client.CommitmentPlansUpdatePlanThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
