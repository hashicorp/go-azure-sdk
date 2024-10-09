
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2024-02-01/replicationprotectionclusters` Documentation

The `replicationprotectionclusters` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2024-02-01/replicationprotectionclusters"
```


### Client Initialization

```go
client := replicationprotectionclusters.NewReplicationProtectionClustersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationProtectionClustersClient.ApplyRecoveryPoint`

```go
ctx := context.TODO()
id := replicationprotectionclusters.NewReplicationProtectionClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationProtectionClusterName")

payload := replicationprotectionclusters.ApplyClusterRecoveryPointInput{
	// ...
}


if err := client.ApplyRecoveryPointThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationProtectionClustersClient.Create`

```go
ctx := context.TODO()
id := replicationprotectionclusters.NewReplicationProtectionClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationProtectionClusterName")

payload := replicationprotectionclusters.ReplicationProtectionCluster{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationProtectionClustersClient.FailoverCommit`

```go
ctx := context.TODO()
id := replicationprotectionclusters.NewReplicationProtectionClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationProtectionClusterName")

if err := client.FailoverCommitThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationProtectionClustersClient.Get`

```go
ctx := context.TODO()
id := replicationprotectionclusters.NewReplicationProtectionClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationProtectionClusterName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationProtectionClustersClient.List`

```go
ctx := context.TODO()
id := replicationprotectionclusters.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.List(ctx, id, replicationprotectionclusters.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, replicationprotectionclusters.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReplicationProtectionClustersClient.ListByReplicationProtectionContainers`

```go
ctx := context.TODO()
id := replicationprotectionclusters.NewReplicationProtectionContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName")

// alternatively `client.ListByReplicationProtectionContainers(ctx, id)` can be used to do batched pagination
items, err := client.ListByReplicationProtectionContainersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReplicationProtectionClustersClient.Purge`

```go
ctx := context.TODO()
id := replicationprotectionclusters.NewReplicationProtectionClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationProtectionClusterName")

if err := client.PurgeThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationProtectionClustersClient.RepairReplication`

```go
ctx := context.TODO()
id := replicationprotectionclusters.NewReplicationProtectionClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationProtectionClusterName")

if err := client.RepairReplicationThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationProtectionClustersClient.TestFailover`

```go
ctx := context.TODO()
id := replicationprotectionclusters.NewReplicationProtectionClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationProtectionClusterName")

payload := replicationprotectionclusters.ClusterTestFailoverInput{
	// ...
}


if err := client.TestFailoverThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationProtectionClustersClient.TestFailoverCleanup`

```go
ctx := context.TODO()
id := replicationprotectionclusters.NewReplicationProtectionClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationProtectionClusterName")

payload := replicationprotectionclusters.ClusterTestFailoverCleanupInput{
	// ...
}


if err := client.TestFailoverCleanupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationProtectionClustersClient.UnplannedFailover`

```go
ctx := context.TODO()
id := replicationprotectionclusters.NewReplicationProtectionClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationProtectionClusterName")

payload := replicationprotectionclusters.ClusterUnplannedFailoverInput{
	// ...
}


if err := client.UnplannedFailoverThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
