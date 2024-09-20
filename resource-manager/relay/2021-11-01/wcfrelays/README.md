
## `github.com/hashicorp/go-azure-sdk/resource-manager/relay/2021-11-01/wcfrelays` Documentation

The `wcfrelays` SDK allows for interaction with Azure Resource Manager `relay` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/relay/2021-11-01/wcfrelays"
```


### Client Initialization

```go
client := wcfrelays.NewWCFRelaysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WCFRelaysClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := wcfrelays.NewWcfRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "relayName")

payload := wcfrelays.WcfRelay{
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


### Example Usage: `WCFRelaysClient.CreateOrUpdateAuthorizationRule`

```go
ctx := context.TODO()
id := wcfrelays.NewWcfRelayAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "relayName", "authorizationRuleName")

payload := wcfrelays.AuthorizationRule{
	// ...
}


read, err := client.CreateOrUpdateAuthorizationRule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WCFRelaysClient.Delete`

```go
ctx := context.TODO()
id := wcfrelays.NewWcfRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "relayName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WCFRelaysClient.DeleteAuthorizationRule`

```go
ctx := context.TODO()
id := wcfrelays.NewWcfRelayAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "relayName", "authorizationRuleName")

read, err := client.DeleteAuthorizationRule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WCFRelaysClient.Get`

```go
ctx := context.TODO()
id := wcfrelays.NewWcfRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "relayName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WCFRelaysClient.GetAuthorizationRule`

```go
ctx := context.TODO()
id := wcfrelays.NewWcfRelayAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "relayName", "authorizationRuleName")

read, err := client.GetAuthorizationRule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WCFRelaysClient.ListAuthorizationRules`

```go
ctx := context.TODO()
id := wcfrelays.NewWcfRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "relayName")

// alternatively `client.ListAuthorizationRules(ctx, id)` can be used to do batched pagination
items, err := client.ListAuthorizationRulesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WCFRelaysClient.ListByNamespace`

```go
ctx := context.TODO()
id := wcfrelays.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

// alternatively `client.ListByNamespace(ctx, id)` can be used to do batched pagination
items, err := client.ListByNamespaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WCFRelaysClient.ListKeys`

```go
ctx := context.TODO()
id := wcfrelays.NewWcfRelayAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "relayName", "authorizationRuleName")

read, err := client.ListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WCFRelaysClient.RegenerateKeys`

```go
ctx := context.TODO()
id := wcfrelays.NewWcfRelayAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "relayName", "authorizationRuleName")

payload := wcfrelays.RegenerateAccessKeyParameters{
	// ...
}


read, err := client.RegenerateKeys(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
