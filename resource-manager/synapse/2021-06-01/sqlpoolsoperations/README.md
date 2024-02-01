
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsoperations` Documentation

The `sqlpoolsoperations` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsoperations"
```


### Client Initialization

```go
client := sqlpoolsoperations.NewSqlPoolsOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsOperationsClient.SqlPoolOperationsList`

```go
ctx := context.TODO()
id := sqlpoolsoperations.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

// alternatively `client.SqlPoolOperationsList(ctx, id)` can be used to do batched pagination
items, err := client.SqlPoolOperationsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
