
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancestagedecisioninstance` Documentation

The `accessreviewdefinitioninstancestagedecisioninstance` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancestagedecisioninstance"
```


### Client Initialization

```go
client := accessreviewdefinitioninstancestagedecisioninstance.NewAccessReviewDefinitionInstanceStageDecisionInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.AcceptAccessReviewDefinitionInstanceStageDecisionInstanceRecommendation`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.AcceptAccessReviewDefinitionInstanceStageDecisionInstanceRecommendation(ctx, id)
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
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecision(ctx, id)
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
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := accessreviewdefinitioninstancestagedecisioninstance.CreateAccessReviewDefinitionInstanceStageDecisionInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceStageDecisionInstanceBatchRecordDecision(ctx, id, payload)
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
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

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
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetAccessReviewDefinitionInstanceStageDecisionInstance(ctx, id, accessreviewdefinitioninstancestagedecisioninstance.DefaultGetAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.ResetAccessReviewDefinitionInstanceStageDecisionInstanceDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.ResetAccessReviewDefinitionInstanceStageDecisionInstanceDecision(ctx, id)
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
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.SendAccessReviewDefinitionInstanceStageDecisionInstanceReminder(ctx, id)
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
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopAccessReviewDefinitionInstanceStageDecisionInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceStageDecisionInstanceClient.StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecision(ctx, id)
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
id := accessreviewdefinitioninstancestagedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := accessreviewdefinitioninstancestagedecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateAccessReviewDefinitionInstanceStageDecisionInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
