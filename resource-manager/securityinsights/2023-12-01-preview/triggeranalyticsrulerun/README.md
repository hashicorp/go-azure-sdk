
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/triggeranalyticsrulerun` Documentation

The `triggeranalyticsrulerun` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/triggeranalyticsrulerun"
```


### Client Initialization

```go
client := triggeranalyticsrulerun.NewTriggerAnalyticsRuleRunClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TriggerAnalyticsRuleRunClient.AlertRuleTriggerRuleRun`

```go
ctx := context.TODO()
id := triggeranalyticsrulerun.NewAlertRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "ruleId")

payload := triggeranalyticsrulerun.AnalyticsRuleRunTrigger{
	// ...
}


if err := client.AlertRuleTriggerRuleRunThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
