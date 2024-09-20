
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/authorizationrulesdisasterrecoveryconfigs` Documentation

The `authorizationrulesdisasterrecoveryconfigs` SDK allows for interaction with Azure Resource Manager `eventhub` (API Version `2022-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/authorizationrulesdisasterrecoveryconfigs"
```


### Client Initialization

```go
client := authorizationrulesdisasterrecoveryconfigs.NewAuthorizationRulesDisasterRecoveryConfigsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AuthorizationRulesDisasterRecoveryConfigsClient.DisasterRecoveryConfigsGetAuthorizationRule`

```go
ctx := context.TODO()
id := authorizationrulesdisasterrecoveryconfigs.NewDisasterRecoveryConfigAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "alias", "authorizationRuleName")

read, err := client.DisasterRecoveryConfigsGetAuthorizationRule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthorizationRulesDisasterRecoveryConfigsClient.DisasterRecoveryConfigsListAuthorizationRules`

```go
ctx := context.TODO()
id := authorizationrulesdisasterrecoveryconfigs.NewDisasterRecoveryConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "alias")

// alternatively `client.DisasterRecoveryConfigsListAuthorizationRules(ctx, id)` can be used to do batched pagination
items, err := client.DisasterRecoveryConfigsListAuthorizationRulesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AuthorizationRulesDisasterRecoveryConfigsClient.DisasterRecoveryConfigsListKeys`

```go
ctx := context.TODO()
id := authorizationrulesdisasterrecoveryconfigs.NewDisasterRecoveryConfigAuthorizationRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "alias", "authorizationRuleName")

read, err := client.DisasterRecoveryConfigsListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
