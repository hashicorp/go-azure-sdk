
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/backupcrrjobs` Documentation

The `backupcrrjobs` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2023-01-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/backupcrrjobs"
```


### Client Initialization

```go
client := backupcrrjobs.NewBackupCrrJobsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupCrrJobsClient.List`

```go
ctx := context.TODO()
id := backupcrrjobs.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := backupcrrjobs.CrrJobRequest{
	// ...
}


// alternatively `client.List(ctx, id, payload, backupcrrjobs.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, payload, backupcrrjobs.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
