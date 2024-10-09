
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/fileresource` Documentation

The `fileresource` SDK allows for interaction with Azure Resource Manager `datamigration` (API Version `2021-06-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/fileresource"
```


### Client Initialization

```go
client := fileresource.NewFileResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FileResourceClient.FilesDelete`

```go
ctx := context.TODO()
id := fileresource.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName", "fileName")

read, err := client.FilesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FileResourceClient.FilesGet`

```go
ctx := context.TODO()
id := fileresource.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName", "fileName")

read, err := client.FilesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FileResourceClient.FilesList`

```go
ctx := context.TODO()
id := fileresource.NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName")

// alternatively `client.FilesList(ctx, id)` can be used to do batched pagination
items, err := client.FilesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FileResourceClient.FilesRead`

```go
ctx := context.TODO()
id := fileresource.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName", "fileName")

read, err := client.FilesRead(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FileResourceClient.FilesReadWrite`

```go
ctx := context.TODO()
id := fileresource.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName", "fileName")

read, err := client.FilesReadWrite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FileResourceClient.FilesUpdate`

```go
ctx := context.TODO()
id := fileresource.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName", "fileName")

payload := fileresource.ProjectFile{
	// ...
}


read, err := client.FilesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
