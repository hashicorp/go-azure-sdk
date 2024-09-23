
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstance` Documentation

The `pendingaccessreviewinstance` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstance"
```


### Client Initialization

```go
client := pendingaccessreviewinstance.NewPendingAccessReviewInstanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceClient.AcceptPendingAccessReviewInstanceRecommendations`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceId")

read, err := client.AcceptPendingAccessReviewInstanceRecommendations(ctx, id, pendingaccessreviewinstance.DefaultAcceptPendingAccessReviewInstanceRecommendationsOperationOptions())
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

payload := pendingaccessreviewinstance.AccessReviewInstance{
	// ...
}


read, err := client.CreatePendingAccessReviewInstance(ctx, payload, pendingaccessreviewinstance.DefaultCreatePendingAccessReviewInstanceOperationOptions())
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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceId")

read, err := client.CreatePendingAccessReviewInstanceApplyDecision(ctx, id, pendingaccessreviewinstance.DefaultCreatePendingAccessReviewInstanceApplyDecisionOperationOptions())
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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceId")

payload := pendingaccessreviewinstance.CreatePendingAccessReviewInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceBatchRecordDecision(ctx, id, payload, pendingaccessreviewinstance.DefaultCreatePendingAccessReviewInstanceBatchRecordDecisionOperationOptions())
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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceId")

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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceId")

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


read, err := client.GetPendingAccessReviewInstancesCount(ctx, pendingaccessreviewinstance.DefaultGetPendingAccessReviewInstancesCountOperationOptions())
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


// alternatively `client.ListPendingAccessReviewInstances(ctx, pendingaccessreviewinstance.DefaultListPendingAccessReviewInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListPendingAccessReviewInstancesComplete(ctx, pendingaccessreviewinstance.DefaultListPendingAccessReviewInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PendingAccessReviewInstanceClient.ResetPendingAccessReviewInstanceDecisions`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceId")

read, err := client.ResetPendingAccessReviewInstanceDecisions(ctx, id, pendingaccessreviewinstance.DefaultResetPendingAccessReviewInstanceDecisionsOperationOptions())
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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceId")

read, err := client.SendPendingAccessReviewInstanceReminder(ctx, id, pendingaccessreviewinstance.DefaultSendPendingAccessReviewInstanceReminderOperationOptions())
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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceId")

read, err := client.StopPendingAccessReviewInstance(ctx, id, pendingaccessreviewinstance.DefaultStopPendingAccessReviewInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceClient.StopPendingAccessReviewInstanceApplyDecisions`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceId")

read, err := client.StopPendingAccessReviewInstanceApplyDecisions(ctx, id, pendingaccessreviewinstance.DefaultStopPendingAccessReviewInstanceApplyDecisionsOperationOptions())
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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceId")

payload := pendingaccessreviewinstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdatePendingAccessReviewInstance(ctx, id, payload, pendingaccessreviewinstance.DefaultUpdatePendingAccessReviewInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
