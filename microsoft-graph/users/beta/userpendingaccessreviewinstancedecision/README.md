
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstancedecision` Documentation

The `userpendingaccessreviewinstancedecision` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstancedecision"
```


### Client Initialization

```go
client := userpendingaccessreviewinstancedecision.NewUserPendingAccessReviewInstanceDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionClient.CreateUserByIdPendingAccessReviewInstanceByIdDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecision.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

payload := userpendingaccessreviewinstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionClient.CreateUserByIdPendingAccessReviewInstanceByIdDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecision.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

payload := userpendingaccessreviewinstancedecision.CreateUserByIdPendingAccessReviewInstanceByIdDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdDecisionRecordAllDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionClient.DeleteUserByIdPendingAccessReviewInstanceByIdDecisionById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecision.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeleteUserByIdPendingAccessReviewInstanceByIdDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionClient.GetUserByIdPendingAccessReviewInstanceByIdDecisionById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecision.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetUserByIdPendingAccessReviewInstanceByIdDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionClient.GetUserByIdPendingAccessReviewInstanceByIdDecisionCount`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecision.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.GetUserByIdPendingAccessReviewInstanceByIdDecisionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionClient.ListUserByIdPendingAccessReviewInstanceByIdDecisions`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecision.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

// alternatively `client.ListUserByIdPendingAccessReviewInstanceByIdDecisions(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdPendingAccessReviewInstanceByIdDecisionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserPendingAccessReviewInstanceDecisionClient.UpdateUserByIdPendingAccessReviewInstanceByIdDecisionById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstancedecision.NewUserPendingAccessReviewInstanceDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := userpendingaccessreviewinstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdateUserByIdPendingAccessReviewInstanceByIdDecisionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
