
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstancestagedecision` Documentation

The `userpendingaccessreviewinstancestagedecision` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstancestagedecision"
```


### Client Initialization

```go
client := userpendingaccessreviewinstancestagedecision.NewUserPendingAccessReviewInstanceStageDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionClient.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecision.NewUserPendingAccessReviewInstanceStageID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue")

payload := userpendingaccessreviewinstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionClient.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecision.NewUserPendingAccessReviewInstanceStageID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue")

payload := userpendingaccessreviewinstancestagedecision.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionRecordAllDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionClient.DeleteUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecision.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeleteUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionClient.GetUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecision.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionClient.GetUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionCount`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecision.NewUserPendingAccessReviewInstanceStageID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue")

read, err := client.GetUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionClient.ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisions`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecision.NewUserPendingAccessReviewInstanceStageID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue")

// alternatively `client.ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisions(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserPendingAccessReviewInstanceStageDecisionClient.UpdateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancestagedecision.NewUserPendingAccessReviewInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := userpendingaccessreviewinstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdateUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
