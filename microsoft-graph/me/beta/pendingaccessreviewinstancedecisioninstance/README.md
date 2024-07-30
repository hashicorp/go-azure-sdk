
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstancedecisioninstance` Documentation

The `pendingaccessreviewinstancedecisioninstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstancedecisioninstance"
```


### Client Initialization

```go
client := pendingaccessreviewinstancedecisioninstance.NewPendingAccessReviewInstanceDecisionInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.CreatePendingAccessReviewInstanceDecisionInstanceAcceptRecommendation`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreatePendingAccessReviewInstanceDecisionInstanceAcceptRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.CreatePendingAccessReviewInstanceDecisionInstanceApplyDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreatePendingAccessReviewInstanceDecisionInstanceApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.CreatePendingAccessReviewInstanceDecisionInstanceBatchRecordDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := pendingaccessreviewinstancedecisioninstance.CreatePendingAccessReviewInstanceDecisionInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceDecisionInstanceBatchRecordDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.CreatePendingAccessReviewInstanceDecisionInstanceResetDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreatePendingAccessReviewInstanceDecisionInstanceResetDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.CreatePendingAccessReviewInstanceDecisionInstanceSendReminder`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreatePendingAccessReviewInstanceDecisionInstanceSendReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.DeletePendingAccessReviewInstanceDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeletePendingAccessReviewInstanceDecisionInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.GetPendingAccessReviewInstanceDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetPendingAccessReviewInstanceDecisionInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.StopMePendingAccessReviewInstanceDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopMePendingAccessReviewInstanceDecisionInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.StopMePendingAccessReviewInstanceDecisionInstanceApplyDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopMePendingAccessReviewInstanceDecisionInstanceApplyDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.UpdatePendingAccessReviewInstanceDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := pendingaccessreviewinstancedecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdatePendingAccessReviewInstanceDecisionInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
