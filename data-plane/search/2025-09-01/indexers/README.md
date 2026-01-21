
## `github.com/hashicorp/go-azure-sdk/data-plane/search/2025-09-01/indexers` Documentation

The `indexers` SDK allows for interaction with <unknown source data type> `search` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/search/2025-09-01/indexers"
```


### Client Initialization

```go
client := indexers.NewIndexersClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `IndexersClient.Create`

```go
ctx := context.TODO()

payload := indexers.SearchIndexer{
	// ...
}


read, err := client.Create(ctx, payload, indexers.DefaultCreateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IndexersClient.CreateOrUpdate`

```go
ctx := context.TODO()

payload := indexers.SearchIndexer{
	// ...
}


read, err := client.CreateOrUpdate(ctx, payload, indexers.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IndexersClient.Delete`

```go
ctx := context.TODO()


read, err := client.Delete(ctx, indexers.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IndexersClient.Get`

```go
ctx := context.TODO()


read, err := client.Get(ctx, indexers.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IndexersClient.GetStatus`

```go
ctx := context.TODO()


read, err := client.GetStatus(ctx, indexers.DefaultGetStatusOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IndexersClient.List`

```go
ctx := context.TODO()


read, err := client.List(ctx, indexers.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IndexersClient.Reset`

```go
ctx := context.TODO()


read, err := client.Reset(ctx, indexers.DefaultResetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IndexersClient.Run`

```go
ctx := context.TODO()


read, err := client.Run(ctx, indexers.DefaultRunOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
