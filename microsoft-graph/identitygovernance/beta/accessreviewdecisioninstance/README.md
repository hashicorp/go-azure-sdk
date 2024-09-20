
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdecisioninstance` Documentation

The `accessreviewdecisioninstance` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdecisioninstance"
```


### Client Initialization

```go
client := accessreviewdecisioninstance.NewAccessReviewDecisionInstanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDecisionInstanceClient.AcceptAccessReviewDecisionInstanceRecommendations`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

read, err := client.AcceptAccessReviewDecisionInstanceRecommendations(ctx, id, accessreviewdecisioninstance.DefaultAcceptAccessReviewDecisionInstanceRecommendationsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceClient.CreateAccessReviewDecisionInstanceApplyDecision`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

read, err := client.CreateAccessReviewDecisionInstanceApplyDecision(ctx, id, accessreviewdecisioninstance.DefaultCreateAccessReviewDecisionInstanceApplyDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceClient.CreateAccessReviewDecisionInstanceBatchRecordDecision`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

payload := accessreviewdecisioninstance.CreateAccessReviewDecisionInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDecisionInstanceBatchRecordDecision(ctx, id, payload, accessreviewdecisioninstance.DefaultCreateAccessReviewDecisionInstanceBatchRecordDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceClient.DeleteAccessReviewDecisionInstance`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

read, err := client.DeleteAccessReviewDecisionInstance(ctx, id, accessreviewdecisioninstance.DefaultDeleteAccessReviewDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceClient.GetAccessReviewDecisionInstance`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

read, err := client.GetAccessReviewDecisionInstance(ctx, id, accessreviewdecisioninstance.DefaultGetAccessReviewDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceClient.ResetAccessReviewDecisionInstanceDecisions`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

read, err := client.ResetAccessReviewDecisionInstanceDecisions(ctx, id, accessreviewdecisioninstance.DefaultResetAccessReviewDecisionInstanceDecisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceClient.SendAccessReviewDecisionInstanceReminder`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

read, err := client.SendAccessReviewDecisionInstanceReminder(ctx, id, accessreviewdecisioninstance.DefaultSendAccessReviewDecisionInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceClient.StopAccessReviewDecisionInstance`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

read, err := client.StopAccessReviewDecisionInstance(ctx, id, accessreviewdecisioninstance.DefaultStopAccessReviewDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceClient.StopAccessReviewDecisionInstanceApplyDecisions`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

read, err := client.StopAccessReviewDecisionInstanceApplyDecisions(ctx, id, accessreviewdecisioninstance.DefaultStopAccessReviewDecisionInstanceApplyDecisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceClient.UpdateAccessReviewDecisionInstance`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemId")

payload := accessreviewdecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateAccessReviewDecisionInstance(ctx, id, payload, accessreviewdecisioninstance.DefaultUpdateAccessReviewDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
