
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpools` Documentation

The `sqlpools` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpools"
```


### Client Initialization

```go
client := sqlpools.NewSqlPoolsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsClient.Create`

```go
ctx := context.TODO()
id := sqlpools.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "sqlPoolName")

payload := sqlpools.SqlPool{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SqlPoolsClient.Delete`

```go
ctx := context.TODO()
id := sqlpools.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "sqlPoolName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SqlPoolsClient.Get`

```go
ctx := context.TODO()
id := sqlpools.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "sqlPoolName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsClient.ListByWorkspace`

```go
ctx := context.TODO()
id := sqlpools.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.ListByWorkspace(ctx, id)` can be used to do batched pagination
items, err := client.ListByWorkspaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SqlPoolsClient.Pause`

```go
ctx := context.TODO()
id := sqlpools.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "sqlPoolName")

if err := client.PauseThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SqlPoolsClient.Resume`

```go
ctx := context.TODO()
id := sqlpools.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "sqlPoolName")

if err := client.ResumeThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SqlPoolsClient.SqlPoolMetadataSyncConfigsCreate`

```go
ctx := context.TODO()
id := sqlpools.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "sqlPoolName")

payload := sqlpools.MetadataSyncConfig{
	// ...
}


read, err := client.SqlPoolMetadataSyncConfigsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsClient.SqlPoolMetadataSyncConfigsGet`

```go
ctx := context.TODO()
id := sqlpools.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "sqlPoolName")

read, err := client.SqlPoolMetadataSyncConfigsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsClient.Update`

```go
ctx := context.TODO()
id := sqlpools.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "sqlPoolName")

payload := sqlpools.SqlPoolPatchInfo{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
