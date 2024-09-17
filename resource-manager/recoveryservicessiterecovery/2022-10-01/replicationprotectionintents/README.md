
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2022-10-01/replicationprotectionintents` Documentation

The `replicationprotectionintents` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2022-10-01/replicationprotectionintents"
```


### Client Initialization

```go
client := replicationprotectionintents.NewReplicationProtectionIntentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationProtectionIntentsClient.Create`

```go
ctx := context.TODO()
id := replicationprotectionintents.NewReplicationProtectionIntentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationProtectionIntentValue")

payload := replicationprotectionintents.CreateProtectionIntentInput{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationProtectionIntentsClient.Get`

```go
ctx := context.TODO()
id := replicationprotectionintents.NewReplicationProtectionIntentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationProtectionIntentValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationProtectionIntentsClient.List`

```go
ctx := context.TODO()
id := replicationprotectionintents.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.List(ctx, id, replicationprotectionintents.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, replicationprotectionintents.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
