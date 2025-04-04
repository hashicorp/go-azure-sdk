
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/ipv6firewallrules` Documentation

The `ipv6firewallrules` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/ipv6firewallrules"
```


### Client Initialization

```go
client := ipv6firewallrules.NewIPv6FirewallRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IPv6FirewallRulesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := ipv6firewallrules.NewIPv6FirewallRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "ipv6FirewallRuleName")

payload := ipv6firewallrules.IPv6FirewallRule{
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


### Example Usage: `IPv6FirewallRulesClient.Delete`

```go
ctx := context.TODO()
id := ipv6firewallrules.NewIPv6FirewallRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "ipv6FirewallRuleName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IPv6FirewallRulesClient.Get`

```go
ctx := context.TODO()
id := ipv6firewallrules.NewIPv6FirewallRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "ipv6FirewallRuleName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IPv6FirewallRulesClient.ListByServer`

```go
ctx := context.TODO()
id := commonids.NewSqlServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
