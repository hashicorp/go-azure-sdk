
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstance` Documentation

The `pendingaccessreviewinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstance"
```


### Client Initialization

```go
client := pendingaccessreviewinstance.NewPendingAccessReviewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceClient.AcceptPendingAccessReviewInstanceRecommendation`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.AcceptPendingAccessReviewInstanceRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceClient.CreatePendingAccessReviewInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserID("userIdValue")

payload := pendingaccessreviewinstance.AccessReviewInstance{
	// ...
}


read, err := client.CreatePendingAccessReviewInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceClient.CreatePendingAccessReviewInstanceApplyDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.CreatePendingAccessReviewInstanceApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceClient.CreatePendingAccessReviewInstanceBatchRecordDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

payload := pendingaccessreviewinstance.CreatePendingAccessReviewInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceBatchRecordDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceClient.DeletePendingAccessReviewInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.DeletePendingAccessReviewInstance(ctx, id, pendingaccessreviewinstance.DefaultDeletePendingAccessReviewInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceClient.GetPendingAccessReviewInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.GetPendingAccessReviewInstance(ctx, id, pendingaccessreviewinstance.DefaultGetPendingAccessReviewInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceClient.GetPendingAccessReviewInstancesCount`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserID("userIdValue")

read, err := client.GetPendingAccessReviewInstancesCount(ctx, id, pendingaccessreviewinstance.DefaultGetPendingAccessReviewInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceClient.ListPendingAccessReviewInstances`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserID("userIdValue")

// alternatively `client.ListPendingAccessReviewInstances(ctx, id, pendingaccessreviewinstance.DefaultListPendingAccessReviewInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListPendingAccessReviewInstancesComplete(ctx, id, pendingaccessreviewinstance.DefaultListPendingAccessReviewInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PendingAccessReviewInstanceClient.ResetPendingAccessReviewInstanceDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.ResetPendingAccessReviewInstanceDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceClient.SendPendingAccessReviewInstanceReminder`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.SendPendingAccessReviewInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceClient.StopPendingAccessReviewInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.StopPendingAccessReviewInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceClient.StopPendingAccessReviewInstanceApplyDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.StopPendingAccessReviewInstanceApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceClient.UpdatePendingAccessReviewInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

payload := pendingaccessreviewinstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdatePendingAccessReviewInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
