
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-07-01-preview/automationrules` Documentation

The `automationrules` SDK allows for interaction with the Azure Resource Manager Service `securityinsights` (API Version `2022-07-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-07-01-preview/automationrules"
```


### Client Initialization

```go
client := automationrules.NewAutomationRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AutomationRulesClient.AutomationRulesCreateOrUpdate`

```go
ctx := context.TODO()
id := automationrules.NewAutomationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "automationRuleIdValue")

payload := automationrules.AutomationRule{
	// ...
}


read, err := client.AutomationRulesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AutomationRulesClient.AutomationRulesDelete`

```go
ctx := context.TODO()
id := automationrules.NewAutomationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "automationRuleIdValue")

read, err := client.AutomationRulesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AutomationRulesClient.AutomationRulesGet`

```go
ctx := context.TODO()
id := automationrules.NewAutomationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "automationRuleIdValue")

read, err := client.AutomationRulesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AutomationRulesClient.AutomationRulesList`

```go
ctx := context.TODO()
id := automationrules.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.AutomationRulesList(ctx, id)` can be used to do batched pagination
items, err := client.AutomationRulesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
