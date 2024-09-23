
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdecisioninstancestagedecision` Documentation

The `accessreviewdecisioninstancestagedecision` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdecisioninstancestagedecision"
```


### Client Initialization

```go
client := accessreviewdecisioninstancestagedecision.NewAccessReviewDecisionInstanceStageDecisionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDecisionInstanceStageDecisionClient.CreateAccessReviewDecisionInstanceStageDecision`

```go
ctx := context.TODO()
id := accessreviewdecisioninstancestagedecision.NewIdentityGovernanceAccessReviewDecisionIdInstanceStageID("accessReviewInstanceDecisionItemId", "accessReviewStageId")

payload := accessreviewdecisioninstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateAccessReviewDecisionInstanceStageDecision(ctx, id, payload, accessreviewdecisioninstancestagedecision.DefaultCreateAccessReviewDecisionInstanceStageDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceStageDecisionClient.CreateAccessReviewDecisionInstanceStageDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := accessreviewdecisioninstancestagedecision.NewIdentityGovernanceAccessReviewDecisionIdInstanceStageID("accessReviewInstanceDecisionItemId", "accessReviewStageId")

payload := accessreviewdecisioninstancestagedecision.CreateAccessReviewDecisionInstanceStageDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDecisionInstanceStageDecisionRecordAllDecision(ctx, id, payload, accessreviewdecisioninstancestagedecision.DefaultCreateAccessReviewDecisionInstanceStageDecisionRecordAllDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceStageDecisionClient.GetAccessReviewDecisionInstanceStageDecisionsCount`

```go
ctx := context.TODO()
id := accessreviewdecisioninstancestagedecision.NewIdentityGovernanceAccessReviewDecisionIdInstanceStageID("accessReviewInstanceDecisionItemId", "accessReviewStageId")

read, err := client.GetAccessReviewDecisionInstanceStageDecisionsCount(ctx, id, accessreviewdecisioninstancestagedecision.DefaultGetAccessReviewDecisionInstanceStageDecisionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceStageDecisionClient.ListAccessReviewDecisionInstanceStageDecisions`

```go
ctx := context.TODO()
id := accessreviewdecisioninstancestagedecision.NewIdentityGovernanceAccessReviewDecisionIdInstanceStageID("accessReviewInstanceDecisionItemId", "accessReviewStageId")

// alternatively `client.ListAccessReviewDecisionInstanceStageDecisions(ctx, id, accessreviewdecisioninstancestagedecision.DefaultListAccessReviewDecisionInstanceStageDecisionsOperationOptions())` can be used to do batched pagination
items, err := client.ListAccessReviewDecisionInstanceStageDecisionsComplete(ctx, id, accessreviewdecisioninstancestagedecision.DefaultListAccessReviewDecisionInstanceStageDecisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
