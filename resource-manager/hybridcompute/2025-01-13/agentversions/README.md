
## `github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/agentversions` Documentation

The `agentversions` SDK allows for interaction with Azure Resource Manager `hybridcompute` (API Version `2025-01-13`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/agentversions"
```


### Client Initialization

```go
client := agentversions.NewAgentVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AgentVersionsClient.AgentVersionGet`

```go
ctx := context.TODO()
id := agentversions.NewAgentVersionID("osTypeName", "agentVersionName")

read, err := client.AgentVersionGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentVersionsClient.AgentVersionList`

```go
ctx := context.TODO()
id := agentversions.NewOsTypeID("osTypeName")

// alternatively `client.AgentVersionList(ctx, id)` can be used to do batched pagination
items, err := client.AgentVersionListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
