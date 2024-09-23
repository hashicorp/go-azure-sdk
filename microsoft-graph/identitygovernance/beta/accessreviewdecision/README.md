
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdecision` Documentation

The `accessreviewdecision` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdecision"
```


### Client Initialization

```go
client := accessreviewdecision.NewAccessReviewDecisionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDecisionClient.CreateAccessReviewDecision`

```go
ctx := context.TODO()

payload := accessreviewdecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateAccessReviewDecision(ctx, payload, accessreviewdecision.DefaultCreateAccessReviewDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionClient.CreateAccessReviewDecisionRecordAllDecision`

```go
ctx := context.TODO()

payload := accessreviewdecision.CreateAccessReviewDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDecisionRecordAllDecision(ctx, payload, accessreviewdecision.DefaultCreateAccessReviewDecisionRecordAllDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionClient.DeleteAccessReviewDecision`

```go
ctx := context.TODO()
id := accessreviewdecision.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

read, err := client.DeleteAccessReviewDecision(ctx, id, accessreviewdecision.DefaultDeleteAccessReviewDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionClient.GetAccessReviewDecision`

```go
ctx := context.TODO()
id := accessreviewdecision.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

read, err := client.GetAccessReviewDecision(ctx, id, accessreviewdecision.DefaultGetAccessReviewDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionClient.GetAccessReviewDecisionsCount`

```go
ctx := context.TODO()


read, err := client.GetAccessReviewDecisionsCount(ctx, accessreviewdecision.DefaultGetAccessReviewDecisionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionClient.ListAccessReviewDecisions`

```go
ctx := context.TODO()


// alternatively `client.ListAccessReviewDecisions(ctx, accessreviewdecision.DefaultListAccessReviewDecisionsOperationOptions())` can be used to do batched pagination
items, err := client.ListAccessReviewDecisionsComplete(ctx, accessreviewdecision.DefaultListAccessReviewDecisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AccessReviewDecisionClient.UpdateAccessReviewDecision`

```go
ctx := context.TODO()
id := accessreviewdecision.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

payload := accessreviewdecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdateAccessReviewDecision(ctx, id, payload, accessreviewdecision.DefaultUpdateAccessReviewDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
