
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstancedecisioninstance` Documentation

The `pendingaccessreviewinstancedecisioninstance` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstancedecisioninstance"
```


### Client Initialization

```go
client := pendingaccessreviewinstancedecisioninstance.NewPendingAccessReviewInstanceDecisionInstanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.AcceptPendingAccessReviewInstanceDecisionInstanceRecommendations`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.AcceptPendingAccessReviewInstanceDecisionInstanceRecommendations(ctx, id, pendingaccessreviewinstancedecisioninstance.DefaultAcceptPendingAccessReviewInstanceDecisionInstanceRecommendationsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.CreatePendingAccessReviewInstanceDecisionInstanceApplyDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.CreatePendingAccessReviewInstanceDecisionInstanceApplyDecision(ctx, id, pendingaccessreviewinstancedecisioninstance.DefaultCreatePendingAccessReviewInstanceDecisionInstanceApplyDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.CreatePendingAccessReviewInstanceDecisionInstanceBatchRecordDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

payload := pendingaccessreviewinstancedecisioninstance.CreatePendingAccessReviewInstanceDecisionInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceDecisionInstanceBatchRecordDecision(ctx, id, payload, pendingaccessreviewinstancedecisioninstance.DefaultCreatePendingAccessReviewInstanceDecisionInstanceBatchRecordDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.DeletePendingAccessReviewInstanceDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.DeletePendingAccessReviewInstanceDecisionInstance(ctx, id, pendingaccessreviewinstancedecisioninstance.DefaultDeletePendingAccessReviewInstanceDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.GetPendingAccessReviewInstanceDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.GetPendingAccessReviewInstanceDecisionInstance(ctx, id, pendingaccessreviewinstancedecisioninstance.DefaultGetPendingAccessReviewInstanceDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.ResetPendingAccessReviewInstanceDecisionInstanceDecisions`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.ResetPendingAccessReviewInstanceDecisionInstanceDecisions(ctx, id, pendingaccessreviewinstancedecisioninstance.DefaultResetPendingAccessReviewInstanceDecisionInstanceDecisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.SendPendingAccessReviewInstanceDecisionInstanceReminder`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.SendPendingAccessReviewInstanceDecisionInstanceReminder(ctx, id, pendingaccessreviewinstancedecisioninstance.DefaultSendPendingAccessReviewInstanceDecisionInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.StopPendingAccessReviewInstanceDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.StopPendingAccessReviewInstanceDecisionInstance(ctx, id, pendingaccessreviewinstancedecisioninstance.DefaultStopPendingAccessReviewInstanceDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.StopPendingAccessReviewInstanceDecisionInstanceApplyDecisions`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.StopPendingAccessReviewInstanceDecisionInstanceApplyDecisions(ctx, id, pendingaccessreviewinstancedecisioninstance.DefaultStopPendingAccessReviewInstanceDecisionInstanceApplyDecisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.UpdatePendingAccessReviewInstanceDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

payload := pendingaccessreviewinstancedecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdatePendingAccessReviewInstanceDecisionInstance(ctx, id, payload, pendingaccessreviewinstancedecisioninstance.DefaultUpdatePendingAccessReviewInstanceDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
