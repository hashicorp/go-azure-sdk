
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstancestagedecisioninstance` Documentation

The `userpendingaccessreviewinstancestagedecisioninstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstancestagedecisioninstance"
```


### Client Initialization

```go
client := userpendingaccessreviewinstancestagedecisioninstance.NewUserPendingAccessReviewInstanceStageDecisionInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceAcceptRecommendation`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstance.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceAcceptRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceApplyDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstance.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceBatchRecordDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstance.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := userpendingaccessreviewinstancestagedecisioninstance.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceBatchRecordDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceResetDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstance.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceResetDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceSendReminder`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstance.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceSendReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceClient.DeleteUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstance.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeleteUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceClient.GetUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstance.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceClient.StopUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstance.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceClient.StopUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceApplyDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstance.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceClient.UpdateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstance.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := userpendingaccessreviewinstancestagedecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
