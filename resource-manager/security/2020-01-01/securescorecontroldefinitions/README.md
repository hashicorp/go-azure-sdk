
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securescorecontroldefinitions` Documentation

The `securescorecontroldefinitions` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securescorecontroldefinitions"
```


### Client Initialization

```go
client := securescorecontroldefinitions.NewSecureScoreControlDefinitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SecureScoreControlDefinitionsClient.SecureScoreControlDefinitionsList`

```go
ctx := context.TODO()


// alternatively `client.SecureScoreControlDefinitionsList(ctx)` can be used to do batched pagination
items, err := client.SecureScoreControlDefinitionsListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SecureScoreControlDefinitionsClient.SecureScoreControlDefinitionsListBySubscription`

```go
ctx := context.TODO()
id := securescorecontroldefinitions.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.SecureScoreControlDefinitionsListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.SecureScoreControlDefinitionsListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
