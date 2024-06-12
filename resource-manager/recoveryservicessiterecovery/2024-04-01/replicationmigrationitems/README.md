
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2024-04-01/replicationmigrationitems` Documentation

The `replicationmigrationitems` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicessiterecovery` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2024-04-01/replicationmigrationitems"
```


### Client Initialization

```go
client := replicationmigrationitems.NewReplicationMigrationItemsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationMigrationItemsClient.Create`

```go
ctx := context.TODO()
id := replicationmigrationitems.NewReplicationMigrationItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationMigrationItemValue")

payload := replicationmigrationitems.EnableMigrationInput{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationMigrationItemsClient.Delete`

```go
ctx := context.TODO()
id := replicationmigrationitems.NewReplicationMigrationItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationMigrationItemValue")

if err := client.DeleteThenPoll(ctx, id, replicationmigrationitems.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationMigrationItemsClient.Get`

```go
ctx := context.TODO()
id := replicationmigrationitems.NewReplicationMigrationItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationMigrationItemValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationMigrationItemsClient.List`

```go
ctx := context.TODO()
id := replicationmigrationitems.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.List(ctx, id, replicationmigrationitems.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, replicationmigrationitems.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReplicationMigrationItemsClient.ListByReplicationProtectionContainers`

```go
ctx := context.TODO()
id := replicationmigrationitems.NewReplicationProtectionContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue")

// alternatively `client.ListByReplicationProtectionContainers(ctx, id, replicationmigrationitems.DefaultListByReplicationProtectionContainersOperationOptions())` can be used to do batched pagination
items, err := client.ListByReplicationProtectionContainersComplete(ctx, id, replicationmigrationitems.DefaultListByReplicationProtectionContainersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReplicationMigrationItemsClient.Migrate`

```go
ctx := context.TODO()
id := replicationmigrationitems.NewReplicationMigrationItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationMigrationItemValue")

payload := replicationmigrationitems.MigrateInput{
	// ...
}


if err := client.MigrateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationMigrationItemsClient.PauseReplication`

```go
ctx := context.TODO()
id := replicationmigrationitems.NewReplicationMigrationItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationMigrationItemValue")

payload := replicationmigrationitems.PauseReplicationInput{
	// ...
}


if err := client.PauseReplicationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationMigrationItemsClient.ResumeReplication`

```go
ctx := context.TODO()
id := replicationmigrationitems.NewReplicationMigrationItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationMigrationItemValue")

payload := replicationmigrationitems.ResumeReplicationInput{
	// ...
}


if err := client.ResumeReplicationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationMigrationItemsClient.Resync`

```go
ctx := context.TODO()
id := replicationmigrationitems.NewReplicationMigrationItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationMigrationItemValue")

payload := replicationmigrationitems.ResyncInput{
	// ...
}


if err := client.ResyncThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationMigrationItemsClient.TestMigrate`

```go
ctx := context.TODO()
id := replicationmigrationitems.NewReplicationMigrationItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationMigrationItemValue")

payload := replicationmigrationitems.TestMigrateInput{
	// ...
}


if err := client.TestMigrateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationMigrationItemsClient.TestMigrateCleanup`

```go
ctx := context.TODO()
id := replicationmigrationitems.NewReplicationMigrationItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationMigrationItemValue")

payload := replicationmigrationitems.TestMigrateCleanupInput{
	// ...
}


if err := client.TestMigrateCleanupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationMigrationItemsClient.Update`

```go
ctx := context.TODO()
id := replicationmigrationitems.NewReplicationMigrationItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationMigrationItemValue")

payload := replicationmigrationitems.UpdateMigrationItemInput{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
