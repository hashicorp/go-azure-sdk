
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstancedecisioninstance` Documentation

The `mependingaccessreviewinstancedecisioninstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstancedecisioninstance"
```


### Client Initialization

```go
client := mependingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceDecisionInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceClient.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceAcceptRecommendation`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceAcceptRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceClient.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceApplyDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceClient.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceBatchRecordDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := mependingaccessreviewinstancedecisioninstance.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceBatchRecordDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceClient.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceResetDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceResetDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceClient.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceSendReminder`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateMePendingAccessReviewInstanceByIdDecisionByIdInstanceSendReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceClient.DeleteMePendingAccessReviewInstanceByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeleteMePendingAccessReviewInstanceByIdDecisionByIdInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceClient.GetMePendingAccessReviewInstanceByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetMePendingAccessReviewInstanceByIdDecisionByIdInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceClient.StopMePendingAccessReviewInstanceByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopMePendingAccessReviewInstanceByIdDecisionByIdInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceClient.StopMePendingAccessReviewInstanceByIdDecisionByIdInstanceApplyDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopMePendingAccessReviewInstanceByIdDecisionByIdInstanceApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionInstanceClient.UpdateMePendingAccessReviewInstanceByIdDecisionByIdInstance`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := mependingaccessreviewinstancedecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateMePendingAccessReviewInstanceByIdDecisionByIdInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
