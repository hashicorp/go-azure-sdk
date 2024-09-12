
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdecisioninstancedecision` Documentation

The `accessreviewdecisioninstancedecision` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdecisioninstancedecision"
```


### Client Initialization

```go
client := accessreviewdecisioninstancedecision.NewAccessReviewDecisionInstanceDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDecisionInstanceDecisionClient.CreateAccessReviewDecisionInstanceDecision`

```go
ctx := context.TODO()
id := accessreviewdecisioninstancedecision.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

payload := accessreviewdecisioninstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateAccessReviewDecisionInstanceDecision(ctx, id, payload)
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
id := accessreviewdecisioninstancedecision.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

payload := accessreviewdecisioninstancedecision.CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDecisionInstanceDecisionRecordAllDecision(ctx, id, payload)
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
id := accessreviewdecisioninstancedecision.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

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
id := accessreviewdecisioninstancedecision.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

// alternatively `client.ListAccessReviewDecisionInstanceDecisions(ctx, id, accessreviewdecisioninstancedecision.DefaultListAccessReviewDecisionInstanceDecisionsOperationOptions())` can be used to do batched pagination
items, err := client.ListAccessReviewDecisionInstanceDecisionsComplete(ctx, id, accessreviewdecisioninstancedecision.DefaultListAccessReviewDecisionInstanceDecisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
