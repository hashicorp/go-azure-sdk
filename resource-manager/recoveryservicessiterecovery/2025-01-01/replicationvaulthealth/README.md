
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2025-01-01/replicationvaulthealth` Documentation

The `replicationvaulthealth` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2025-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2025-01-01/replicationvaulthealth"
```


### Client Initialization

```go
client := replicationvaulthealth.NewReplicationVaultHealthClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationVaultHealthClient.Get`

```go
ctx := context.TODO()
id := replicationvaulthealth.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationVaultHealthClient.Refresh`

```go
ctx := context.TODO()
id := replicationvaulthealth.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

if err := client.RefreshThenPoll(ctx, id); err != nil {
	// handle the error
}
```
