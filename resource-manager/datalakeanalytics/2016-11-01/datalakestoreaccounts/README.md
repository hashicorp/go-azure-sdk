
## `github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/datalakestoreaccounts` Documentation

The `datalakestoreaccounts` SDK allows for interaction with Azure Resource Manager `datalakeanalytics` (API Version `2016-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/datalakestoreaccounts"
```


### Client Initialization

```go
client := datalakestoreaccounts.NewDataLakeStoreAccountsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataLakeStoreAccountsClient.Add`

```go
ctx := context.TODO()
id := datalakestoreaccounts.NewDataLakeStoreAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "dataLakeStoreAccountValue")

payload := datalakestoreaccounts.AddDataLakeStoreParameters{
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


### Example Usage: `DataLakeStoreAccountsClient.Delete`

```go
ctx := context.TODO()
id := datalakestoreaccounts.NewDataLakeStoreAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "dataLakeStoreAccountValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataLakeStoreAccountsClient.Get`

```go
ctx := context.TODO()
id := datalakestoreaccounts.NewDataLakeStoreAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "dataLakeStoreAccountValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataLakeStoreAccountsClient.ListByAccount`

```go
ctx := context.TODO()
id := datalakestoreaccounts.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

// alternatively `client.ListByAccount(ctx, id, datalakestoreaccounts.DefaultListByAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByAccountComplete(ctx, id, datalakestoreaccounts.DefaultListByAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
