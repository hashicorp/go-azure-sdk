
## `github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2025-05-01/dnssecurityrules` Documentation

The `dnssecurityrules` SDK allows for interaction with Azure Resource Manager `dnsresolver` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2025-05-01/dnssecurityrules"
```


### Client Initialization

```go
client := dnssecurityrules.NewDnsSecurityRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DnsSecurityRulesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := dnssecurityrules.NewDnsSecurityRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName", "dnsSecurityRuleName")

payload := dnssecurityrules.DnsSecurityRule{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, dnssecurityrules.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DnsSecurityRulesClient.Delete`

```go
ctx := context.TODO()
id := dnssecurityrules.NewDnsSecurityRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName", "dnsSecurityRuleName")

if err := client.DeleteThenPoll(ctx, id, dnssecurityrules.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DnsSecurityRulesClient.Get`

```go
ctx := context.TODO()
id := dnssecurityrules.NewDnsSecurityRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName", "dnsSecurityRuleName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DnsSecurityRulesClient.List`

```go
ctx := context.TODO()
id := dnssecurityrules.NewDnsResolverPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName")

// alternatively `client.List(ctx, id, dnssecurityrules.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, dnssecurityrules.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DnsSecurityRulesClient.Update`

```go
ctx := context.TODO()
id := dnssecurityrules.NewDnsSecurityRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName", "dnsSecurityRuleName")

payload := dnssecurityrules.DnsSecurityRulePatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, dnssecurityrules.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
