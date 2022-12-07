
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2022-10-01/backupusagesummaries` Documentation

The `backupusagesummaries` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2022-10-01/backupusagesummaries"
```


### Client Initialization

```go
client := backupusagesummaries.NewBackupUsageSummariesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupUsageSummariesClient.List`

```go
ctx := context.TODO()
id := backupusagesummaries.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

read, err := client.List(ctx, id, backupusagesummaries.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
