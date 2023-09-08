
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotenotebooksection` Documentation

The `grouponenotenotebooksection` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotenotebooksection"
```


### Client Initialization

```go
client := grouponenotenotebooksection.NewGroupOnenoteNotebookSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupOnenoteNotebookSectionClient.CreateGroupByIdOnenoteNotebookByIdSection`

```go
ctx := context.TODO()
id := grouponenotenotebooksection.NewGroupOnenoteNotebookID("groupIdValue", "notebookIdValue")

payload := grouponenotenotebooksection.OnenoteSection{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionClient.CreateGroupByIdOnenoteNotebookByIdSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := grouponenotenotebooksection.NewGroupOnenoteNotebookSectionID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue")

payload := grouponenotenotebooksection.CreateGroupByIdOnenoteNotebookByIdSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionClient.CreateGroupByIdOnenoteNotebookByIdSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := grouponenotenotebooksection.NewGroupOnenoteNotebookSectionID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue")

payload := grouponenotenotebooksection.CreateGroupByIdOnenoteNotebookByIdSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionClient.DeleteGroupByIdOnenoteNotebookByIdSectionById`

```go
ctx := context.TODO()
id := grouponenotenotebooksection.NewGroupOnenoteNotebookSectionID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue")

read, err := client.DeleteGroupByIdOnenoteNotebookByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionClient.GetGroupByIdOnenoteNotebookByIdSectionById`

```go
ctx := context.TODO()
id := grouponenotenotebooksection.NewGroupOnenoteNotebookSectionID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdOnenoteNotebookByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionClient.GetGroupByIdOnenoteNotebookByIdSectionCount`

```go
ctx := context.TODO()
id := grouponenotenotebooksection.NewGroupOnenoteNotebookID("groupIdValue", "notebookIdValue")

read, err := client.GetGroupByIdOnenoteNotebookByIdSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionClient.ListGroupByIdOnenoteNotebookByIdSections`

```go
ctx := context.TODO()
id := grouponenotenotebooksection.NewGroupOnenoteNotebookID("groupIdValue", "notebookIdValue")

// alternatively `client.ListGroupByIdOnenoteNotebookByIdSections(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdOnenoteNotebookByIdSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupOnenoteNotebookSectionClient.UpdateGroupByIdOnenoteNotebookByIdSectionById`

```go
ctx := context.TODO()
id := grouponenotenotebooksection.NewGroupOnenoteNotebookSectionID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue")

payload := grouponenotenotebooksection.OnenoteSection{
	// ...
}


read, err := client.UpdateGroupByIdOnenoteNotebookByIdSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
