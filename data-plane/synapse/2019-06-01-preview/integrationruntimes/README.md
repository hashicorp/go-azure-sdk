
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/integrationruntimes` Documentation

The `integrationruntimes` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2019-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/integrationruntimes"
```


### Client Initialization

```go
client := integrationruntimes.NewIntegrationRuntimesClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `IntegrationRuntimesClient.Get`

```go
ctx := context.TODO()
id := integrationruntimes.NewIntegrationRuntimeID("integrationRuntimeName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntegrationRuntimesClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx)` can be used to do batched pagination
items, err := client.ListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
