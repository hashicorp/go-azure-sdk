
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/alertssuppressionrules` Documentation

The `alertssuppressionrules` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2019-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/alertssuppressionrules"
```


### Client Initialization

```go
client := alertssuppressionrules.NewAlertsSuppressionRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AlertsSuppressionRulesClient.Delete`

```go
ctx := context.TODO()
id := alertssuppressionrules.NewAlertsSuppressionRuleID("12345678-1234-9876-4563-123456789012", "alertsSuppressionRuleValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsSuppressionRulesClient.Get`

```go
ctx := context.TODO()
id := alertssuppressionrules.NewAlertsSuppressionRuleID("12345678-1234-9876-4563-123456789012", "alertsSuppressionRuleValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsSuppressionRulesClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, alertssuppressionrules.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, alertssuppressionrules.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AlertsSuppressionRulesClient.Update`

```go
ctx := context.TODO()
id := alertssuppressionrules.NewAlertsSuppressionRuleID("12345678-1234-9876-4563-123456789012", "alertsSuppressionRuleValue")

payload := alertssuppressionrules.AlertsSuppressionRule{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
