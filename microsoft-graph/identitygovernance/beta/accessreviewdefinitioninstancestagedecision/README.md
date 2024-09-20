
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancestagedecision` Documentation

The `accessreviewdefinitioninstancestagedecision` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancestagedecision"
```


### Client Initialization

```go
client := accessreviewdefinitioninstancestagedecision.NewAccessReviewDefinitionInstanceStageDecisionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionClient.CreateAccessReviewDefinitionInstanceStageDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId")

payload := accessreviewdefinitioninstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceStageDecision(ctx, id, payload, accessreviewdefinitioninstancestagedecision.DefaultCreateAccessReviewDefinitionInstanceStageDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionClient.CreateAccessReviewDefinitionInstanceStageDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId")

payload := accessreviewdefinitioninstancestagedecision.CreateAccessReviewDefinitionInstanceStageDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceStageDecisionRecordAllDecision(ctx, id, payload, accessreviewdefinitioninstancestagedecision.DefaultCreateAccessReviewDefinitionInstanceStageDecisionRecordAllDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionClient.DeleteAccessReviewDefinitionInstanceStageDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.DeleteAccessReviewDefinitionInstanceStageDecision(ctx, id, accessreviewdefinitioninstancestagedecision.DefaultDeleteAccessReviewDefinitionInstanceStageDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionClient.GetAccessReviewDefinitionInstanceStageDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.GetAccessReviewDefinitionInstanceStageDecision(ctx, id, accessreviewdefinitioninstancestagedecision.DefaultGetAccessReviewDefinitionInstanceStageDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionClient.GetAccessReviewDefinitionInstanceStageDecisionsCount`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId")

read, err := client.GetAccessReviewDefinitionInstanceStageDecisionsCount(ctx, id, accessreviewdefinitioninstancestagedecision.DefaultGetAccessReviewDefinitionInstanceStageDecisionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionClient.ListAccessReviewDefinitionInstanceStageDecisions`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId")

// alternatively `client.ListAccessReviewDefinitionInstanceStageDecisions(ctx, id, accessreviewdefinitioninstancestagedecision.DefaultListAccessReviewDefinitionInstanceStageDecisionsOperationOptions())` can be used to do batched pagination
items, err := client.ListAccessReviewDefinitionInstanceStageDecisionsComplete(ctx, id, accessreviewdefinitioninstancestagedecision.DefaultListAccessReviewDefinitionInstanceStageDecisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionClient.UpdateAccessReviewDefinitionInstanceStageDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

payload := accessreviewdefinitioninstancestagedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdateAccessReviewDefinitionInstanceStageDecision(ctx, id, payload, accessreviewdefinitioninstancestagedecision.DefaultUpdateAccessReviewDefinitionInstanceStageDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
