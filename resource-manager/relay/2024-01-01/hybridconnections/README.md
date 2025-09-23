
## `github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/hybridconnections` Documentation

The `hybridconnections` SDK allows for interaction with Azure Resource Manager `relay` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/hybridconnections"
```


### Client Initialization

```go
client := hybridconnections.NewHybridConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HybridConnectionsClient.CreateOrUpdateAuthorizationRule`

```go
ctx := context.TODO()
id := hybridconnections.NewHybridConnectionAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "hybridConnectionName", "authorizationRuleName")

payload := hybridconnections.AuthorizationRule{
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


### Example Usage: `HybridConnectionsClient.DeleteAuthorizationRule`

```go
ctx := context.TODO()
id := hybridconnections.NewHybridConnectionAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "hybridConnectionName", "authorizationRuleName")

read, err := client.DeleteAuthorizationRule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionsClient.GetAuthorizationRule`

```go
ctx := context.TODO()
id := hybridconnections.NewHybridConnectionAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "hybridConnectionName", "authorizationRuleName")

read, err := client.GetAuthorizationRule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionsClient.ListAuthorizationRules`

```go
ctx := context.TODO()
id := hybridconnections.NewHybridConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "hybridConnectionName")

// alternatively `client.ListAuthorizationRules(ctx, id)` can be used to do batched pagination
items, err := client.ListAuthorizationRulesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `HybridConnectionsClient.ListKeys`

```go
ctx := context.TODO()
id := hybridconnections.NewHybridConnectionAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "hybridConnectionName", "authorizationRuleName")

read, err := client.ListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionsClient.RegenerateKeys`

```go
ctx := context.TODO()
id := hybridconnections.NewHybridConnectionAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "hybridConnectionName", "authorizationRuleName")

payload := hybridconnections.RegenerateAccessKeyParameters{
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
