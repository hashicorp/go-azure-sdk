
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstancestagedecision` Documentation

The `mependingaccessreviewinstancestagedecision` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstancestagedecision"
```


### Client Initialization

```go
client := mependingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceStageDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionClient.CreateMePendingAccessReviewInstanceByIdStageByIdDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceStageID("accessReviewInstanceIdValue", "accessReviewStageIdValue")

payload := mependingaccessreviewinstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateMePendingAccessReviewInstanceByIdStageByIdDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionClient.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceStageID("accessReviewInstanceIdValue", "accessReviewStageIdValue")

payload := mependingaccessreviewinstancestagedecision.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionRecordAllDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionClient.DeleteMePendingAccessReviewInstanceByIdStageByIdDecisionById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeleteMePendingAccessReviewInstanceByIdStageByIdDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionClient.GetMePendingAccessReviewInstanceByIdStageByIdDecisionById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetMePendingAccessReviewInstanceByIdStageByIdDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionClient.GetMePendingAccessReviewInstanceByIdStageByIdDecisionCount`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceStageID("accessReviewInstanceIdValue", "accessReviewStageIdValue")

read, err := client.GetMePendingAccessReviewInstanceByIdStageByIdDecisionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionClient.ListMePendingAccessReviewInstanceByIdStageByIdDecisions`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceStageID("accessReviewInstanceIdValue", "accessReviewStageIdValue")

// alternatively `client.ListMePendingAccessReviewInstanceByIdStageByIdDecisions(ctx, id)` can be used to do batched pagination
items, err := client.ListMePendingAccessReviewInstanceByIdStageByIdDecisionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionClient.UpdateMePendingAccessReviewInstanceByIdStageByIdDecisionById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := mependingaccessreviewinstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdateMePendingAccessReviewInstanceByIdStageByIdDecisionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
