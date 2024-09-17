
## `github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/virtualnetworkrules` Documentation

The `virtualnetworkrules` SDK allows for interaction with Azure Resource Manager `datalakestore` (API Version `2016-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/virtualnetworkrules"
```


### Client Initialization

```go
client := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualNetworkRulesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := virtualnetworkrules.NewVirtualNetworkRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "virtualNetworkRuleValue")

payload := virtualnetworkrules.CreateOrUpdateVirtualNetworkRuleParameters{
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


### Example Usage: `VirtualNetworkRulesClient.Delete`

```go
ctx := context.TODO()
id := virtualnetworkrules.NewVirtualNetworkRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "virtualNetworkRuleValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualNetworkRulesClient.Get`

```go
ctx := context.TODO()
id := virtualnetworkrules.NewVirtualNetworkRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "virtualNetworkRuleValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualNetworkRulesClient.ListByAccount`

```go
ctx := context.TODO()
id := virtualnetworkrules.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

// alternatively `client.ListByAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualNetworkRulesClient.Update`

```go
ctx := context.TODO()
id := virtualnetworkrules.NewVirtualNetworkRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "virtualNetworkRuleValue")

payload := virtualnetworkrules.UpdateVirtualNetworkRuleParameters{
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
