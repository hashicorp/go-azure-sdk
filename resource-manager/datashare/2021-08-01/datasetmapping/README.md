
## `github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/datasetmapping` Documentation

The `datasetmapping` SDK allows for interaction with the Azure Resource Manager Service `datashare` (API Version `2021-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/datasetmapping"
```


### Client Initialization

```go
client := datasetmapping.NewDataSetMappingClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataSetMappingClient.Create`

```go
ctx := context.TODO()
id := datasetmapping.NewDataSetMappingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareSubscriptionValue", "dataSetMappingValue")

payload := datasetmapping.DataSetMapping{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSetMappingClient.Delete`

```go
ctx := context.TODO()
id := datasetmapping.NewDataSetMappingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareSubscriptionValue", "dataSetMappingValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSetMappingClient.Get`

```go
ctx := context.TODO()
id := datasetmapping.NewDataSetMappingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareSubscriptionValue", "dataSetMappingValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSetMappingClient.ListByShareSubscription`

```go
ctx := context.TODO()
id := datasetmapping.NewShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareSubscriptionValue")

// alternatively `client.ListByShareSubscription(ctx, id, datasetmapping.DefaultListByShareSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListByShareSubscriptionComplete(ctx, id, datasetmapping.DefaultListByShareSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
