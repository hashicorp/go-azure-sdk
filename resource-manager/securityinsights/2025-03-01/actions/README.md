
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-03-01/actions` Documentation

The `actions` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-03-01/actions"
```


### Client Initialization

```go
client := actions.NewActionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ActionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := actions.NewActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "ruleId", "actionId")

payload := actions.ActionRequest{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ActionsClient.Delete`

```go
ctx := context.TODO()
id := actions.NewActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "ruleId", "actionId")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ActionsClient.Get`

```go
ctx := context.TODO()
id := actions.NewActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "ruleId", "actionId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ActionsClient.ListByAlertRule`

```go
ctx := context.TODO()
id := actions.NewAlertRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "ruleId")

// alternatively `client.ListByAlertRule(ctx, id)` can be used to do batched pagination
items, err := client.ListByAlertRuleComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
