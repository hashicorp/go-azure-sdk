
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/manageddatabasesensitivitylabels` Documentation

The `manageddatabasesensitivitylabels` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/manageddatabasesensitivitylabels"
```


### Client Initialization

```go
client := manageddatabasesensitivitylabels.NewManagedDatabaseSensitivityLabelsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewTableColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName", "schemaName", "tableName", "columnName")

payload := manageddatabasesensitivitylabels.SensitivityLabel{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.Delete`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewTableColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName", "schemaName", "tableName", "columnName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.DisableRecommendation`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewTableColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName", "schemaName", "tableName", "columnName")

read, err := client.DisableRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.EnableRecommendation`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewTableColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName", "schemaName", "tableName", "columnName")

read, err := client.EnableRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.Get`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewSensitivityLabelSensitivityLabelSourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName", "schemaName", "tableName", "columnName", "current")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.ListByDatabase`

```go
ctx := context.TODO()
id := commonids.NewSqlManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName")

// alternatively `client.ListByDatabase(ctx, id, manageddatabasesensitivitylabels.DefaultListByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id, manageddatabasesensitivitylabels.DefaultListByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.ListCurrentByDatabase`

```go
ctx := context.TODO()
id := commonids.NewSqlManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName")

// alternatively `client.ListCurrentByDatabase(ctx, id, manageddatabasesensitivitylabels.DefaultListCurrentByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListCurrentByDatabaseComplete(ctx, id, manageddatabasesensitivitylabels.DefaultListCurrentByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.ListRecommendedByDatabase`

```go
ctx := context.TODO()
id := commonids.NewSqlManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName")

// alternatively `client.ListRecommendedByDatabase(ctx, id, manageddatabasesensitivitylabels.DefaultListRecommendedByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListRecommendedByDatabaseComplete(ctx, id, manageddatabasesensitivitylabels.DefaultListRecommendedByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.ManagedDatabaseRecommendedSensitivityLabelsUpdate`

```go
ctx := context.TODO()
id := commonids.NewSqlManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName")

payload := manageddatabasesensitivitylabels.RecommendedSensitivityLabelUpdateList{
	// ...
}


read, err := client.ManagedDatabaseRecommendedSensitivityLabelsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.Update`

```go
ctx := context.TODO()
id := commonids.NewSqlManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName")

payload := manageddatabasesensitivitylabels.SensitivityLabelUpdateList{
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
