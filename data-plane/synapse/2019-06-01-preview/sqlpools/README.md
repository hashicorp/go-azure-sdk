
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/sqlpools` Documentation

The `sqlpools` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2019-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/sqlpools"
```


### Client Initialization

```go
client := sqlpools.NewSqlPoolsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsClient.Get`

```go
ctx := context.TODO()
id := sqlpools.NewSqlPoolID("sqlPoolName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsClient.List`

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
