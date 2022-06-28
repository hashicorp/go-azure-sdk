
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/dimensions` Documentation

The `dimensions` SDK allows for interaction with the Azure Resource Manager Service `costmanagement` (API Version `2021-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/dimensions"
```


### Client Initialization

```go
client := dimensions.NewDimensionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DimensionsClient.ByExternalCloudProviderType`

```go
ctx := context.TODO()
id := dimensions.NewExternalCloudProviderTypeID("externalBillingAccounts", "externalCloudProviderIdValue")

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
id := dimensions.NewScopeID()

read, err := client.List(ctx, id, dimensions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
