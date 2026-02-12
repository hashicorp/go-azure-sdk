
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/sqlscripts` Documentation

The `sqlscripts` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2019-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/sqlscripts"
```


### Client Initialization

```go
client := sqlscripts.NewSqlScriptsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlScriptsClient.SqlScriptCreateOrUpdateSqlScript`

```go
ctx := context.TODO()
id := sqlscripts.NewSqlScriptID("sqlScriptName")

payload := sqlscripts.SqlScriptResource{
	// ...
}


if err := client.SqlScriptCreateOrUpdateSqlScriptThenPoll(ctx, id, payload, sqlscripts.DefaultSqlScriptCreateOrUpdateSqlScriptOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `SqlScriptsClient.SqlScriptDeleteSqlScript`

```go
ctx := context.TODO()
id := sqlscripts.NewSqlScriptID("sqlScriptName")

if err := client.SqlScriptDeleteSqlScriptThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SqlScriptsClient.SqlScriptGetSqlScript`

```go
ctx := context.TODO()
id := sqlscripts.NewSqlScriptID("sqlScriptName")

read, err := client.SqlScriptGetSqlScript(ctx, id, sqlscripts.DefaultSqlScriptGetSqlScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlScriptsClient.SqlScriptGetSqlScriptsByWorkspace`

```go
ctx := context.TODO()


// alternatively `client.SqlScriptGetSqlScriptsByWorkspace(ctx)` can be used to do batched pagination
items, err := client.SqlScriptGetSqlScriptsByWorkspaceComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SqlScriptsClient.SqlScriptRenameSqlScript`

```go
ctx := context.TODO()
id := sqlscripts.NewSqlScriptID("sqlScriptName")

payload := sqlscripts.ArtifactRenameRequest{
	// ...
}


if err := client.SqlScriptRenameSqlScriptThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
