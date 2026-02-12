
## `github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1.0/revisions` Documentation

The `revisions` SDK allows for interaction with <unknown source data type> `appconfiguration` (API Version `1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1.0/revisions"
```


### Client Initialization

```go
client := revisions.NewRevisionsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `RevisionsClient.CheckRevisions`

```go
ctx := context.TODO()


read, err := client.CheckRevisions(ctx, revisions.DefaultCheckRevisionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RevisionsClient.GetRevisions`

```go
ctx := context.TODO()


// alternatively `client.GetRevisions(ctx, revisions.DefaultGetRevisionsOperationOptions())` can be used to do batched pagination
items, err := client.GetRevisionsComplete(ctx, revisions.DefaultGetRevisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
