
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstancedecisioninstancestagedecision` Documentation

The `mependingaccessreviewinstancedecisioninstancestagedecision` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstancedecisioninstancestagedecision"
```


### Client Initialization

```go
client := mependingaccessreviewinstancedecisioninstancestagedecision.NewMePendingAccessReviewInstanceDecisionInstanceStageDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceStageDecisionClient.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstancestagedecision.NewMePendingAccessReviewInstanceDecisionInstanceStageID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue")

payload := mependingaccessreviewinstancedecisioninstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceStageDecisionClient.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstancestagedecision.NewMePendingAccessReviewInstanceDecisionInstanceStageID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue")

payload := mependingaccessreviewinstancedecisioninstancestagedecision.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionRecordAllDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceStageDecisionClient.DeleteMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstancestagedecision.NewMePendingAccessReviewInstanceDecisionInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemId1Value")

read, err := client.DeleteMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceStageDecisionClient.GetMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstancestagedecision.NewMePendingAccessReviewInstanceDecisionInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemId1Value")

read, err := client.GetMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceStageDecisionClient.GetMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionCount`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstancestagedecision.NewMePendingAccessReviewInstanceDecisionInstanceStageID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue")

read, err := client.GetMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceStageDecisionClient.ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisions`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstancestagedecision.NewMePendingAccessReviewInstanceDecisionInstanceStageID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue")

// alternatively `client.ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisions(ctx, id)` can be used to do batched pagination
items, err := client.ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceStageDecisionClient.UpdateMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstancestagedecision.NewMePendingAccessReviewInstanceDecisionInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemId1Value")

payload := mependingaccessreviewinstancedecisioninstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdateMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
