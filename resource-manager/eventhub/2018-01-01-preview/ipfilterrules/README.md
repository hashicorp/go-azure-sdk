
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/ipfilterrules` Documentation

The `ipfilterrules` SDK allows for interaction with the Azure Resource Manager Service `eventhub` (API Version `2018-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/ipfilterrules"
```


### Client Initialization

```go
client := ipfilterrules.NewIpFilterRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `IpFilterRulesClient.NamespacesCreateOrUpdateIpFilterRule`

```go
ctx := context.TODO()
id := ipfilterrules.NewIpfilterruleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "ipFilterRuleValue")

payload := ipfilterrules.IpFilterRule{
	// ...
}

read, err := client.NamespacesCreateOrUpdateIpFilterRule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IpFilterRulesClient.NamespacesDeleteIpFilterRule`

```go
ctx := context.TODO()
id := ipfilterrules.NewIpfilterruleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "ipFilterRuleValue")
read, err := client.NamespacesDeleteIpFilterRule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IpFilterRulesClient.NamespacesGetIpFilterRule`

```go
ctx := context.TODO()
id := ipfilterrules.NewIpfilterruleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "ipFilterRuleValue")
read, err := client.NamespacesGetIpFilterRule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IpFilterRulesClient.NamespacesListIPFilterRules`

```go
ctx := context.TODO()
id := ipfilterrules.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")
// alternatively `client.NamespacesListIPFilterRules(ctx, id)` can be used to do batched pagination
items, err := client.NamespacesListIPFilterRulesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
