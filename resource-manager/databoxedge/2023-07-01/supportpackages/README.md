
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/supportpackages` Documentation

The `supportpackages` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2023-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/supportpackages"
```


### Client Initialization

```go
client := supportpackages.NewSupportPackagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SupportPackagesClient.TriggerSupportPackage`

```go
ctx := context.TODO()
id := supportpackages.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue")

payload := supportpackages.TriggerSupportPackageRequest{
	// ...
}


if err := client.TriggerSupportPackageThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
