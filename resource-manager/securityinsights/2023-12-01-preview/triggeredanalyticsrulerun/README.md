
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/triggeredanalyticsrulerun` Documentation

The `triggeredanalyticsrulerun` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/triggeredanalyticsrulerun"
```


### Client Initialization

```go
client := triggeredanalyticsrulerun.NewTriggeredAnalyticsRuleRunClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TriggeredAnalyticsRuleRunClient.Get`

```go
ctx := context.TODO()
id := triggeredanalyticsrulerun.NewTriggeredAnalyticsRuleRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "ruleRunId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
