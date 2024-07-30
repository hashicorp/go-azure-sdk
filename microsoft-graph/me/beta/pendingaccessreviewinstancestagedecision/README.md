
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstancestagedecision` Documentation

The `pendingaccessreviewinstancestagedecision` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstancestagedecision"
```


### Client Initialization

```go
client := pendingaccessreviewinstancestagedecision.NewPendingAccessReviewInstanceStageDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionClient.CreatePendingAccessReviewInstanceStageDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageID("accessReviewInstanceIdValue", "accessReviewStageIdValue")

payload := pendingaccessreviewinstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceStageDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionClient.CreatePendingAccessReviewInstanceStageDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageID("accessReviewInstanceIdValue", "accessReviewStageIdValue")

payload := pendingaccessreviewinstancestagedecision.CreatePendingAccessReviewInstanceStageDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceStageDecisionRecordAllDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionClient.DeletePendingAccessReviewInstanceStageDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageIdDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeletePendingAccessReviewInstanceStageDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionClient.GetPendingAccessReviewInstanceStageDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageIdDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetPendingAccessReviewInstanceStageDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionClient.GetPendingAccessReviewInstanceStageDecisionCount`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageID("accessReviewInstanceIdValue", "accessReviewStageIdValue")

read, err := client.GetPendingAccessReviewInstanceStageDecisionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionClient.ListPendingAccessReviewInstanceStageDecisions`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageID("accessReviewInstanceIdValue", "accessReviewStageIdValue")

// alternatively `client.ListPendingAccessReviewInstanceStageDecisions(ctx, id)` can be used to do batched pagination
items, err := client.ListPendingAccessReviewInstanceStageDecisionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionClient.UpdatePendingAccessReviewInstanceStageDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageIdDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := pendingaccessreviewinstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdatePendingAccessReviewInstanceStageDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
