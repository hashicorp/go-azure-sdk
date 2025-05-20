
## `github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterupdate` Documentation

The `cloudhsmclusterupdate` SDK allows for interaction with Azure Resource Manager `hardwaresecuritymodules` (API Version `2025-03-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterupdate"
```


### Client Initialization

```go
client := cloudhsmclusterupdate.NewCloudHSMClusterUpdateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudHSMClusterUpdateClient.CloudHsmClustersUpdate`

```go
ctx := context.TODO()
id := cloudhsmclusterupdate.NewCloudHsmClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudHsmClusterName")

payload := cloudhsmclusterupdate.CloudHsmClusterPatchParameters{
	// ...
}


if err := client.CloudHsmClustersUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
