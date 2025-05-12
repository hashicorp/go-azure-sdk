
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/dimensions` Documentation

The `dimensions` SDK allows for interaction with Azure Resource Manager `costmanagement` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/dimensions"
```


### Client Initialization

```go
client := dimensions.NewDimensionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DimensionsClient.ByExternalCloudProviderType`

```go
ctx := context.TODO()
id := dimensions.NewExternalCloudProviderTypeID("externalBillingAccounts", "externalCloudProviderId")

read, err := client.ByExternalCloudProviderType(ctx, id, dimensions.DefaultByExternalCloudProviderTypeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DimensionsClient.List`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.List(ctx, id, dimensions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
