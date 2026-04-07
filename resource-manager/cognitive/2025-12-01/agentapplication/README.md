
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/agentapplication` Documentation

The `agentapplication` SDK allows for interaction with Azure Resource Manager `cognitive` (API Version `2025-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/agentapplication"
```


### Client Initialization

```go
client := agentapplication.NewAgentApplicationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AgentApplicationClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := agentapplication.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "applicationName")

payload := agentapplication.AgentApplication{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AgentApplicationClient.Delete`

```go
ctx := context.TODO()
id := agentapplication.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "applicationName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AgentApplicationClient.Disable`

```go
ctx := context.TODO()
id := agentapplication.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "applicationName")

read, err := client.Disable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentApplicationClient.Enable`

```go
ctx := context.TODO()
id := agentapplication.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "applicationName")

read, err := client.Enable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentApplicationClient.Get`

```go
ctx := context.TODO()
id := agentapplication.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "applicationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentApplicationClient.List`

```go
ctx := context.TODO()
id := agentapplication.NewProjectID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName")

// alternatively `client.List(ctx, id, agentapplication.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, agentapplication.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AgentApplicationClient.ListAgents`

```go
ctx := context.TODO()
id := agentapplication.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "applicationName")

// alternatively `client.ListAgents(ctx, id)` can be used to do batched pagination
items, err := client.ListAgentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
