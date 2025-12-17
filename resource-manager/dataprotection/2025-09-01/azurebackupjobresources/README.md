
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-09-01/azurebackupjobresources` Documentation

The `azurebackupjobresources` SDK allows for interaction with Azure Resource Manager `dataprotection` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-09-01/azurebackupjobresources"
```


### Client Initialization

```go
client := azurebackupjobresources.NewAzureBackupJobResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AzureBackupJobResourcesClient.JobsGet`

```go
ctx := context.TODO()
id := azurebackupjobresources.NewBackupJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultName", "jobId")

read, err := client.JobsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AzureBackupJobResourcesClient.JobsList`

```go
ctx := context.TODO()
id := azurebackupjobresources.NewBackupVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultName")

// alternatively `client.JobsList(ctx, id)` can be used to do batched pagination
items, err := client.JobsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
