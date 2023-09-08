
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meonenotenotebooksectiongroupsection` Documentation

The `meonenotenotebooksectiongroupsection` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meonenotenotebooksectiongroupsection"
```


### Client Initialization

```go
client := meonenotenotebooksectiongroupsection.NewMeOnenoteNotebookSectionGroupSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionClient.CreateMeOnenoteNotebookByIdSectionGroupByIdSection`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsection.NewMeOnenoteNotebookSectionGroupID("notebookIdValue", "sectionGroupIdValue")

payload := meonenotenotebooksectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdSectionGroupByIdSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionClient.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsection.NewMeOnenoteNotebookSectionGroupSectionID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := meonenotenotebooksectiongroupsection.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionClient.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsection.NewMeOnenoteNotebookSectionGroupSectionID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := meonenotenotebooksectiongroupsection.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionClient.DeleteMeOnenoteNotebookByIdSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsection.NewMeOnenoteNotebookSectionGroupSectionID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.DeleteMeOnenoteNotebookByIdSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionClient.GetMeOnenoteNotebookByIdSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsection.NewMeOnenoteNotebookSectionGroupSectionID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetMeOnenoteNotebookByIdSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionClient.GetMeOnenoteNotebookByIdSectionGroupByIdSectionCount`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsection.NewMeOnenoteNotebookSectionGroupID("notebookIdValue", "sectionGroupIdValue")

read, err := client.GetMeOnenoteNotebookByIdSectionGroupByIdSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionClient.ListMeOnenoteNotebookByIdSectionGroupByIdSections`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsection.NewMeOnenoteNotebookSectionGroupID("notebookIdValue", "sectionGroupIdValue")

// alternatively `client.ListMeOnenoteNotebookByIdSectionGroupByIdSections(ctx, id)` can be used to do batched pagination
items, err := client.ListMeOnenoteNotebookByIdSectionGroupByIdSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionClient.UpdateMeOnenoteNotebookByIdSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsection.NewMeOnenoteNotebookSectionGroupSectionID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := meonenotenotebooksectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.UpdateMeOnenoteNotebookByIdSectionGroupByIdSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
