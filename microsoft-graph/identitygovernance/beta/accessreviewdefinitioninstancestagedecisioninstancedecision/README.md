
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancestagedecisioninstancedecision` Documentation

The `accessreviewdefinitioninstancestagedecisioninstancedecision` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancestagedecisioninstancedecision"
```


### Client Initialization

```go
client := accessreviewdefinitioninstancestagedecisioninstancedecision.NewAccessReviewDefinitionInstanceStageDecisionInstanceDecisionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceDecisionClient.CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstancedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

payload := accessreviewdefinitioninstancestagedecisioninstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecision(ctx, id, payload, accessreviewdefinitioninstancestagedecisioninstancedecision.DefaultCreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceDecisionClient.CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstancedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

payload := accessreviewdefinitioninstancestagedecisioninstancedecision.CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecision(ctx, id, payload, accessreviewdefinitioninstancestagedecisioninstancedecision.DefaultCreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceDecisionClient.GetAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCount`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstancedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.GetAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCount(ctx, id, accessreviewdefinitioninstancestagedecisioninstancedecision.DefaultGetAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceDecisionClient.ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisions`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstancedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

// alternatively `client.ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisions(ctx, id, accessreviewdefinitioninstancestagedecisioninstancedecision.DefaultListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions())` can be used to do batched pagination
items, err := client.ListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsComplete(ctx, id, accessreviewdefinitioninstancestagedecisioninstancedecision.DefaultListAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
