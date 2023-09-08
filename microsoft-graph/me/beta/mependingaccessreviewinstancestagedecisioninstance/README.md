
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstancestagedecisioninstance` Documentation

The `mependingaccessreviewinstancestagedecisioninstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstancestagedecisioninstance"
```


### Client Initialization

```go
client := mependingaccessreviewinstancestagedecisioninstance.NewMePendingAccessReviewInstanceStageDecisionInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceClient.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceAcceptRecommendation`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstance.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceAcceptRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceClient.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceApplyDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstance.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceClient.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceBatchRecordDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstance.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := mependingaccessreviewinstancestagedecisioninstance.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceBatchRecordDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceClient.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceResetDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstance.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceResetDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceClient.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceSendReminder`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstance.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceSendReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceClient.DeleteMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstance.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeleteMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceClient.GetMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstance.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceClient.StopMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstance.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceClient.StopMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceApplyDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstance.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceStageDecisionInstanceClient.UpdateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancestagedecisioninstance.NewMePendingAccessReviewInstanceStageDecisionID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := mependingaccessreviewinstancestagedecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
