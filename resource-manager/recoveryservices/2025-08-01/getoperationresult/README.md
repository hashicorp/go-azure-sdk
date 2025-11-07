
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-08-01/getoperationresult` Documentation

The `getoperationresult` SDK allows for interaction with Azure Resource Manager `recoveryservices` (API Version `2025-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-08-01/getoperationresult"
```


### Client Initialization

```go
client := getoperationresult.NewGetOperationResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GetOperationResultClient.VaultsListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.VaultsListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.VaultsListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
