
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstancedecisioninstancestagedecision` Documentation

The `userpendingaccessreviewinstancedecisioninstancestagedecision` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstancedecisioninstancestagedecision"
```


### Client Initialization

```go
client := userpendingaccessreviewinstancedecisioninstancestagedecision.NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstancestagedecision.NewUserPendingAccessReviewInstanceDecisionInstanceStageID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue")

payload := userpendingaccessreviewinstancedecisioninstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstancestagedecision.NewUserPendingAccessReviewInstanceDecisionInstanceStageID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue")

payload := userpendingaccessreviewinstancedecisioninstancestagedecision.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionRecordAllDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient.DeleteUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstancestagedecision.NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemId1Value")

read, err := client.DeleteUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient.GetUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstancestagedecision.NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemId1Value")

read, err := client.GetUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient.GetUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionCount`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstancestagedecision.NewUserPendingAccessReviewInstanceDecisionInstanceStageID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue")

read, err := client.GetUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient.ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisions`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstancestagedecision.NewUserPendingAccessReviewInstanceDecisionInstanceStageID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue")

// alternatively `client.ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisions(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient.UpdateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecisioninstancestagedecision.NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemId1Value")

payload := userpendingaccessreviewinstancedecisioninstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdateUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
