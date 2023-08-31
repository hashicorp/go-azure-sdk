
## `github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2022-07-01-preview/accounts` Documentation

The `accounts` SDK allows for interaction with the Azure Resource Manager Service `newrelic` (API Version `2022-07-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2022-07-01-preview/accounts"
```


### Client Initialization

```go
client := accounts.NewAccountsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccountsClient.List`

```go
ctx := context.TODO()
id := accounts.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, accounts.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, accounts.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
