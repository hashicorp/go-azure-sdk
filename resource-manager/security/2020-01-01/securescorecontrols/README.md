
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securescorecontrols` Documentation

The `securescorecontrols` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securescorecontrols"
```


### Client Initialization

```go
client := securescorecontrols.NewSecureScoreControlsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SecureScoreControlsClient.List`

```go
ctx := context.TODO()
id := securescorecontrols.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, securescorecontrols.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, securescorecontrols.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
