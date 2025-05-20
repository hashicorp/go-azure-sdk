
## `github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclustersrestoreoperationcreate` Documentation

The `cloudhsmclustersrestoreoperationcreate` SDK allows for interaction with Azure Resource Manager `hardwaresecuritymodules` (API Version `2025-03-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclustersrestoreoperationcreate"
```


### Client Initialization

```go
client := cloudhsmclustersrestoreoperationcreate.NewCloudHSMClustersRestoreOperationCreateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudHSMClustersRestoreOperationCreateClient.CloudHsmClustersRestore`

```go
ctx := context.TODO()
id := cloudhsmclustersrestoreoperationcreate.NewCloudHsmClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudHsmClusterName")

payload := cloudhsmclustersrestoreoperationcreate.RestoreRequestProperties{
	// ...
}


if err := client.CloudHsmClustersRestoreThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
