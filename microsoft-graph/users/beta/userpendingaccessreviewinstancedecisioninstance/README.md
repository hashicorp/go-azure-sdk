
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstancedecisioninstance` Documentation

The `userpendingaccessreviewinstancedecisioninstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstancedecisioninstance"
```


### Client Initialization

```go
client := userpendingaccessreviewinstancedecisioninstance.NewUserPendingAccessReviewInstanceDecisionInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceAcceptRecommendation`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstance.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceAcceptRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceApplyDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstance.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceBatchRecordDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstance.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := userpendingaccessreviewinstancedecisioninstance.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceBatchRecordDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceResetDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstance.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceResetDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceSendReminder`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstance.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceSendReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceClient.DeleteUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstance.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeleteUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceClient.GetUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstance.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceClient.StopUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstance.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceClient.StopUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceApplyDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstance.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceClient.UpdateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstance.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := userpendingaccessreviewinstancedecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
