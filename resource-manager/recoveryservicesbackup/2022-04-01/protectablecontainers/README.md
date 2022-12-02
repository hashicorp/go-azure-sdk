
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2022-04-01/protectablecontainers` Documentation

The `protectablecontainers` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2022-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2022-04-01/protectablecontainers"
```


### Client Initialization

```go
client := protectablecontainers.NewProtectableContainersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProtectableContainersClient.List`

```go
ctx := context.TODO()
id := protectablecontainers.NewBackupFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "fabricValue")

// alternatively `client.List(ctx, id, protectablecontainers.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, protectablecontainers.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
