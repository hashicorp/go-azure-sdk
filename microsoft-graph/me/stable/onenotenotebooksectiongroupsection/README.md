
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/onenotenotebooksectiongroupsection` Documentation

The `onenotenotebooksectiongroupsection` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/onenotenotebooksectiongroupsection"
```


### Client Initialization

```go
client := onenotenotebooksectiongroupsection.NewOnenoteNotebookSectionGroupSectionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnenoteNotebookSectionGroupSectionClient.CopyOnenoteNotebookSectionGroupSectionToNotebook`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsection.NewMeOnenoteNotebookIdSectionGroupIdSectionID("notebookId", "sectionGroupId", "onenoteSectionId")

payload := onenotenotebooksectiongroupsection.CopyOnenoteNotebookSectionGroupSectionToNotebookRequest{
	// ...
}


read, err := client.CopyOnenoteNotebookSectionGroupSectionToNotebook(ctx, id, payload, onenotenotebooksectiongroupsection.DefaultCopyOnenoteNotebookSectionGroupSectionToNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionClient.CopyOnenoteNotebookSectionGroupSectionToSectionGroup`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsection.NewMeOnenoteNotebookIdSectionGroupIdSectionID("notebookId", "sectionGroupId", "onenoteSectionId")

payload := onenotenotebooksectiongroupsection.CopyOnenoteNotebookSectionGroupSectionToSectionGroupRequest{
	// ...
}


read, err := client.CopyOnenoteNotebookSectionGroupSectionToSectionGroup(ctx, id, payload, onenotenotebooksectiongroupsection.DefaultCopyOnenoteNotebookSectionGroupSectionToSectionGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionClient.CreateOnenoteNotebookSectionGroupSection`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsection.NewMeOnenoteNotebookIdSectionGroupID("notebookId", "sectionGroupId")

payload := onenotenotebooksectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.CreateOnenoteNotebookSectionGroupSection(ctx, id, payload, onenotenotebooksectiongroupsection.DefaultCreateOnenoteNotebookSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionClient.DeleteOnenoteNotebookSectionGroupSection`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsection.NewMeOnenoteNotebookIdSectionGroupIdSectionID("notebookId", "sectionGroupId", "onenoteSectionId")

read, err := client.DeleteOnenoteNotebookSectionGroupSection(ctx, id, onenotenotebooksectiongroupsection.DefaultDeleteOnenoteNotebookSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionClient.GetOnenoteNotebookSectionGroupSection`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsection.NewMeOnenoteNotebookIdSectionGroupIdSectionID("notebookId", "sectionGroupId", "onenoteSectionId")

read, err := client.GetOnenoteNotebookSectionGroupSection(ctx, id, onenotenotebooksectiongroupsection.DefaultGetOnenoteNotebookSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionClient.GetOnenoteNotebookSectionGroupSectionsCount`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsection.NewMeOnenoteNotebookIdSectionGroupID("notebookId", "sectionGroupId")

read, err := client.GetOnenoteNotebookSectionGroupSectionsCount(ctx, id, onenotenotebooksectiongroupsection.DefaultGetOnenoteNotebookSectionGroupSectionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionClient.ListOnenoteNotebookSectionGroupSections`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsection.NewMeOnenoteNotebookIdSectionGroupID("notebookId", "sectionGroupId")

// alternatively `client.ListOnenoteNotebookSectionGroupSections(ctx, id, onenotenotebooksectiongroupsection.DefaultListOnenoteNotebookSectionGroupSectionsOperationOptions())` can be used to do batched pagination
items, err := client.ListOnenoteNotebookSectionGroupSectionsComplete(ctx, id, onenotenotebooksectiongroupsection.DefaultListOnenoteNotebookSectionGroupSectionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionClient.UpdateOnenoteNotebookSectionGroupSection`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsection.NewMeOnenoteNotebookIdSectionGroupIdSectionID("notebookId", "sectionGroupId", "onenoteSectionId")

payload := onenotenotebooksectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.UpdateOnenoteNotebookSectionGroupSection(ctx, id, payload, onenotenotebooksectiongroupsection.DefaultUpdateOnenoteNotebookSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
