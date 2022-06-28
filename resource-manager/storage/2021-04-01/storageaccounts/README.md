
## `github.com/hashicorp/go-azure-sdk/resource-manager/storage/2021-04-01/storageaccounts` Documentation

The `storageaccounts` SDK allows for interaction with the Azure Resource Manager Service `storage` (API Version `2021-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storage/2021-04-01/storageaccounts"
```


### Client Initialization

```go
client := storageaccounts.NewStorageAccountsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `StorageAccountsClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := storageaccounts.NewSubscriptionID()

payload := storageaccounts.StorageAccountCheckNameAvailabilityParameters{
	// ...
}

read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.Create`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := storageaccounts.StorageAccountCreateParameters{
	// ...
}

future, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `StorageAccountsClient.Delete`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")
read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.Failover`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")
future, err := client.Failover(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `StorageAccountsClient.GetProperties`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")
read, err := client.GetProperties(ctx, id, storageaccounts.DefaultGetPropertiesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.List`

```go
ctx := context.TODO()
id := storageaccounts.NewSubscriptionID()
// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StorageAccountsClient.ListAccountSAS`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := storageaccounts.AccountSasParameters{
	// ...
}

read, err := client.ListAccountSAS(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := storageaccounts.NewResourceGroupID()
// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StorageAccountsClient.ListKeys`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")
read, err := client.ListKeys(ctx, id, storageaccounts.DefaultListKeysOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.ListServiceSAS`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := storageaccounts.ServiceSasParameters{
	// ...
}

read, err := client.ListServiceSAS(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.RegenerateKey`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := storageaccounts.StorageAccountRegenerateKeyParameters{
	// ...
}

read, err := client.RegenerateKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.RestoreBlobRanges`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := storageaccounts.BlobRestoreParameters{
	// ...
}

future, err := client.RestoreBlobRanges(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `StorageAccountsClient.RevokeUserDelegationKeys`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")
read, err := client.RevokeUserDelegationKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.Update`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := storageaccounts.StorageAccountUpdateParameters{
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
