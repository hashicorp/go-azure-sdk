
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2022-05-01/replicationappliances` Documentation

The `replicationappliances` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicessiterecovery` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2022-05-01/replicationappliances"
```


### Client Initialization

```go
client := replicationappliances.NewReplicationAppliancesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationAppliancesClient.List`

```go
ctx := context.TODO()
id := replicationappliances.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue")

// alternatively `client.List(ctx, id, replicationappliances.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, replicationappliances.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
