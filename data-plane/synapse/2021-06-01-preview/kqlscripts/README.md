
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/kqlscripts` Documentation

The `kqlscripts` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2021-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/kqlscripts"
```


### Client Initialization

```go
client := kqlscripts.NewKqlScriptsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `KqlScriptsClient.GetAll`

```go
ctx := context.TODO()


// alternatively `client.GetAll(ctx)` can be used to do batched pagination
items, err := client.GetAllComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `KqlScriptsClient.KqlScriptCreateOrUpdate`

```go
ctx := context.TODO()
id := kqlscripts.NewKqlScriptID("kqlScriptName")

payload := kqlscripts.KqlScriptResource{
	// ...
}


if err := client.KqlScriptCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `KqlScriptsClient.KqlScriptDeleteByName`

```go
ctx := context.TODO()
id := kqlscripts.NewKqlScriptID("kqlScriptName")

if err := client.KqlScriptDeleteByNameThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `KqlScriptsClient.KqlScriptGetByName`

```go
ctx := context.TODO()
id := kqlscripts.NewKqlScriptID("kqlScriptName")

read, err := client.KqlScriptGetByName(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KqlScriptsClient.KqlScriptRename`

```go
ctx := context.TODO()
id := kqlscripts.NewKqlScriptID("kqlScriptName")

payload := kqlscripts.ArtifactRenameRequest{
	// ...
}


if err := client.KqlScriptRenameThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
