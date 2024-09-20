
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstance` Documentation

The `accessreviewdefinitioninstance` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstance"
```


### Client Initialization

```go
client := accessreviewdefinitioninstance.NewAccessReviewDefinitionInstanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDefinitionInstanceClient.AcceptAccessReviewDefinitionInstanceRecommendations`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionId", "accessReviewInstanceId")

read, err := client.AcceptAccessReviewDefinitionInstanceRecommendations(ctx, id, accessreviewdefinitioninstance.DefaultAcceptAccessReviewDefinitionInstanceRecommendationsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.CreateAccessReviewDefinitionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionID("accessReviewScheduleDefinitionId")

payload := accessreviewdefinitioninstance.AccessReviewInstance{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstance(ctx, id, payload, accessreviewdefinitioninstance.DefaultCreateAccessReviewDefinitionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.CreateAccessReviewDefinitionInstanceApplyDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionId", "accessReviewInstanceId")

read, err := client.CreateAccessReviewDefinitionInstanceApplyDecision(ctx, id, accessreviewdefinitioninstance.DefaultCreateAccessReviewDefinitionInstanceApplyDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.CreateAccessReviewDefinitionInstanceBatchRecordDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionId", "accessReviewInstanceId")

payload := accessreviewdefinitioninstance.CreateAccessReviewDefinitionInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceBatchRecordDecision(ctx, id, payload, accessreviewdefinitioninstance.DefaultCreateAccessReviewDefinitionInstanceBatchRecordDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.DeleteAccessReviewDefinitionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionId", "accessReviewInstanceId")

read, err := client.DeleteAccessReviewDefinitionInstance(ctx, id, accessreviewdefinitioninstance.DefaultDeleteAccessReviewDefinitionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.GetAccessReviewDefinitionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionId", "accessReviewInstanceId")

read, err := client.GetAccessReviewDefinitionInstance(ctx, id, accessreviewdefinitioninstance.DefaultGetAccessReviewDefinitionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.GetAccessReviewDefinitionInstancesCount`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionID("accessReviewScheduleDefinitionId")

read, err := client.GetAccessReviewDefinitionInstancesCount(ctx, id, accessreviewdefinitioninstance.DefaultGetAccessReviewDefinitionInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.ListAccessReviewDefinitionInstances`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionID("accessReviewScheduleDefinitionId")

// alternatively `client.ListAccessReviewDefinitionInstances(ctx, id, accessreviewdefinitioninstance.DefaultListAccessReviewDefinitionInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListAccessReviewDefinitionInstancesComplete(ctx, id, accessreviewdefinitioninstance.DefaultListAccessReviewDefinitionInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.ResetAccessReviewDefinitionInstanceDecisions`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionId", "accessReviewInstanceId")

read, err := client.ResetAccessReviewDefinitionInstanceDecisions(ctx, id, accessreviewdefinitioninstance.DefaultResetAccessReviewDefinitionInstanceDecisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.SendAccessReviewDefinitionInstanceReminder`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionId", "accessReviewInstanceId")

read, err := client.SendAccessReviewDefinitionInstanceReminder(ctx, id, accessreviewdefinitioninstance.DefaultSendAccessReviewDefinitionInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.StopAccessReviewDefinitionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionId", "accessReviewInstanceId")

read, err := client.StopAccessReviewDefinitionInstance(ctx, id, accessreviewdefinitioninstance.DefaultStopAccessReviewDefinitionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.StopAccessReviewDefinitionInstanceApplyDecisions`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionId", "accessReviewInstanceId")

read, err := client.StopAccessReviewDefinitionInstanceApplyDecisions(ctx, id, accessreviewdefinitioninstance.DefaultStopAccessReviewDefinitionInstanceApplyDecisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.UpdateAccessReviewDefinitionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionId", "accessReviewInstanceId")

payload := accessreviewdefinitioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateAccessReviewDefinitionInstance(ctx, id, payload, accessreviewdefinitioninstance.DefaultUpdateAccessReviewDefinitionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
