
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/firewallrules` Documentation

The `firewallrules` SDK allows for interaction with the Azure Resource Manager Service `postgresqlhsc` (API Version `2020-10-05-privatepreview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/firewallrules"
```


### Client Initialization

```go
client := firewallrules.NewFirewallRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `FirewallRulesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := firewallrules.NewFirewallRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupValue", "firewallRuleValue")

payload := firewallrules.FirewallRule{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `FirewallRulesClient.Delete`

```go
ctx := context.TODO()
id := firewallrules.NewFirewallRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupValue", "firewallRuleValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `FirewallRulesClient.Get`

```go
ctx := context.TODO()
id := firewallrules.NewFirewallRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupValue", "firewallRuleValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FirewallRulesClient.ListByServerGroup`

```go
ctx := context.TODO()
id := firewallrules.NewServerGroupsv2ID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupValue")
read, err := client.ListByServerGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
