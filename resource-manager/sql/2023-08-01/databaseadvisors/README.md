
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/databaseadvisors` Documentation

The `databaseadvisors` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/databaseadvisors"
```


### Client Initialization

```go
client := databaseadvisors.NewDatabaseAdvisorsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseAdvisorsClient.Get`

```go
ctx := context.TODO()
id := databaseadvisors.NewDatabaseAdvisorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "advisorName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseAdvisorsClient.ListByDatabase`

```go
ctx := context.TODO()
id := commonids.NewSqlDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName")

read, err := client.ListByDatabase(ctx, id, databaseadvisors.DefaultListByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseAdvisorsClient.Update`

```go
ctx := context.TODO()
id := databaseadvisors.NewDatabaseAdvisorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "advisorName")

payload := databaseadvisors.Advisor{
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
