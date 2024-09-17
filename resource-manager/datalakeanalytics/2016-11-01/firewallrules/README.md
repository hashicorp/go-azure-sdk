
## `github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/firewallrules` Documentation

The `firewallrules` SDK allows for interaction with Azure Resource Manager `datalakeanalytics` (API Version `2016-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/firewallrules"
```


### Client Initialization

```go
client := firewallrules.NewFirewallRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FirewallRulesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := firewallrules.NewFirewallRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "firewallRuleValue")

payload := firewallrules.CreateOrUpdateFirewallRuleParameters{
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


### Example Usage: `FirewallRulesClient.Delete`

```go
ctx := context.TODO()
id := firewallrules.NewFirewallRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "firewallRuleValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FirewallRulesClient.Get`

```go
ctx := context.TODO()
id := firewallrules.NewFirewallRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "firewallRuleValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FirewallRulesClient.ListByAccount`

```go
ctx := context.TODO()
id := firewallrules.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

// alternatively `client.ListByAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FirewallRulesClient.Update`

```go
ctx := context.TODO()
id := firewallrules.NewFirewallRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "firewallRuleValue")

payload := firewallrules.UpdateFirewallRuleParameters{
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
