
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstance` Documentation

The `pendingaccessreviewinstancestagedecisioninstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstance"
```


### Client Initialization

```go
client := pendingaccessreviewinstancestagedecisioninstance.NewPendingAccessReviewInstanceStageDecisionInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.CreatePendingAccessReviewInstanceStageDecisionInstanceAcceptRecommendation`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreatePendingAccessReviewInstanceStageDecisionInstanceAcceptRecommendation(ctx, id)
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
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreatePendingAccessReviewInstanceStageDecisionInstanceApplyDecision(ctx, id)
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
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := pendingaccessreviewinstancestagedecisioninstance.CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.CreatePendingAccessReviewInstanceStageDecisionInstanceResetDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreatePendingAccessReviewInstanceStageDecisionInstanceResetDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.CreatePendingAccessReviewInstanceStageDecisionInstanceSendReminder`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreatePendingAccessReviewInstanceStageDecisionInstanceSendReminder(ctx, id)
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
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeletePendingAccessReviewInstanceStageDecisionInstance(ctx, id)
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
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetPendingAccessReviewInstanceStageDecisionInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.StopUserPendingAccessReviewInstanceStageDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopUserPendingAccessReviewInstanceStageDecisionInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceStageDecisionInstanceClient.StopUserPendingAccessReviewInstanceStageDecisionInstanceApplyDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopUserPendingAccessReviewInstanceStageDecisionInstanceApplyDecision(ctx, id)
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
id := pendingaccessreviewinstancestagedecisioninstance.NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := pendingaccessreviewinstancestagedecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdatePendingAccessReviewInstanceStageDecisionInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
