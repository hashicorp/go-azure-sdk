
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securescore` Documentation

The `securescore` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securescore"
```


### Client Initialization

```go
client := securescore.NewSecureScoreClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SecureScoreClient.ControlsListBySecureScore`

```go
ctx := context.TODO()
id := securescore.NewSecureScoreID("12345678-1234-9876-4563-123456789012", "secureScoreValue")

// alternatively `client.ControlsListBySecureScore(ctx, id, securescore.DefaultControlsListBySecureScoreOperationOptions())` can be used to do batched pagination
items, err := client.ControlsListBySecureScoreComplete(ctx, id, securescore.DefaultControlsListBySecureScoreOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SecureScoreClient.Get`

```go
ctx := context.TODO()
id := securescore.NewSecureScoreID("12345678-1234-9876-4563-123456789012", "secureScoreValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecureScoreClient.List`

```go
ctx := context.TODO()
id := securescore.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
