
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotenotebook` Documentation

The `grouponenotenotebook` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotenotebook"
```


### Client Initialization

```go
client := grouponenotenotebook.NewGroupOnenoteNotebookClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupOnenoteNotebookClient.CreateGroupByIdOnenoteNotebook`

```go
ctx := context.TODO()
id := grouponenotenotebook.NewGroupID("groupIdValue")

payload := grouponenotenotebook.Notebook{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookClient.CreateGroupByIdOnenoteNotebookByIdCopyNotebook`

```go
ctx := context.TODO()
id := grouponenotenotebook.NewGroupOnenoteNotebookID("groupIdValue", "notebookIdValue")

payload := grouponenotenotebook.CreateGroupByIdOnenoteNotebookByIdCopyNotebookRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdCopyNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookClient.DeleteGroupByIdOnenoteNotebookById`

```go
ctx := context.TODO()
id := grouponenotenotebook.NewGroupOnenoteNotebookID("groupIdValue", "notebookIdValue")

read, err := client.DeleteGroupByIdOnenoteNotebookById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookClient.GetGroupByIdOnenoteNotebookById`

```go
ctx := context.TODO()
id := grouponenotenotebook.NewGroupOnenoteNotebookID("groupIdValue", "notebookIdValue")

read, err := client.GetGroupByIdOnenoteNotebookById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookClient.GetGroupByIdOnenoteNotebookCount`

```go
ctx := context.TODO()
id := grouponenotenotebook.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdOnenoteNotebookCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookClient.GetGroupByIdOnenoteNotebooksNotebookFromWebUrl`

```go
ctx := context.TODO()
id := grouponenotenotebook.NewGroupID("groupIdValue")

payload := grouponenotenotebook.GetGroupByIdOnenoteNotebooksNotebookFromWebUrlRequest{
	// ...
}


read, err := client.GetGroupByIdOnenoteNotebooksNotebookFromWebUrl(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookClient.ListGroupByIdOnenoteNotebooks`

```go
ctx := context.TODO()
id := grouponenotenotebook.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdOnenoteNotebooks(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdOnenoteNotebooksComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupOnenoteNotebookClient.UpdateGroupByIdOnenoteNotebookById`

```go
ctx := context.TODO()
id := grouponenotenotebook.NewGroupOnenoteNotebookID("groupIdValue", "notebookIdValue")

payload := grouponenotenotebook.Notebook{
	// ...
}


read, err := client.UpdateGroupByIdOnenoteNotebookById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
