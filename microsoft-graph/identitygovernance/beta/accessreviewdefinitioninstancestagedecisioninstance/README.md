
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancestagedecisioninstance` Documentation

The `accessreviewdefinitioninstancestagedecisioninstance` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancestagedecisioninstance"
```


### Client Initialization

```go
client := accessreviewdefinitioninstancestagedecisioninstance.NewAccessReviewDefinitionInstanceStageDecisionInstanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.AcceptAccessReviewDefinitionInstanceStageDecisionInstanceRecommendations`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.AcceptAccessReviewDefinitionInstanceStageDecisionInstanceRecommendations(ctx, id, accessreviewdefinitioninstancestagedecisioninstance.DefaultAcceptAccessReviewDefinitionInstanceStageDecisionInstanceRecommendationsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.CreateAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.CreateAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecision(ctx, id, accessreviewdefinitioninstancestagedecisioninstance.DefaultCreateAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.CreateAccessReviewDefinitionInstanceStageDecisionInstanceBatchRecordDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

payload := accessreviewdefinitioninstancestagedecisioninstance.CreateAccessReviewDefinitionInstanceStageDecisionInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceStageDecisionInstanceBatchRecordDecision(ctx, id, payload, accessreviewdefinitioninstancestagedecisioninstance.DefaultCreateAccessReviewDefinitionInstanceStageDecisionInstanceBatchRecordDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.DeleteAccessReviewDefinitionInstanceStageDecisionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.DeleteAccessReviewDefinitionInstanceStageDecisionInstance(ctx, id, accessreviewdefinitioninstancestagedecisioninstance.DefaultDeleteAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.GetAccessReviewDefinitionInstanceStageDecisionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.GetAccessReviewDefinitionInstanceStageDecisionInstance(ctx, id, accessreviewdefinitioninstancestagedecisioninstance.DefaultGetAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.ResetAccessReviewDefinitionInstanceStageDecisionInstanceDecisions`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.ResetAccessReviewDefinitionInstanceStageDecisionInstanceDecisions(ctx, id, accessreviewdefinitioninstancestagedecisioninstance.DefaultResetAccessReviewDefinitionInstanceStageDecisionInstanceDecisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.SendAccessReviewDefinitionInstanceStageDecisionInstanceReminder`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.SendAccessReviewDefinitionInstanceStageDecisionInstanceReminder(ctx, id, accessreviewdefinitioninstancestagedecisioninstance.DefaultSendAccessReviewDefinitionInstanceStageDecisionInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.StopAccessReviewDefinitionInstanceStageDecisionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.StopAccessReviewDefinitionInstanceStageDecisionInstance(ctx, id, accessreviewdefinitioninstancestagedecisioninstance.DefaultStopAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisions`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

read, err := client.StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisions(ctx, id, accessreviewdefinitioninstancestagedecisioninstance.DefaultStopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.UpdateAccessReviewDefinitionInstanceStageDecisionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewStageId", "accessReviewInstanceDecisionItemId")

payload := accessreviewdefinitioninstancestagedecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateAccessReviewDefinitionInstanceStageDecisionInstance(ctx, id, payload, accessreviewdefinitioninstancestagedecisioninstance.DefaultUpdateAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
