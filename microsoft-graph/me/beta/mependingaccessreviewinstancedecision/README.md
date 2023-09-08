
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstancedecision` Documentation

The `mependingaccessreviewinstancedecision` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mependingaccessreviewinstancedecision"
```


### Client Initialization

```go
client := mependingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MePendingAccessReviewInstanceDecisionClient.CreateMePendingAccessReviewInstanceByIdDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

payload := mependingaccessreviewinstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateMePendingAccessReviewInstanceByIdDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionClient.CreateMePendingAccessReviewInstanceByIdDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

payload := mependingaccessreviewinstancedecision.CreateMePendingAccessReviewInstanceByIdDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateMePendingAccessReviewInstanceByIdDecisionRecordAllDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionClient.DeleteMePendingAccessReviewInstanceByIdDecisionById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeleteMePendingAccessReviewInstanceByIdDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionClient.GetMePendingAccessReviewInstanceByIdDecisionById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetMePendingAccessReviewInstanceByIdDecisionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionClient.GetMePendingAccessReviewInstanceByIdDecisionCount`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

read, err := client.GetMePendingAccessReviewInstanceByIdDecisionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionClient.ListMePendingAccessReviewInstanceByIdDecisions`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceID("accessReviewInstanceIdValue")

// alternatively `client.ListMePendingAccessReviewInstanceByIdDecisions(ctx, id)` can be used to do batched pagination
items, err := client.ListMePendingAccessReviewInstanceByIdDecisionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MePendingAccessReviewInstanceDecisionClient.UpdateMePendingAccessReviewInstanceByIdDecisionById`

```go
ctx := context.TODO()
id := mependingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceDecisionID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := mependingaccessreviewinstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdateMePendingAccessReviewInstanceByIdDecisionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
