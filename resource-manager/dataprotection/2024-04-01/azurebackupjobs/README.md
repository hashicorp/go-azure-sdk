
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/azurebackupjobs` Documentation

The `azurebackupjobs` SDK allows for interaction with the Azure Resource Manager Service `dataprotection` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/azurebackupjobs"
```


### Client Initialization

```go
client := azurebackupjobs.NewAzureBackupJobsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AzureBackupJobsClient.JobsList`

```go
ctx := context.TODO()
id := azurebackupjobs.NewBackupVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultValue")

// alternatively `client.JobsList(ctx, id)` can be used to do batched pagination
items, err := client.JobsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
