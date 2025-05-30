
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2025-02-01/agentpools` Documentation

The `agentpools` SDK allows for interaction with Azure Resource Manager `containerservice` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2025-02-01/agentpools"
```


### Client Initialization

```go
client := agentpools.NewAgentPoolsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AgentPoolsClient.AbortLatestOperation`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "agentPoolName")

if err := client.AbortLatestOperationThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AgentPoolsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "agentPoolName")

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
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "agentPoolName")

if err := client.DeleteThenPoll(ctx, id, agentpools.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `AgentPoolsClient.DeleteMachines`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "agentPoolName")

payload := agentpools.AgentPoolDeleteMachinesParameter{
	// ...
}


if err := client.DeleteMachinesThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AgentPoolsClient.Get`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "agentPoolName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentPoolsClient.GetAvailableAgentPoolVersions`

```go
ctx := context.TODO()
id := commonids.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName")

read, err := client.GetAvailableAgentPoolVersions(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentPoolsClient.GetUpgradeProfile`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "agentPoolName")

read, err := client.GetUpgradeProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentPoolsClient.List`

```go
ctx := context.TODO()
id := commonids.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AgentPoolsClient.UpgradeNodeImageVersion`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "agentPoolName")

if err := client.UpgradeNodeImageVersionThenPoll(ctx, id); err != nil {
	// handle the error
}
```
