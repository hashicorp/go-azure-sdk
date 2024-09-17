
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-05-01/azurebackupjob` Documentation

The `azurebackupjob` SDK allows for interaction with Azure Resource Manager `dataprotection` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-05-01/azurebackupjob"
```


### Client Initialization

```go
client := azurebackupjob.NewAzureBackupJobClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AzureBackupJobClient.ExportJobsOperationResultGet`

```go
ctx := context.TODO()
id := azurebackupjob.NewOperationIdID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultValue", "operationIdValue")

read, err := client.ExportJobsOperationResultGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AzureBackupJobClient.ExportJobsTrigger`

```go
ctx := context.TODO()
id := azurebackupjob.NewBackupVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultValue")

if err := client.ExportJobsTriggerThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AzureBackupJobClient.JobsGet`

```go
ctx := context.TODO()
id := azurebackupjob.NewBackupJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultValue", "jobIdValue")

read, err := client.JobsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
