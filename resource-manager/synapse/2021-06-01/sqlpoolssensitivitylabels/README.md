
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolssensitivitylabels` Documentation

The `sqlpoolssensitivitylabels` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolssensitivitylabels"
```


### Client Initialization

```go
client := sqlpoolssensitivitylabels.NewSqlPoolsSensitivityLabelsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsSensitivityLabelsClient.SqlPoolRecommendedSensitivityLabelsUpdate`

```go
ctx := context.TODO()
id := sqlpoolssensitivitylabels.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

payload := sqlpoolssensitivitylabels.RecommendedSensitivityLabelUpdateList{
	// ...
}


read, err := client.SqlPoolRecommendedSensitivityLabelsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsSensitivityLabelsClient.SqlPoolSensitivityLabelsCreateOrUpdate`

```go
ctx := context.TODO()
id := sqlpoolssensitivitylabels.NewColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "schemaValue", "tableValue", "columnValue")

payload := sqlpoolssensitivitylabels.SensitivityLabel{
	// ...
}


read, err := client.SqlPoolSensitivityLabelsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsSensitivityLabelsClient.SqlPoolSensitivityLabelsDelete`

```go
ctx := context.TODO()
id := sqlpoolssensitivitylabels.NewColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "schemaValue", "tableValue", "columnValue")

read, err := client.SqlPoolSensitivityLabelsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsSensitivityLabelsClient.SqlPoolSensitivityLabelsGet`

```go
ctx := context.TODO()
id := sqlpoolssensitivitylabels.NewSensitivityLabelSourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "schemaValue", "tableValue", "columnValue", "current")

read, err := client.SqlPoolSensitivityLabelsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsSensitivityLabelsClient.SqlPoolSensitivityLabelsListCurrent`

```go
ctx := context.TODO()
id := sqlpoolssensitivitylabels.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

// alternatively `client.SqlPoolSensitivityLabelsListCurrent(ctx, id, sqlpoolssensitivitylabels.DefaultSqlPoolSensitivityLabelsListCurrentOperationOptions())` can be used to do batched pagination
items, err := client.SqlPoolSensitivityLabelsListCurrentComplete(ctx, id, sqlpoolssensitivitylabels.DefaultSqlPoolSensitivityLabelsListCurrentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SqlPoolsSensitivityLabelsClient.SqlPoolSensitivityLabelsListRecommended`

```go
ctx := context.TODO()
id := sqlpoolssensitivitylabels.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

// alternatively `client.SqlPoolSensitivityLabelsListRecommended(ctx, id, sqlpoolssensitivitylabels.DefaultSqlPoolSensitivityLabelsListRecommendedOperationOptions())` can be used to do batched pagination
items, err := client.SqlPoolSensitivityLabelsListRecommendedComplete(ctx, id, sqlpoolssensitivitylabels.DefaultSqlPoolSensitivityLabelsListRecommendedOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SqlPoolsSensitivityLabelsClient.SqlPoolSensitivityLabelsUpdate`

```go
ctx := context.TODO()
id := sqlpoolssensitivitylabels.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

payload := sqlpoolssensitivitylabels.SensitivityLabelUpdateList{
	// ...
}


read, err := client.SqlPoolSensitivityLabelsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
