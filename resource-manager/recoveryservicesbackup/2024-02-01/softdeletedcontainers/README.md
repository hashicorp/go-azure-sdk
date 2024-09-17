
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-02-01/softdeletedcontainers` Documentation

The `softdeletedcontainers` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-02-01/softdeletedcontainers"
```


### Client Initialization

```go
client := softdeletedcontainers.NewSoftDeletedContainersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SoftDeletedContainersClient.DeletedProtectionContainersList`

```go
ctx := context.TODO()
id := softdeletedcontainers.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.DeletedProtectionContainersList(ctx, id, softdeletedcontainers.DefaultDeletedProtectionContainersListOperationOptions())` can be used to do batched pagination
items, err := client.DeletedProtectionContainersListComplete(ctx, id, softdeletedcontainers.DefaultDeletedProtectionContainersListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
