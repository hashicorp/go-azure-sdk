
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/agentdeployment` Documentation

The `agentdeployment` SDK allows for interaction with Azure Resource Manager `cognitive` (API Version `2025-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/agentdeployment"
```


### Client Initialization

```go
client := agentdeployment.NewAgentDeploymentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AgentDeploymentClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := agentdeployment.NewAgentDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "applicationName", "agentDeploymentName")

payload := agentdeployment.AgentDeployment{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AgentDeploymentClient.Delete`

```go
ctx := context.TODO()
id := agentdeployment.NewAgentDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "applicationName", "agentDeploymentName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AgentDeploymentClient.Get`

```go
ctx := context.TODO()
id := agentdeployment.NewAgentDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "applicationName", "agentDeploymentName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentDeploymentClient.List`

```go
ctx := context.TODO()
id := agentdeployment.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "applicationName")

// alternatively `client.List(ctx, id, agentdeployment.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, agentdeployment.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AgentDeploymentClient.Start`

```go
ctx := context.TODO()
id := agentdeployment.NewAgentDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "applicationName", "agentDeploymentName")

read, err := client.Start(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentDeploymentClient.Stop`

```go
ctx := context.TODO()
id := agentdeployment.NewAgentDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "applicationName", "agentDeploymentName")

read, err := client.Stop(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
