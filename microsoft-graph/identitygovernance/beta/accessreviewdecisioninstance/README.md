
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdecisioninstance` Documentation

The `accessreviewdecisioninstance` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdecisioninstance"
```


### Client Initialization

```go
client := accessreviewdecisioninstance.NewAccessReviewDecisionInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDecisionInstanceClient.AcceptAccessReviewDecisionInstanceRecommendation`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

read, err := client.AcceptAccessReviewDecisionInstanceRecommendation(ctx, id)
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
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

read, err := client.CreateAccessReviewDecisionInstanceApplyDecision(ctx, id)
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
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

payload := accessreviewdecisioninstance.CreateAccessReviewDecisionInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDecisionInstanceBatchRecordDecision(ctx, id, payload)
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
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

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
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

read, err := client.GetAccessReviewDecisionInstance(ctx, id, accessreviewdecisioninstance.DefaultGetAccessReviewDecisionInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceClient.ResetAccessReviewDecisionInstanceDecision`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

read, err := client.ResetAccessReviewDecisionInstanceDecision(ctx, id)
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
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

read, err := client.SendAccessReviewDecisionInstanceReminder(ctx, id)
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
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

read, err := client.StopAccessReviewDecisionInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDecisionInstanceClient.StopAccessReviewDecisionInstanceApplyDecision`

```go
ctx := context.TODO()
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

read, err := client.StopAccessReviewDecisionInstanceApplyDecision(ctx, id)
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
id := accessreviewdecisioninstance.NewIdentityGovernanceAccessReviewDecisionID("accessReviewInstanceDecisionItemIdValue")

payload := accessreviewdecisioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateAccessReviewDecisionInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
