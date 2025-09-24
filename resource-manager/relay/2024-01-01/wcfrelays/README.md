
## `github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/wcfrelays` Documentation

The `wcfrelays` SDK allows for interaction with Azure Resource Manager `relay` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/wcfrelays"
```


### Client Initialization

```go
client := wcfrelays.NewWCFRelaysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WCFRelaysClient.CreateOrUpdateAuthorizationRule`

```go
ctx := context.TODO()
id := wcfrelays.NewWcfRelayAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "wcfRelayName", "authorizationRuleName")

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


### Example Usage: `WCFRelaysClient.DeleteAuthorizationRule`

```go
ctx := context.TODO()
id := wcfrelays.NewWcfRelayAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "wcfRelayName", "authorizationRuleName")

read, err := client.DeleteAuthorizationRule(ctx, id)
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
id := wcfrelays.NewWcfRelayAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "wcfRelayName", "authorizationRuleName")

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
id := wcfrelays.NewWcfRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "wcfRelayName")

// alternatively `client.ListAuthorizationRules(ctx, id)` can be used to do batched pagination
items, err := client.ListAuthorizationRulesComplete(ctx, id)
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
id := wcfrelays.NewWcfRelayAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "wcfRelayName", "authorizationRuleName")

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
id := wcfrelays.NewWcfRelayAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "wcfRelayName", "authorizationRuleName")

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
