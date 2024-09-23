
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedecisioninstancestagedecision` Documentation

The `pendingaccessreviewinstancedecisioninstancestagedecision` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedecisioninstancestagedecision"
```


### Client Initialization

```go
client := pendingaccessreviewinstancedecisioninstancestagedecision.NewPendingAccessReviewInstanceDecisionInstanceStageDecisionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceStageDecisionClient.CreatePendingAccessReviewInstanceDecisionInstanceStageDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstancestagedecision.NewUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageID("userId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId", "accessReviewStageId")

payload := pendingaccessreviewinstancedecisioninstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceDecisionInstanceStageDecision(ctx, id, payload, pendingaccessreviewinstancedecisioninstancestagedecision.DefaultCreatePendingAccessReviewInstanceDecisionInstanceStageDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceStageDecisionClient.CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstancestagedecision.NewUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageID("userId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId", "accessReviewStageId")

payload := pendingaccessreviewinstancedecisioninstancestagedecision.CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecision(ctx, id, payload, pendingaccessreviewinstancedecisioninstancestagedecision.DefaultCreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceStageDecisionClient.GetPendingAccessReviewInstanceDecisionInstanceStageDecisionsCount`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstancestagedecision.NewUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageID("userId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId", "accessReviewStageId")

read, err := client.GetPendingAccessReviewInstanceDecisionInstanceStageDecisionsCount(ctx, id, pendingaccessreviewinstancedecisioninstancestagedecision.DefaultGetPendingAccessReviewInstanceDecisionInstanceStageDecisionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceStageDecisionClient.ListPendingAccessReviewInstanceDecisionInstanceStageDecisions`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstancestagedecision.NewUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageID("userId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId", "accessReviewStageId")

// alternatively `client.ListPendingAccessReviewInstanceDecisionInstanceStageDecisions(ctx, id, pendingaccessreviewinstancedecisioninstancestagedecision.DefaultListPendingAccessReviewInstanceDecisionInstanceStageDecisionsOperationOptions())` can be used to do batched pagination
items, err := client.ListPendingAccessReviewInstanceDecisionInstanceStageDecisionsComplete(ctx, id, pendingaccessreviewinstancedecisioninstancestagedecision.DefaultListPendingAccessReviewInstanceDecisionInstanceStageDecisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
