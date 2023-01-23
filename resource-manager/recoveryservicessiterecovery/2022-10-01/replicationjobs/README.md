
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2022-10-01/replicationjobs` Documentation

The `replicationjobs` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicessiterecovery` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2022-10-01/replicationjobs"
```


### Client Initialization

```go
client := replicationjobs.NewReplicationJobsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationJobsClient.Cancel`

```go
ctx := context.TODO()
id := replicationjobs.NewReplicationJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationJobValue")

if err := client.CancelThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationJobsClient.Export`

```go
ctx := context.TODO()
id := replicationjobs.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

payload := replicationjobs.JobQueryParameter{
	// ...
}


if err := client.ExportThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationJobsClient.Get`

```go
ctx := context.TODO()
id := replicationjobs.NewReplicationJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationJobValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationJobsClient.List`

```go
ctx := context.TODO()
id := replicationjobs.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.List(ctx, id, replicationjobs.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, replicationjobs.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReplicationJobsClient.Restart`

```go
ctx := context.TODO()
id := replicationjobs.NewReplicationJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationJobValue")

if err := client.RestartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ReplicationJobsClient.Resume`

```go
ctx := context.TODO()
id := replicationjobs.NewReplicationJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationJobValue")

payload := replicationjobs.ResumeJobParams{
	// ...
}


if err := client.ResumeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
