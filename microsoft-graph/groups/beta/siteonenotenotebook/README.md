
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotenotebook` Documentation

The `siteonenotenotebook` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotenotebook"
```


### Client Initialization

```go
client := siteonenotenotebook.NewSiteOnenoteNotebookClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteOnenoteNotebookClient.CopySiteOnenoteNotebookNotebook`

```go
ctx := context.TODO()
id := siteonenotenotebook.NewGroupIdSiteIdOnenoteNotebookID("groupId", "siteId", "notebookId")

payload := siteonenotenotebook.CopySiteOnenoteNotebookNotebookRequest{
	// ...
}


read, err := client.CopySiteOnenoteNotebookNotebook(ctx, id, payload, siteonenotenotebook.DefaultCopySiteOnenoteNotebookNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookClient.CreateSiteOnenoteNotebook`

```go
ctx := context.TODO()
id := siteonenotenotebook.NewGroupIdSiteID("groupId", "siteId")

payload := siteonenotenotebook.Notebook{
	// ...
}


read, err := client.CreateSiteOnenoteNotebook(ctx, id, payload, siteonenotenotebook.DefaultCreateSiteOnenoteNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookClient.DeleteSiteOnenoteNotebook`

```go
ctx := context.TODO()
id := siteonenotenotebook.NewGroupIdSiteIdOnenoteNotebookID("groupId", "siteId", "notebookId")

read, err := client.DeleteSiteOnenoteNotebook(ctx, id, siteonenotenotebook.DefaultDeleteSiteOnenoteNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookClient.GetSiteOnenoteNotebook`

```go
ctx := context.TODO()
id := siteonenotenotebook.NewGroupIdSiteIdOnenoteNotebookID("groupId", "siteId", "notebookId")

read, err := client.GetSiteOnenoteNotebook(ctx, id, siteonenotenotebook.DefaultGetSiteOnenoteNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookClient.GetSiteOnenoteNotebooksCount`

```go
ctx := context.TODO()
id := siteonenotenotebook.NewGroupIdSiteID("groupId", "siteId")

read, err := client.GetSiteOnenoteNotebooksCount(ctx, id, siteonenotenotebook.DefaultGetSiteOnenoteNotebooksCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookClient.GetSiteOnenoteNotebooksFromWebUrl`

```go
ctx := context.TODO()
id := siteonenotenotebook.NewGroupIdSiteID("groupId", "siteId")

payload := siteonenotenotebook.GetSiteOnenoteNotebooksFromWebUrlRequest{
	// ...
}


read, err := client.GetSiteOnenoteNotebooksFromWebUrl(ctx, id, payload, siteonenotenotebook.DefaultGetSiteOnenoteNotebooksFromWebUrlOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookClient.ListSiteOnenoteNotebooks`

```go
ctx := context.TODO()
id := siteonenotenotebook.NewGroupIdSiteID("groupId", "siteId")

// alternatively `client.ListSiteOnenoteNotebooks(ctx, id, siteonenotenotebook.DefaultListSiteOnenoteNotebooksOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteOnenoteNotebooksComplete(ctx, id, siteonenotenotebook.DefaultListSiteOnenoteNotebooksOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteOnenoteNotebookClient.UpdateSiteOnenoteNotebook`

```go
ctx := context.TODO()
id := siteonenotenotebook.NewGroupIdSiteIdOnenoteNotebookID("groupId", "siteId", "notebookId")

payload := siteonenotenotebook.Notebook{
	// ...
}


read, err := client.UpdateSiteOnenoteNotebook(ctx, id, payload, siteonenotenotebook.DefaultUpdateSiteOnenoteNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
