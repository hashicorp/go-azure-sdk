
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-08-01/targetcomputesizes` Documentation

The `targetcomputesizes` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-08-01/targetcomputesizes"
```


### Client Initialization

```go
client := targetcomputesizes.NewTargetComputeSizesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TargetComputeSizesClient.ListByReplicationProtectedItems`

```go
ctx := context.TODO()
id := targetcomputesizes.NewReplicationProtectedItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationProtectedItemValue")

// alternatively `client.ListByReplicationProtectedItems(ctx, id)` can be used to do batched pagination
items, err := client.ListByReplicationProtectedItemsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
