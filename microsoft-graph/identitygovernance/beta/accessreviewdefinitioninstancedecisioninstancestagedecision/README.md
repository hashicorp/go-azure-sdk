
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancedecisioninstancestagedecision` Documentation

The `accessreviewdefinitioninstancedecisioninstancestagedecision` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancedecisioninstancestagedecision"
```


### Client Initialization

```go
client := accessreviewdefinitioninstancedecisioninstancestagedecision.NewAccessReviewDefinitionInstanceDecisionInstanceStageDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceStageDecisionClient.CreateAccessReviewDefinitionInstanceDecisionInstanceStageDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstancestagedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue")

payload := accessreviewdefinitioninstancedecisioninstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceDecisionInstanceStageDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceStageDecisionClient.CreateAccessReviewDefinitionInstanceDecisionInstanceStageDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstancestagedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue")

payload := accessreviewdefinitioninstancedecisioninstancestagedecision.CreateAccessReviewDefinitionInstanceDecisionInstanceStageDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceDecisionInstanceStageDecisionRecordAllDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceStageDecisionClient.GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCount`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstancestagedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue")

read, err := client.GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCount(ctx, id, accessreviewdefinitioninstancedecisioninstancestagedecision.DefaultGetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceStageDecisionClient.ListAccessReviewDefinitionInstanceDecisionInstanceStageDecisions`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstancestagedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue")

// alternatively `client.ListAccessReviewDefinitionInstanceDecisionInstanceStageDecisions(ctx, id, accessreviewdefinitioninstancedecisioninstancestagedecision.DefaultListAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsOperationOptions())` can be used to do batched pagination
items, err := client.ListAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsComplete(ctx, id, accessreviewdefinitioninstancedecisioninstancestagedecision.DefaultListAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
