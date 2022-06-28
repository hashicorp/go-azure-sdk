
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/virtualnetworkrules` Documentation

The `virtualnetworkrules` SDK allows for interaction with the Azure Resource Manager Service `eventhub` (API Version `2018-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/virtualnetworkrules"
```


### Client Initialization

```go
client := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkRulesClient.NamespacesCreateOrUpdateVirtualNetworkRule`

```go
ctx := context.TODO()
id := virtualnetworkrules.NewVirtualnetworkruleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "virtualNetworkRuleValue")

payload := virtualnetworkrules.VirtualNetworkRule{
	// ...
}

read, err := client.NamespacesCreateOrUpdateVirtualNetworkRule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualNetworkRulesClient.NamespacesDeleteVirtualNetworkRule`

```go
ctx := context.TODO()
id := virtualnetworkrules.NewVirtualnetworkruleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "virtualNetworkRuleValue")
read, err := client.NamespacesDeleteVirtualNetworkRule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualNetworkRulesClient.NamespacesGetVirtualNetworkRule`

```go
ctx := context.TODO()
id := virtualnetworkrules.NewVirtualnetworkruleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "virtualNetworkRuleValue")
read, err := client.NamespacesGetVirtualNetworkRule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualNetworkRulesClient.NamespacesListVirtualNetworkRules`

```go
ctx := context.TODO()
id := virtualnetworkrules.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")
// alternatively `client.NamespacesListVirtualNetworkRules(ctx, id)` can be used to do batched pagination
items, err := client.NamespacesListVirtualNetworkRulesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
