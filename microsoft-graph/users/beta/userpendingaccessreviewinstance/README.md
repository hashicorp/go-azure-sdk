
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstance` Documentation

The `userpendingaccessreviewinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpendingaccessreviewinstance"
```


### Client Initialization

```go
client := userpendingaccessreviewinstance.NewUserPendingAccessReviewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserPendingAccessReviewInstanceClient.CreateUserByIdPendingAccessReviewInstance`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserID("userIdValue")

payload := userpendingaccessreviewinstance.AccessReviewInstance{
	// ...
}


read, err := client.CreateUserByIdPendingAccessReviewInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdAcceptRecommendation`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdAcceptRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdApplyDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdBatchRecordDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

payload := userpendingaccessreviewinstance.CreateUserByIdPendingAccessReviewInstanceByIdBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdBatchRecordDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdResetDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdResetDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceClient.CreateUserByIdPendingAccessReviewInstanceByIdSendReminder`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.CreateUserByIdPendingAccessReviewInstanceByIdSendReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceClient.DeleteUserByIdPendingAccessReviewInstanceById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.DeleteUserByIdPendingAccessReviewInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceClient.GetUserByIdPendingAccessReviewInstanceById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.GetUserByIdPendingAccessReviewInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceClient.GetUserByIdPendingAccessReviewInstanceCount`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserID("userIdValue")

read, err := client.GetUserByIdPendingAccessReviewInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceClient.ListUserByIdPendingAccessReviewInstances`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserID("userIdValue")

// alternatively `client.ListUserByIdPendingAccessReviewInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdPendingAccessReviewInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserPendingAccessReviewInstanceClient.StopUserByIdPendingAccessReviewInstanceById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.StopUserByIdPendingAccessReviewInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceClient.StopUserByIdPendingAccessReviewInstanceByIdApplyDecision`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.StopUserByIdPendingAccessReviewInstanceByIdApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPendingAccessReviewInstanceClient.UpdateUserByIdPendingAccessReviewInstanceById`

```go
ctx := context.TODO()
id := userpendingaccessreviewinstance.NewUserPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

payload := userpendingaccessreviewinstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateUserByIdPendingAccessReviewInstanceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
