
## `github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2024-11-01/mhsmlistregions` Documentation

The `mhsmlistregions` SDK allows for interaction with Azure Resource Manager `keyvault` (API Version `2024-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2024-11-01/mhsmlistregions"
```


### Client Initialization

```go
client := mhsmlistregions.NewMHSMListRegionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MHSMListRegionsClient.MHSMRegionsListByResource`

```go
ctx := context.TODO()
id := mhsmlistregions.NewManagedHSMID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMName")

// alternatively `client.MHSMRegionsListByResource(ctx, id)` can be used to do batched pagination
items, err := client.MHSMRegionsListByResourceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
