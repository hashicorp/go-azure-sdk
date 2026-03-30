
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/agentpools` Documentation

The `agentpools` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/agentpools"
```


### Client Initialization

```go
client := agentpools.NewAgentPoolsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AgentPoolsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "agentPoolName")

payload := agentpools.AgentPool{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, agentpools.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `AgentPoolsClient.Delete`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "agentPoolName")

if err := client.DeleteThenPoll(ctx, id, agentpools.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `AgentPoolsClient.Get`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "agentPoolName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentPoolsClient.ListByKubernetesCluster`

```go
ctx := context.TODO()
id := agentpools.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName")

// alternatively `client.ListByKubernetesCluster(ctx, id, agentpools.DefaultListByKubernetesClusterOperationOptions())` can be used to do batched pagination
items, err := client.ListByKubernetesClusterComplete(ctx, id, agentpools.DefaultListByKubernetesClusterOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AgentPoolsClient.Update`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "agentPoolName")

payload := agentpools.AgentPoolPatchParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, agentpools.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
