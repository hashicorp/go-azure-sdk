
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreviewdefinitioninstance` Documentation

The `accessreviewdefinitioninstance` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreviewdefinitioninstance"
```


### Client Initialization

```go
client := accessreviewdefinitioninstance.NewAccessReviewDefinitionInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDefinitionInstanceClient.CreateAccessReviewDefinitionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionID("accessReviewScheduleDefinitionIdValue")

payload := accessreviewdefinitioninstance.AccessReviewInstance{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.CreateAccessReviewDefinitionInstanceAcceptRecommendation`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

read, err := client.CreateAccessReviewDefinitionInstanceAcceptRecommendation(ctx, id)
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
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

read, err := client.CreateAccessReviewDefinitionInstanceApplyDecision(ctx, id)
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
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

payload := accessreviewdefinitioninstance.CreateAccessReviewDefinitionInstanceBatchRecordDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceBatchRecordDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.CreateAccessReviewDefinitionInstanceResetDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

read, err := client.CreateAccessReviewDefinitionInstanceResetDecision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.CreateAccessReviewDefinitionInstanceSendReminder`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

read, err := client.CreateAccessReviewDefinitionInstanceSendReminder(ctx, id)
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
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

read, err := client.DeleteAccessReviewDefinitionInstance(ctx, id)
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
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

read, err := client.GetAccessReviewDefinitionInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.GetAccessReviewDefinitionInstanceCount`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionID("accessReviewScheduleDefinitionIdValue")

read, err := client.GetAccessReviewDefinitionInstanceCount(ctx, id)
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
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionID("accessReviewScheduleDefinitionIdValue")

// alternatively `client.ListAccessReviewDefinitionInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListAccessReviewDefinitionInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AccessReviewDefinitionInstanceClient.StopIdentityGovernanceAccessReviewDefinitionInstance`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

read, err := client.StopIdentityGovernanceAccessReviewDefinitionInstance(ctx, id)
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
id := accessreviewdefinitioninstance.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

payload := accessreviewdefinitioninstance.AccessReviewInstance{
	// ...
}


read, err := client.UpdateAccessReviewDefinitionInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
