
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


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.AcceptPendingAccessReviewInstanceDecisionInstanceRecommendation`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.AcceptPendingAccessReviewInstanceDecisionInstanceRecommendation(ctx, id)
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


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.DeletePendingAccessReviewInstanceDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeletePendingAccessReviewInstanceDecisionInstance(ctx, id, pendingaccessreviewinstancedecisioninstance.DefaultDeletePendingAccessReviewInstanceDecisionInstanceOperationOptions())
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

read, err := client.GetPendingAccessReviewInstanceDecisionInstance(ctx, id, pendingaccessreviewinstancedecisioninstance.DefaultGetPendingAccessReviewInstanceDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.ResetPendingAccessReviewInstanceDecisionInstanceDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.ResetPendingAccessReviewInstanceDecisionInstanceDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.SendPendingAccessReviewInstanceDecisionInstanceReminder`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.SendPendingAccessReviewInstanceDecisionInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.StopPendingAccessReviewInstanceDecisionInstance`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopPendingAccessReviewInstanceDecisionInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionInstanceClient.StopPendingAccessReviewInstanceDecisionInstanceApplyDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopPendingAccessReviewInstanceDecisionInstanceApplyDecision(ctx, id)
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
