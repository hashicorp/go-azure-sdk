
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancedecisioninstance` Documentation

The `accessreviewdefinitioninstancedecisioninstance` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancedecisioninstance"
```


### Client Initialization

```go
client := accessreviewdefinitioninstancedecisioninstance.NewAccessReviewDefinitionInstanceDecisionInstanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendations`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendations(ctx, id, accessreviewdefinitioninstancedecisioninstance.DefaultAcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendationsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.CreateAccessReviewDefinitionInstanceDecisionInstanceApplyDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.CreateAccessReviewDefinitionInstanceDecisionInstanceApplyDecision(ctx, id, accessreviewdefinitioninstancedecisioninstance.DefaultCreateAccessReviewDefinitionInstanceDecisionInstanceApplyDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.CreateAccessReviewDefinitionInstanceDecisionInstanceBatchRecordDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

payload := accessreviewdefinitioninstancedecisioninstance.CreateAccessReviewDefinitionInstanceDecisionInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceDecisionInstanceBatchRecordDecision(ctx, id, payload, accessreviewdefinitioninstancedecisioninstance.DefaultCreateAccessReviewDefinitionInstanceDecisionInstanceBatchRecordDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.DeleteAccessReviewDefinitionInstanceDecisionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.DeleteAccessReviewDefinitionInstanceDecisionInstance(ctx, id, accessreviewdefinitioninstancedecisioninstance.DefaultDeleteAccessReviewDefinitionInstanceDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.GetAccessReviewDefinitionInstanceDecisionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.GetAccessReviewDefinitionInstanceDecisionInstance(ctx, id, accessreviewdefinitioninstancedecisioninstance.DefaultGetAccessReviewDefinitionInstanceDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.ResetAccessReviewDefinitionInstanceDecisionInstanceDecisions`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.ResetAccessReviewDefinitionInstanceDecisionInstanceDecisions(ctx, id, accessreviewdefinitioninstancedecisioninstance.DefaultResetAccessReviewDefinitionInstanceDecisionInstanceDecisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.SendAccessReviewDefinitionInstanceDecisionInstanceReminder`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.SendAccessReviewDefinitionInstanceDecisionInstanceReminder(ctx, id, accessreviewdefinitioninstancedecisioninstance.DefaultSendAccessReviewDefinitionInstanceDecisionInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.StopAccessReviewDefinitionInstanceDecisionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.StopAccessReviewDefinitionInstanceDecisionInstance(ctx, id, accessreviewdefinitioninstancedecisioninstance.DefaultStopAccessReviewDefinitionInstanceDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.StopAccessReviewDefinitionInstanceDecisionInstanceApplyDecisions`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

read, err := client.StopAccessReviewDefinitionInstanceDecisionInstanceApplyDecisions(ctx, id, accessreviewdefinitioninstancedecisioninstance.DefaultStopAccessReviewDefinitionInstanceDecisionInstanceApplyDecisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.UpdateAccessReviewDefinitionInstanceDecisionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

payload := accessreviewdefinitioninstancedecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateAccessReviewDefinitionInstanceDecisionInstance(ctx, id, payload, accessreviewdefinitioninstancedecisioninstance.DefaultUpdateAccessReviewDefinitionInstanceDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
