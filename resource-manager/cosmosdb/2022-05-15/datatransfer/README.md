
## `github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-05-15/datatransfer` Documentation

The `datatransfer` SDK allows for interaction with the Azure Resource Manager Service `cosmosdb` (API Version `2022-05-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-05-15/datatransfer"
```


### Client Initialization

```go
client := datatransfer.NewDataTransferClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataTransferClient.ServiceCreate`

```go
ctx := context.TODO()
id := datatransfer.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountValue", "serviceValue")

payload := datatransfer.ServiceResourceCreateUpdateParameters{
	// ...
}


if err := client.ServiceCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataTransferClient.ServiceDelete`

```go
ctx := context.TODO()
id := datatransfer.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountValue", "serviceValue")

if err := client.ServiceDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DataTransferClient.ServiceGet`

```go
ctx := context.TODO()
id := datatransfer.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountValue", "serviceValue")

read, err := client.ServiceGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
