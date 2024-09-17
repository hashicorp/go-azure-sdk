
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-06-01/backupjobs` Documentation

The `backupjobs` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-06-01/backupjobs"
```


### Client Initialization

```go
client := backupjobs.NewBackupJobsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupJobsClient.List`

```go
ctx := context.TODO()
id := backupjobs.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.List(ctx, id, backupjobs.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, backupjobs.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
