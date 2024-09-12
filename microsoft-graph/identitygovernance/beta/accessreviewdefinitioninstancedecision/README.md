
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancedecision` Documentation

The `accessreviewdefinitioninstancedecision` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/accessreviewdefinitioninstancedecision"
```


### Client Initialization

```go
client := accessreviewdefinitioninstancedecision.NewAccessReviewDefinitionInstanceDecisionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionClient.CreateAccessReviewDefinitionInstanceDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

payload := accessreviewdefinitioninstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionClient.CreateAccessReviewDefinitionInstanceDecisionRecordAllDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

payload := accessreviewdefinitioninstancedecision.CreateAccessReviewDefinitionInstanceDecisionRecordAllDecisionRequest{
	// ...
}


read, err := client.CreateAccessReviewDefinitionInstanceDecisionRecordAllDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionClient.DeleteAccessReviewDefinitionInstanceDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.DeleteAccessReviewDefinitionInstanceDecision(ctx, id, accessreviewdefinitioninstancedecision.DefaultDeleteAccessReviewDefinitionInstanceDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionClient.GetAccessReviewDefinitionInstanceDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

read, err := client.GetAccessReviewDefinitionInstanceDecision(ctx, id, accessreviewdefinitioninstancedecision.DefaultGetAccessReviewDefinitionInstanceDecisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionClient.GetAccessReviewDefinitionInstanceDecisionsCount`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

read, err := client.GetAccessReviewDefinitionInstanceDecisionsCount(ctx, id, accessreviewdefinitioninstancedecision.DefaultGetAccessReviewDefinitionInstanceDecisionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionClient.ListAccessReviewDefinitionInstanceDecisions`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue")

// alternatively `client.ListAccessReviewDefinitionInstanceDecisions(ctx, id, accessreviewdefinitioninstancedecision.DefaultListAccessReviewDefinitionInstanceDecisionsOperationOptions())` can be used to do batched pagination
items, err := client.ListAccessReviewDefinitionInstanceDecisionsComplete(ctx, id, accessreviewdefinitioninstancedecision.DefaultListAccessReviewDefinitionInstanceDecisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AccessReviewDefinitionInstanceDecisionClient.UpdateAccessReviewDefinitionInstanceDecision`

```go
ctx := context.TODO()
id := accessreviewdefinitioninstancedecision.NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue")

payload := accessreviewdefinitioninstancedecision.AccessReviewInstanceDecisionItem{
	// ...
}


read, err := client.UpdateAccessReviewDefinitionInstanceDecision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
