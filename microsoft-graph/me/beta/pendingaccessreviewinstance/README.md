
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstance` Documentation

The `pendingaccessreviewinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstance"
```


### Client Initialization

```go
client := pendingaccessreviewinstance.NewPendingAccessReviewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceClient.AcceptPendingAccessReviewInstanceRecommendation`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

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

payload := pendingaccessreviewinstance.AccessReviewInstance{
	// ...
}


read, err := client.CreatePendingAccessReviewInstance(ctx, payload)
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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

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


### Example Usage: `PendingAccessReviewInstanceClient.ResetPendingAccessReviewInstanceDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

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
id := pendingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

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
