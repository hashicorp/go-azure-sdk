
## `github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/returnsrestoreoperationstatus` Documentation

The `returnsrestoreoperationstatus` SDK allows for interaction with Azure Resource Manager `hardwaresecuritymodules` (API Version `2025-03-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/returnsrestoreoperationstatus"
```


### Client Initialization

```go
client := returnsrestoreoperationstatus.NewReturnsRestoreOperationStatusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReturnsRestoreOperationStatusClient.CloudHsmClusterRestoreStatusGet`

```go
ctx := context.TODO()
id := returnsrestoreoperationstatus.NewRestoreOperationStatusID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudHsmClusterName", "jobId")

read, err := client.CloudHsmClusterRestoreStatusGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
