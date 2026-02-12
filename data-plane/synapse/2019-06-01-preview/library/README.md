
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/library` Documentation

The `library` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2019-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/library"
```


### Client Initialization

```go
client := library.NewLibraryClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `LibraryClient.Create`

```go
ctx := context.TODO()
id := library.NewLibraryID("libraryName")

if err := client.CreateThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `LibraryClient.Delete`

```go
ctx := context.TODO()
id := library.NewLibraryID("libraryName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `LibraryClient.Flush`

```go
ctx := context.TODO()
id := library.NewLibraryID("libraryName")

if err := client.FlushThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `LibraryClient.Get`

```go
ctx := context.TODO()
id := library.NewLibraryID("libraryName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LibraryClient.GetOperationResult`

```go
ctx := context.TODO()
id := library.NewLibraryOperationResultID("operationId")

read, err := client.GetOperationResult(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LibraryClient.List`

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
