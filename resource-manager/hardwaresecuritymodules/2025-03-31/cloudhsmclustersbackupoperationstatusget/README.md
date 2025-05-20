
## `github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclustersbackupoperationstatusget` Documentation

The `cloudhsmclustersbackupoperationstatusget` SDK allows for interaction with Azure Resource Manager `hardwaresecuritymodules` (API Version `2025-03-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclustersbackupoperationstatusget"
```


### Client Initialization

```go
client := cloudhsmclustersbackupoperationstatusget.NewCloudHSMClustersBackupOperationStatusGetClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudHSMClustersBackupOperationStatusGetClient.CloudHsmClusterBackupStatusGet`

```go
ctx := context.TODO()
id := cloudhsmclustersbackupoperationstatusget.NewBackupOperationStatusID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudHsmClusterName", "jobId")

read, err := client.CloudHsmClusterBackupStatusGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
