
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstance` Documentation

The `mependingaccessreviewinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstance"
```


### Client Initialization

```go
client := mependingaccessreviewinstance.NewMePendingAccessReviewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MePendingAccessReviewInstanceClient.CreateMePendingAccessReviewInstance`

```go
ctx := context.TODO()

payload := mependingaccessreviewinstance.AccessReviewInstance{
	// ...
}


read, err := client.CreateMePendingAccessReviewInstance(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceClient.CreateMePendingAccessReviewInstanceByIdAcceptRecommendation`

```go
ctx := context.TODO()
id := mependingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

read, err := client.CreateMePendingAccessReviewInstanceByIdAcceptRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceClient.CreateMePendingAccessReviewInstanceByIdApplyDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

read, err := client.CreateMePendingAccessReviewInstanceByIdApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceClient.CreateMePendingAccessReviewInstanceByIdBatchRecordDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

payload := mependingaccessreviewinstance.CreateMePendingAccessReviewInstanceByIdBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateMePendingAccessReviewInstanceByIdBatchRecordDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceClient.CreateMePendingAccessReviewInstanceByIdResetDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

read, err := client.CreateMePendingAccessReviewInstanceByIdResetDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceClient.CreateMePendingAccessReviewInstanceByIdSendReminder`

```go
ctx := context.TODO()
id := mependingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

read, err := client.CreateMePendingAccessReviewInstanceByIdSendReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceClient.DeleteMePendingAccessReviewInstanceById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

read, err := client.DeleteMePendingAccessReviewInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceClient.GetMePendingAccessReviewInstanceById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

read, err := client.GetMePendingAccessReviewInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceClient.GetMePendingAccessReviewInstanceCount`

```go
ctx := context.TODO()


read, err := client.GetMePendingAccessReviewInstanceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceClient.ListMePendingAccessReviewInstances`

```go
ctx := context.TODO()


// alternatively `client.ListMePendingAccessReviewInstances(ctx)` can be used to do batched pagination
items, err := client.ListMePendingAccessReviewInstancesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MePendingAccessReviewInstanceClient.StopMePendingAccessReviewInstanceById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

read, err := client.StopMePendingAccessReviewInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceClient.StopMePendingAccessReviewInstanceByIdApplyDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

read, err := client.StopMePendingAccessReviewInstanceByIdApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceClient.UpdateMePendingAccessReviewInstanceById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstance.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

payload := mependingaccessreviewinstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateMePendingAccessReviewInstanceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
