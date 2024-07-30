
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstancedecision` Documentation

The `pendingaccessreviewinstancedecision` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/pendingaccessreviewinstancedecision"
```


### Client Initialization

```go
client := pendingaccessreviewinstancedecision.NewPendingAccessReviewInstanceDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PendingAccessReviewInstanceDecisionClient.CreatePendingAccessReviewInstanceDecision`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

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
id := pendingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

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
id := pendingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeletePendingAccessReviewInstanceDecision(ctx, id)
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
id := pendingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetPendingAccessReviewInstanceDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PendingAccessReviewInstanceDecisionClient.GetPendingAccessReviewInstanceDecisionCount`

```go
ctx := context.TODO()
id := pendingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

read, err := client.GetPendingAccessReviewInstanceDecisionCount(ctx, id)
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
id := pendingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

// alternatively `client.ListPendingAccessReviewInstanceDecisions(ctx, id)` can be used to do batched pagination
items, err := client.ListPendingAccessReviewInstanceDecisionsComplete(ctx, id)
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
id := pendingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceIdDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

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
