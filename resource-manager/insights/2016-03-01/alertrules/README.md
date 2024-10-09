
## `github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/alertrules` Documentation

The `alertrules` SDK allows for interaction with Azure Resource Manager `insights` (API Version `2016-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/alertrules"
```


### Client Initialization

```go
client := alertrules.NewAlertRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AlertRulesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := alertrules.NewAlertRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "alertRuleName")

payload := alertrules.AlertRuleResource{
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


### Example Usage: `AlertRulesClient.Delete`

```go
ctx := context.TODO()
id := alertrules.NewAlertRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "alertRuleName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertRulesClient.Get`

```go
ctx := context.TODO()
id := alertrules.NewAlertRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "alertRuleName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertRulesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.ListByResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertRulesClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ListBySubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
