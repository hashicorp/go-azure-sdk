
## `github.com/hashicorp/go-azure-sdk/data-plane/batch/2020-03-01.11.0/accounts` Documentation

The `accounts` SDK allows for interaction with <unknown source data type> `batch` (API Version `2020-03-01.11.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/batch/2020-03-01.11.0/accounts"
```


### Client Initialization

```go
client := accounts.NewAccountsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccountsClient.AccountListPoolNodeCounts`

```go
ctx := context.TODO()


// alternatively `client.AccountListPoolNodeCounts(ctx, accounts.DefaultAccountListPoolNodeCountsOperationOptions())` can be used to do batched pagination
items, err := client.AccountListPoolNodeCountsComplete(ctx, accounts.DefaultAccountListPoolNodeCountsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AccountsClient.AccountListSupportedImages`

```go
ctx := context.TODO()


// alternatively `client.AccountListSupportedImages(ctx, accounts.DefaultAccountListSupportedImagesOperationOptions())` can be used to do batched pagination
items, err := client.AccountListSupportedImagesComplete(ctx, accounts.DefaultAccountListSupportedImagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
