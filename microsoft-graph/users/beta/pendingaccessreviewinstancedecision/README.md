
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedecision` Documentation

The `pendingaccessreviewinstancedecision` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedecision"
```


### Client Initialization

```go
client := pendingaccessreviewinstancedecision.NewPendingAccessReviewInstanceDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceDecisionClient.CreatePendingAccessReviewInstanceDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecision.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

payload := pendingaccessreviewinstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionClient.CreatePendingAccessReviewInstanceDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecision.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

payload := pendingaccessreviewinstancedecision.CreatePendingAccessReviewInstanceDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreatePendingAccessReviewInstanceDecisionRecordAllDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionClient.DeletePendingAccessReviewInstanceDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecision.NewUserIdPendingAccessReviewInstanceIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeletePendingAccessReviewInstanceDecision(ctx, id, pendingaccessreviewinstancedecision.DefaultDeletePendingAccessReviewInstanceDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionClient.GetPendingAccessReviewInstanceDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecision.NewUserIdPendingAccessReviewInstanceIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetPendingAccessReviewInstanceDecision(ctx, id, pendingaccessreviewinstancedecision.DefaultGetPendingAccessReviewInstanceDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionClient.GetPendingAccessReviewInstanceDecisionsCount`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecision.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

read, err := client.GetPendingAccessReviewInstanceDecisionsCount(ctx, id, pendingaccessreviewinstancedecision.DefaultGetPendingAccessReviewInstanceDecisionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionClient.ListPendingAccessReviewInstanceDecisions`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecision.NewUserIdPendingAccessReviewInstanceID("userIdValue", "accessReviewInstanceIdValue")

// alternatively `client.ListPendingAccessReviewInstanceDecisions(ctx, id, pendingaccessreviewinstancedecision.DefaultListPendingAccessReviewInstanceDecisionsOperationOptions())` can be used to do batched pagination
items, err := client.ListPendingAccessReviewInstanceDecisionsComplete(ctx, id, pendingaccessreviewinstancedecision.DefaultListPendingAccessReviewInstanceDecisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionClient.UpdatePendingAccessReviewInstanceDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecision.NewUserIdPendingAccessReviewInstanceIdDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := pendingaccessreviewinstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdatePendingAccessReviewInstanceDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
