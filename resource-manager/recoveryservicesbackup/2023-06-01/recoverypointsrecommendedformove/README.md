
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-06-01/recoverypointsrecommendedformove` Documentation

The `recoverypointsrecommendedformove` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-06-01/recoverypointsrecommendedformove"
```


### Client Initialization

```go
client := recoverypointsrecommendedformove.NewRecoveryPointsRecommendedForMoveClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecoveryPointsRecommendedForMoveClient.List`

```go
ctx := context.TODO()
id := recoverypointsrecommendedformove.NewProtectedItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupFabricValue", "protectionContainerValue", "protectedItemValue")

payload := recoverypointsrecommendedformove.ListRecoveryPointsRecommendedForMoveRequest{
	// ...
}


// alternatively `client.List(ctx, id, payload)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
