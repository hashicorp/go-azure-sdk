
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-04-01/backupstatus` Documentation

The `backupstatus` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-04-01/backupstatus"
```


### Client Initialization

```go
client := backupstatus.NewBackupStatusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupStatusClient.Get`

```go
ctx := context.TODO()
id := backupstatus.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := backupstatus.BackupStatusRequest{
	// ...
}


read, err := client.Get(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
