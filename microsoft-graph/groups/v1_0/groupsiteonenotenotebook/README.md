
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupsiteonenotenotebook` Documentation

The `groupsiteonenotenotebook` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupsiteonenotenotebook"
```


### Client Initialization

```go
client := groupsiteonenotenotebook.NewGroupSiteOnenoteNotebookClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteOnenoteNotebookClient.CreateGroupByIdSiteByIdOnenoteNotebook`

```go
ctx := context.TODO()
id := groupsiteonenotenotebook.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteonenotenotebook.Notebook{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdCopyNotebook`

```go
ctx := context.TODO()
id := groupsiteonenotenotebook.NewGroupSiteOnenoteNotebookID("groupIdValue", "siteIdValue", "notebookIdValue")

payload := groupsiteonenotenotebook.CreateGroupByIdSiteByIdOnenoteNotebookByIdCopyNotebookRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdCopyNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookClient.DeleteGroupByIdSiteByIdOnenoteNotebookById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebook.NewGroupSiteOnenoteNotebookID("groupIdValue", "siteIdValue", "notebookIdValue")

read, err := client.DeleteGroupByIdSiteByIdOnenoteNotebookById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookClient.GetGroupByIdSiteByIdOnenoteNotebookById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebook.NewGroupSiteOnenoteNotebookID("groupIdValue", "siteIdValue", "notebookIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteNotebookById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookClient.GetGroupByIdSiteByIdOnenoteNotebookCount`

```go
ctx := context.TODO()
id := groupsiteonenotenotebook.NewGroupSiteID("groupIdValue", "siteIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteNotebookCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookClient.GetGroupByIdSiteByIdOnenoteNotebooksNotebookFromWebUrl`

```go
ctx := context.TODO()
id := groupsiteonenotenotebook.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteonenotenotebook.GetGroupByIdSiteByIdOnenoteNotebooksNotebookFromWebUrlRequest{
	// ...
}


read, err := client.GetGroupByIdSiteByIdOnenoteNotebooksNotebookFromWebUrl(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookClient.ListGroupByIdSiteByIdOnenoteNotebooks`

```go
ctx := context.TODO()
id := groupsiteonenotenotebook.NewGroupSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListGroupByIdSiteByIdOnenoteNotebooks(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdOnenoteNotebooksComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteOnenoteNotebookClient.UpdateGroupByIdSiteByIdOnenoteNotebookById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebook.NewGroupSiteOnenoteNotebookID("groupIdValue", "siteIdValue", "notebookIdValue")

payload := groupsiteonenotenotebook.Notebook{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdOnenoteNotebookById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
