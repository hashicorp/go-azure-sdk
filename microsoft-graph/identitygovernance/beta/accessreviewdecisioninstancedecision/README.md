
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdecisioninstancedecision` Documentation

The `accessreviewdecisioninstancedecision` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdecisioninstancedecision"
```


### Client Initialization

```go
client := accessreviewdecisioninstancedecision.NewAccessReviewDecisionInstanceDecisionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDecisionInstanceDecisionClient.CreateAccessReviewDecisionInstanceDecision`

```go
ctx := context.TODO()
id := accessreviewdecisioninstancedecision.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

payload := accessreviewdecisioninstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateAccessReviewDecisionInstanceDecision(ctx, id, payload, accessreviewdecisioninstancedecision.DefaultCreateAccessReviewDecisionInstanceDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceDecisionClient.CreateAccessReviewDecisionInstanceDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := accessreviewdecisioninstancedecision.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

payload := accessreviewdecisioninstancedecision.CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDecisionInstanceDecisionRecordAllDecision(ctx, id, payload, accessreviewdecisioninstancedecision.DefaultCreateAccessReviewDecisionInstanceDecisionRecordAllDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceDecisionClient.GetAccessReviewDecisionInstanceDecisionsCount`

```go
ctx := context.TODO()
id := accessreviewdecisioninstancedecision.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

read, err := client.GetAccessReviewDecisionInstanceDecisionsCount(ctx, id, accessreviewdecisioninstancedecision.DefaultGetAccessReviewDecisionInstanceDecisionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceDecisionClient.ListAccessReviewDecisionInstanceDecisions`

```go
ctx := context.TODO()
id := accessreviewdecisioninstancedecision.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

// alternatively `client.ListAccessReviewDecisionInstanceDecisions(ctx, id, accessreviewdecisioninstancedecision.DefaultListAccessReviewDecisionInstanceDecisionsOperationOptions())` can be used to do batched pagination
items, err := client.ListAccessReviewDecisionInstanceDecisionsComplete(ctx, id, accessreviewdecisioninstancedecision.DefaultListAccessReviewDecisionInstanceDecisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
