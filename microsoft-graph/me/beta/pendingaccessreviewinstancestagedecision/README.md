
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstancestagedecision` Documentation

The `pendingaccessreviewinstancestagedecision` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstancestagedecision"
```


### Client Initialization

```go
client := pendingaccessreviewinstancestagedecision.NewPendingAccessReviewInstanceStageDecisionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionClient.CreatePendingAccessReviewInstanceStageDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageID("accessReviewInstanceId", "accessReviewStageId")

payload := pendingaccessreviewinstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceStageDecision(ctx, id, payload, pendingaccessreviewinstancestagedecision.DefaultCreatePendingAccessReviewInstanceStageDecisionOperationOptions())
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
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageID("accessReviewInstanceId", "accessReviewStageId")

payload := pendingaccessreviewinstancestagedecision.CreatePendingAccessReviewInstanceStageDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceStageDecisionRecordAllDecision(ctx, id, payload, pendingaccessreviewinstancestagedecision.DefaultCreatePendingAccessReviewInstanceStageDecisionRecordAllDecisionOperationOptions())
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
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageIdDecisionID("accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.DeletePendingAccessReviewInstanceStageDecision(ctx, id, pendingaccessreviewinstancestagedecision.DefaultDeletePendingAccessReviewInstanceStageDecisionOperationOptions())
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
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageIdDecisionID("accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.GetPendingAccessReviewInstanceStageDecision(ctx, id, pendingaccessreviewinstancestagedecision.DefaultGetPendingAccessReviewInstanceStageDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionClient.GetPendingAccessReviewInstanceStageDecisionsCount`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageID("accessReviewInstanceId", "accessReviewStageId")

read, err := client.GetPendingAccessReviewInstanceStageDecisionsCount(ctx, id, pendingaccessreviewinstancestagedecision.DefaultGetPendingAccessReviewInstanceStageDecisionsCountOperationOptions())
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
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageID("accessReviewInstanceId", "accessReviewStageId")

// alternatively `client.ListPendingAccessReviewInstanceStageDecisions(ctx, id, pendingaccessreviewinstancestagedecision.DefaultListPendingAccessReviewInstanceStageDecisionsOperationOptions())` can be used to do batched pagination
items, err := client.ListPendingAccessReviewInstanceStageDecisionsComplete(ctx, id, pendingaccessreviewinstancestagedecision.DefaultListPendingAccessReviewInstanceStageDecisionsOperationOptions())
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
id := pendingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceIdStageIdDecisionID("accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

payload := pendingaccessreviewinstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdatePendingAccessReviewInstanceStageDecision(ctx, id, payload, pendingaccessreviewinstancestagedecision.DefaultUpdatePendingAccessReviewInstanceStageDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
