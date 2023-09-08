
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstancestagedecisioninstancedecision` Documentation

The `mependingaccessreviewinstancestagedecisioninstancedecision` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstancestagedecisioninstancedecision"
```


### Client Initialization

```go
client := mependingaccessreviewinstancestagedecisioninstancedecision.NewMePendingAccessReviewInstanceStageDecisionInstanceDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstancedecision.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := mependingaccessreviewinstancestagedecisioninstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstancedecision.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := mependingaccessreviewinstancestagedecisioninstancedecision.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionRecordAllDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient.DeleteMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstancedecision.NewMePendingAccessReviewInstanceStageDecisionInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewInstanceDecisionItemId1Value")

read, err := client.DeleteMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient.GetMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstancedecision.NewMePendingAccessReviewInstanceStageDecisionInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewInstanceDecisionItemId1Value")

read, err := client.GetMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient.GetMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionCount`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstancedecision.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient.ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisions`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstancedecision.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

// alternatively `client.ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisions(ctx, id)` can be used to do batched pagination
items, err := client.ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient.UpdateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstancedecision.NewMePendingAccessReviewInstanceStageDecisionInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewInstanceDecisionItemId1Value")

payload := mependingaccessreviewinstancestagedecisioninstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
