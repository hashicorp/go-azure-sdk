
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolstables` Documentation

The `sqlpoolstables` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolstables"
```


### Client Initialization

```go
client := sqlpoolstables.NewSqlPoolsTablesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsTablesClient.SqlPoolTableColumnsListByTableName`

```go
ctx := context.TODO()
id := sqlpoolstables.NewTableID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "schemaValue", "tableValue")

// alternatively `client.SqlPoolTableColumnsListByTableName(ctx, id, sqlpoolstables.DefaultSqlPoolTableColumnsListByTableNameOperationOptions())` can be used to do batched pagination
items, err := client.SqlPoolTableColumnsListByTableNameComplete(ctx, id, sqlpoolstables.DefaultSqlPoolTableColumnsListByTableNameOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SqlPoolsTablesClient.SqlPoolTablesListBySchema`

```go
ctx := context.TODO()
id := sqlpoolstables.NewSchemaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "schemaValue")

// alternatively `client.SqlPoolTablesListBySchema(ctx, id, sqlpoolstables.DefaultSqlPoolTablesListBySchemaOperationOptions())` can be used to do batched pagination
items, err := client.SqlPoolTablesListBySchemaComplete(ctx, id, sqlpoolstables.DefaultSqlPoolTablesListBySchemaOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
