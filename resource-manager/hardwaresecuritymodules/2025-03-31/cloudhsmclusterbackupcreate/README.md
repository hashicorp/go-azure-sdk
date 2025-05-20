
## `github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterbackupcreate` Documentation

The `cloudhsmclusterbackupcreate` SDK allows for interaction with Azure Resource Manager `hardwaresecuritymodules` (API Version `2025-03-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterbackupcreate"
```


### Client Initialization

```go
client := cloudhsmclusterbackupcreate.NewCloudHSMClusterBackupCreateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudHSMClusterBackupCreateClient.CloudHsmClustersBackup`

```go
ctx := context.TODO()
id := cloudhsmclusterbackupcreate.NewCloudHsmClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudHsmClusterName")

payload := cloudhsmclusterbackupcreate.BackupRestoreRequestBaseProperties{
	// ...
}


if err := client.CloudHsmClustersBackupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
