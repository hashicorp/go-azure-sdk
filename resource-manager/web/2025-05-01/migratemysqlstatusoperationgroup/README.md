
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/migratemysqlstatusoperationgroup` Documentation

The `migratemysqlstatusoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/migratemysqlstatusoperationgroup"
```


### Client Initialization

```go
client := migratemysqlstatusoperationgroup.NewMigrateMySqlStatusOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MigrateMySqlStatusOperationGroupClient.WebAppsGetMigrateMySqlStatusSlot`

```go
ctx := context.TODO()
id := migratemysqlstatusoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.WebAppsGetMigrateMySqlStatusSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
