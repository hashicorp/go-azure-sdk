
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstancestagedecisioninstancedecision` Documentation

The `userpendingaccessreviewinstancestagedecisioninstancedecision` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstancestagedecisioninstancedecision"
```


### Client Initialization

```go
client := userpendingaccessreviewinstancestagedecisioninstancedecision.NewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstancedecision.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := userpendingaccessreviewinstancestagedecisioninstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstancedecision.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := userpendingaccessreviewinstancestagedecisioninstancedecision.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionRecordAllDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient.DeleteUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstancedecision.NewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewInstanceDecisionItemId1Value")

read, err := client.DeleteUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient.GetUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstancedecision.NewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewInstanceDecisionItemId1Value")

read, err := client.GetUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient.GetUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionCount`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstancedecision.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient.ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisions`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstancedecision.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

// alternatively `client.ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisions(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient.UpdateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecisioninstancedecision.NewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewInstanceDecisionItemId1Value")

payload := userpendingaccessreviewinstancestagedecisioninstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
