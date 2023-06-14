
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/datawarehouseuseractivities` Documentation

The `datawarehouseuseractivities` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2017-03-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/datawarehouseuseractivities"
```


### Client Initialization

```go
client := datawarehouseuseractivities.NewDataWarehouseUserActivitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataWarehouseUserActivitiesClient.Get`

```go
ctx := context.TODO()
id := datawarehouseuseractivities.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
