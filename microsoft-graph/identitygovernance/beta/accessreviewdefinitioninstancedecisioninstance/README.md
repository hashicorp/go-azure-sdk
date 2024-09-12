
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancedecisioninstance` Documentation

The `accessreviewdefinitioninstancedecisioninstance` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancedecisioninstance"
```


### Client Initialization

```go
client := accessreviewdefinitioninstancedecisioninstance.NewAccessReviewDefinitionInstanceDecisionInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendation`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendation(ctx, id)
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
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateAccessReviewDefinitionInstanceDecisionInstanceApplyDecision(ctx, id)
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
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := accessreviewdefinitioninstancedecisioninstance.CreateAccessReviewDefinitionInstanceDecisionInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceDecisionInstanceBatchRecordDecision(ctx, id, payload)
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
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

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
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetAccessReviewDefinitionInstanceDecisionInstance(ctx, id, accessreviewdefinitioninstancedecisioninstance.DefaultGetAccessReviewDefinitionInstanceDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.ResetAccessReviewDefinitionInstanceDecisionInstanceDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.ResetAccessReviewDefinitionInstanceDecisionInstanceDecision(ctx, id)
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
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.SendAccessReviewDefinitionInstanceDecisionInstanceReminder(ctx, id)
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
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopAccessReviewDefinitionInstanceDecisionInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionInstanceClient.StopAccessReviewDefinitionInstanceDecisionInstanceApplyDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.StopAccessReviewDefinitionInstanceDecisionInstanceApplyDecision(ctx, id)
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
id := accessreviewdefinitioninstancedecisioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := accessreviewdefinitioninstancedecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateAccessReviewDefinitionInstanceDecisionInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
