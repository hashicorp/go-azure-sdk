
## `github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/files` Documentation

The `files` SDK allows for interaction with <unknown source data type> `batch` (API Version `2022-01-01.15.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/files"
```


### Client Initialization

```go
client := files.NewFilesClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `FilesClient.FileDeleteFromComputeNode`

```go
ctx := context.TODO()
id := files.NewNodeFileID("https://endpoint-url.example.com", "poolId", "nodeId")

read, err := client.FileDeleteFromComputeNode(ctx, id, files.DefaultFileDeleteFromComputeNodeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FilesClient.FileDeleteFromTask`

```go
ctx := context.TODO()
id := files.NewFileID("https://endpoint-url.example.com", "jobId", "taskId")

read, err := client.FileDeleteFromTask(ctx, id, files.DefaultFileDeleteFromTaskOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FilesClient.FileGetFromComputeNode`

```go
ctx := context.TODO()
id := files.NewNodeFileID("https://endpoint-url.example.com", "poolId", "nodeId")

read, err := client.FileGetFromComputeNode(ctx, id, files.DefaultFileGetFromComputeNodeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FilesClient.FileGetFromTask`

```go
ctx := context.TODO()
id := files.NewFileID("https://endpoint-url.example.com", "jobId", "taskId")

read, err := client.FileGetFromTask(ctx, id, files.DefaultFileGetFromTaskOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FilesClient.FileGetPropertiesFromComputeNode`

```go
ctx := context.TODO()
id := files.NewNodeFileID("https://endpoint-url.example.com", "poolId", "nodeId")

read, err := client.FileGetPropertiesFromComputeNode(ctx, id, files.DefaultFileGetPropertiesFromComputeNodeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FilesClient.FileGetPropertiesFromTask`

```go
ctx := context.TODO()
id := files.NewFileID("https://endpoint-url.example.com", "jobId", "taskId")

read, err := client.FileGetPropertiesFromTask(ctx, id, files.DefaultFileGetPropertiesFromTaskOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FilesClient.FileListFromComputeNode`

```go
ctx := context.TODO()
id := files.NewNodeID("poolId", "nodeId")

// alternatively `client.FileListFromComputeNode(ctx, id, files.DefaultFileListFromComputeNodeOperationOptions())` can be used to do batched pagination
items, err := client.FileListFromComputeNodeComplete(ctx, id, files.DefaultFileListFromComputeNodeOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FilesClient.FileListFromTask`

```go
ctx := context.TODO()
id := files.NewTaskID("jobId", "taskId")

// alternatively `client.FileListFromTask(ctx, id, files.DefaultFileListFromTaskOperationOptions())` can be used to do batched pagination
items, err := client.FileListFromTaskComplete(ctx, id, files.DefaultFileListFromTaskOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
