
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2025-06-30/databasemigrationssqlvm` Documentation

The `databasemigrationssqlvm` SDK allows for interaction with Azure Resource Manager `datamigration` (API Version `2025-06-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2025-06-30/databasemigrationssqlvm"
```


### Client Initialization

```go
client := databasemigrationssqlvm.NewDatabaseMigrationsSqlVMClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseMigrationsSqlVMClient.Delete`

```go
ctx := context.TODO()
id := databasemigrationssqlvm.NewSqlVirtualMachineProviders2DatabaseMigrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlVirtualMachineName", "databaseMigrationName")

if err := client.DeleteThenPoll(ctx, id, databasemigrationssqlvm.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```
