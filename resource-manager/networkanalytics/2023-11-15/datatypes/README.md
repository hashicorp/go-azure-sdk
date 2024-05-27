
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkanalytics/2023-11-15/datatypes` Documentation

The `datatypes` SDK allows for interaction with the Azure Resource Manager Service `networkanalytics` (API Version `2023-11-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkanalytics/2023-11-15/datatypes"
```


### Client Initialization

```go
client := datatypes.NewDataTypesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataTypesClient.Create`

```go
ctx := context.TODO()
id := datatypes.NewDataTypeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue", "dataTypeValue")

payload := datatypes.DataType{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataTypesClient.Delete`

```go
ctx := context.TODO()
id := datatypes.NewDataTypeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue", "dataTypeValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DataTypesClient.DeleteData`

```go
ctx := context.TODO()
id := datatypes.NewDataTypeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue", "dataTypeValue")
var payload interface{}

if err := client.DeleteDataThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataTypesClient.GenerateStorageContainerSasToken`

```go
ctx := context.TODO()
id := datatypes.NewDataTypeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue", "dataTypeValue")

payload := datatypes.ContainerSaS{
	// ...
}


read, err := client.GenerateStorageContainerSasToken(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataTypesClient.Get`

```go
ctx := context.TODO()
id := datatypes.NewDataTypeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue", "dataTypeValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataTypesClient.ListByDataProduct`

```go
ctx := context.TODO()
id := datatypes.NewDataProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue")

// alternatively `client.ListByDataProduct(ctx, id)` can be used to do batched pagination
items, err := client.ListByDataProductComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataTypesClient.Update`

```go
ctx := context.TODO()
id := datatypes.NewDataTypeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataProductValue", "dataTypeValue")

payload := datatypes.DataTypeUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
