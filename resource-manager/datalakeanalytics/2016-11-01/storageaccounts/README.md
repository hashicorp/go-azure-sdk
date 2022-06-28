
## `github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/storageaccounts` Documentation

The `storageaccounts` SDK allows for interaction with the Azure Resource Manager Service `datalakeanalytics` (API Version `2016-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/storageaccounts"
```


### Client Initialization

```go
client := storageaccounts.NewStorageAccountsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StorageAccountsClient.Add`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "storageAccountValue")

payload := storageaccounts.AddStorageAccountParameters{
	// ...
}


read, err := client.Add(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.Delete`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "storageAccountValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.Get`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "storageAccountValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.GetStorageContainer`

```go
ctx := context.TODO()
id := storageaccounts.NewContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "storageAccountValue", "containerValue")

read, err := client.GetStorageContainer(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.ListByAccount`

```go
ctx := context.TODO()
id := storageaccounts.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

// alternatively `client.ListByAccount(ctx, id, storageaccounts.DefaultListByAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByAccountComplete(ctx, id, storageaccounts.DefaultListByAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StorageAccountsClient.ListSasTokens`

```go
ctx := context.TODO()
id := storageaccounts.NewContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "storageAccountValue", "containerValue")

// alternatively `client.ListSasTokens(ctx, id)` can be used to do batched pagination
items, err := client.ListSasTokensComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StorageAccountsClient.ListStorageContainers`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "storageAccountValue")

// alternatively `client.ListStorageContainers(ctx, id)` can be used to do batched pagination
items, err := client.ListStorageContainersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StorageAccountsClient.Update`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "storageAccountValue")

payload := storageaccounts.UpdateStorageAccountParameters{
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
