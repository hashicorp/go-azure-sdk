
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/sensitivitylabels` Documentation

The `sensitivitylabels` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2022-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/sensitivitylabels"
```


### Client Initialization

```go
client := sensitivitylabels.NewSensitivityLabelsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SensitivityLabelsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := sensitivitylabels.NewColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "schemaValue", "tableValue", "columnValue")

payload := sensitivitylabels.SensitivityLabel{
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


### Example Usage: `SensitivityLabelsClient.Delete`

```go
ctx := context.TODO()
id := sensitivitylabels.NewColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "schemaValue", "tableValue", "columnValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SensitivityLabelsClient.DisableRecommendation`

```go
ctx := context.TODO()
id := sensitivitylabels.NewColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "schemaValue", "tableValue", "columnValue")

read, err := client.DisableRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SensitivityLabelsClient.EnableRecommendation`

```go
ctx := context.TODO()
id := sensitivitylabels.NewColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "schemaValue", "tableValue", "columnValue")

read, err := client.EnableRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SensitivityLabelsClient.Get`

```go
ctx := context.TODO()
id := sensitivitylabels.NewColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "schemaValue", "tableValue", "columnValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SensitivityLabelsClient.ListByDatabase`

```go
ctx := context.TODO()
id := sensitivitylabels.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

// alternatively `client.ListByDatabase(ctx, id, sensitivitylabels.DefaultListByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id, sensitivitylabels.DefaultListByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SensitivityLabelsClient.ListCurrentByDatabase`

```go
ctx := context.TODO()
id := sensitivitylabels.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

// alternatively `client.ListCurrentByDatabase(ctx, id, sensitivitylabels.DefaultListCurrentByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListCurrentByDatabaseComplete(ctx, id, sensitivitylabels.DefaultListCurrentByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SensitivityLabelsClient.ListRecommendedByDatabase`

```go
ctx := context.TODO()
id := sensitivitylabels.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

// alternatively `client.ListRecommendedByDatabase(ctx, id, sensitivitylabels.DefaultListRecommendedByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListRecommendedByDatabaseComplete(ctx, id, sensitivitylabels.DefaultListRecommendedByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SensitivityLabelsClient.RecommendedSensitivityLabelsUpdate`

```go
ctx := context.TODO()
id := sensitivitylabels.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

payload := sensitivitylabels.RecommendedSensitivityLabelUpdateList{
	// ...
}


read, err := client.RecommendedSensitivityLabelsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SensitivityLabelsClient.Update`

```go
ctx := context.TODO()
id := sensitivitylabels.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

payload := sensitivitylabels.SensitivityLabelUpdateList{
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
