
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstancedecision` Documentation

The `pendingaccessreviewinstancestagedecisioninstancedecision` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstancedecision"
```


### Client Initialization

```go
client := pendingaccessreviewinstancestagedecisioninstancedecision.NewPendingAccessReviewInstanceStageDecisionInstanceDecisionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceDecisionClient.CreatePendingAccessReviewInstanceStageDecisionInstanceDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstancedecision.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

payload := pendingaccessreviewinstancestagedecisioninstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceStageDecisionInstanceDecision(ctx, id, payload, pendingaccessreviewinstancestagedecisioninstancedecision.DefaultCreatePendingAccessReviewInstanceStageDecisionInstanceDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceDecisionClient.CreatePendingAccessReviewInstanceStageDecisionInstanceDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstancedecision.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

payload := pendingaccessreviewinstancestagedecisioninstancedecision.CreatePendingAccessReviewInstanceStageDecisionInstanceDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceStageDecisionInstanceDecisionRecordAllDecision(ctx, id, payload, pendingaccessreviewinstancestagedecisioninstancedecision.DefaultCreatePendingAccessReviewInstanceStageDecisionInstanceDecisionRecordAllDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceDecisionClient.GetPendingAccessReviewInstanceStageDecisionInstanceDecisionsCount`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstancedecision.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.GetPendingAccessReviewInstanceStageDecisionInstanceDecisionsCount(ctx, id, pendingaccessreviewinstancestagedecisioninstancedecision.DefaultGetPendingAccessReviewInstanceStageDecisionInstanceDecisionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceDecisionClient.ListPendingAccessReviewInstanceStageDecisionInstanceDecisions`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstancedecision.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

// alternatively `client.ListPendingAccessReviewInstanceStageDecisionInstanceDecisions(ctx, id, pendingaccessreviewinstancestagedecisioninstancedecision.DefaultListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions())` can be used to do batched pagination
items, err := client.ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsComplete(ctx, id, pendingaccessreviewinstancestagedecisioninstancedecision.DefaultListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
