
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/sparkjobdefinitions` Documentation

The `sparkjobdefinitions` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2019-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/sparkjobdefinitions"
```


### Client Initialization

```go
client := sparkjobdefinitions.NewSparkJobDefinitionsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `SparkJobDefinitionsClient.SparkJobDefinitionCreateOrUpdateSparkJobDefinition`

```go
ctx := context.TODO()
id := sparkjobdefinitions.NewSparkJobDefinitionID("sparkJobDefinitionName")

payload := sparkjobdefinitions.SparkJobDefinitionResource{
	// ...
}


if err := client.SparkJobDefinitionCreateOrUpdateSparkJobDefinitionThenPoll(ctx, id, payload, sparkjobdefinitions.DefaultSparkJobDefinitionCreateOrUpdateSparkJobDefinitionOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `SparkJobDefinitionsClient.SparkJobDefinitionDebugSparkJobDefinition`

```go
ctx := context.TODO()

payload := sparkjobdefinitions.SparkJobDefinitionResource{
	// ...
}


if err := client.SparkJobDefinitionDebugSparkJobDefinitionThenPoll(ctx, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SparkJobDefinitionsClient.SparkJobDefinitionDeleteSparkJobDefinition`

```go
ctx := context.TODO()
id := sparkjobdefinitions.NewSparkJobDefinitionID("sparkJobDefinitionName")

if err := client.SparkJobDefinitionDeleteSparkJobDefinitionThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SparkJobDefinitionsClient.SparkJobDefinitionExecuteSparkJobDefinition`

```go
ctx := context.TODO()
id := sparkjobdefinitions.NewSparkJobDefinitionID("sparkJobDefinitionName")

if err := client.SparkJobDefinitionExecuteSparkJobDefinitionThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SparkJobDefinitionsClient.SparkJobDefinitionGetSparkJobDefinition`

```go
ctx := context.TODO()
id := sparkjobdefinitions.NewSparkJobDefinitionID("sparkJobDefinitionName")

read, err := client.SparkJobDefinitionGetSparkJobDefinition(ctx, id, sparkjobdefinitions.DefaultSparkJobDefinitionGetSparkJobDefinitionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SparkJobDefinitionsClient.SparkJobDefinitionGetSparkJobDefinitionsByWorkspace`

```go
ctx := context.TODO()


// alternatively `client.SparkJobDefinitionGetSparkJobDefinitionsByWorkspace(ctx)` can be used to do batched pagination
items, err := client.SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SparkJobDefinitionsClient.SparkJobDefinitionRenameSparkJobDefinition`

```go
ctx := context.TODO()
id := sparkjobdefinitions.NewSparkJobDefinitionID("sparkJobDefinitionName")

payload := sparkjobdefinitions.ArtifactRenameRequest{
	// ...
}


if err := client.SparkJobDefinitionRenameSparkJobDefinitionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
