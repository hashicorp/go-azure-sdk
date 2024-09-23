
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstance` Documentation

The `pendingaccessreviewinstancestagedecisioninstance` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstance"
```


### Client Initialization

```go
client := pendingaccessreviewinstancestagedecisioninstance.NewPendingAccessReviewInstanceStageDecisionInstanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendations`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendations(ctx, id, pendingaccessreviewinstancestagedecisioninstance.DefaultAcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendationsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.CreatePendingAccessReviewInstanceStageDecisionInstanceApplyDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.CreatePendingAccessReviewInstanceStageDecisionInstanceApplyDecision(ctx, id, pendingaccessreviewinstancestagedecisioninstance.DefaultCreatePendingAccessReviewInstanceStageDecisionInstanceApplyDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

payload := pendingaccessreviewinstancestagedecisioninstance.CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecision(ctx, id, payload, pendingaccessreviewinstancestagedecisioninstance.DefaultCreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.DeletePendingAccessReviewInstanceStageDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.DeletePendingAccessReviewInstanceStageDecisionInstance(ctx, id, pendingaccessreviewinstancestagedecisioninstance.DefaultDeletePendingAccessReviewInstanceStageDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.GetPendingAccessReviewInstanceStageDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.GetPendingAccessReviewInstanceStageDecisionInstance(ctx, id, pendingaccessreviewinstancestagedecisioninstance.DefaultGetPendingAccessReviewInstanceStageDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.ResetPendingAccessReviewInstanceStageDecisionInstanceDecisions`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.ResetPendingAccessReviewInstanceStageDecisionInstanceDecisions(ctx, id, pendingaccessreviewinstancestagedecisioninstance.DefaultResetPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.SendPendingAccessReviewInstanceStageDecisionInstanceReminder`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.SendPendingAccessReviewInstanceStageDecisionInstanceReminder(ctx, id, pendingaccessreviewinstancestagedecisioninstance.DefaultSendPendingAccessReviewInstanceStageDecisionInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.StopPendingAccessReviewInstanceStageDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.StopPendingAccessReviewInstanceStageDecisionInstance(ctx, id, pendingaccessreviewinstancestagedecisioninstance.DefaultStopPendingAccessReviewInstanceStageDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.StopPendingAccessReviewInstanceStageDecisionInstanceApplyDecisions`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.StopPendingAccessReviewInstanceStageDecisionInstanceApplyDecisions(ctx, id, pendingaccessreviewinstancestagedecisioninstance.DefaultStopPendingAccessReviewInstanceStageDecisionInstanceApplyDecisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.UpdatePendingAccessReviewInstanceStageDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

payload := pendingaccessreviewinstancestagedecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdatePendingAccessReviewInstanceStageDecisionInstance(ctx, id, payload, pendingaccessreviewinstancestagedecisioninstance.DefaultUpdatePendingAccessReviewInstanceStageDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
