
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meonenotenotebooksection` Documentation

The `meonenotenotebooksection` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meonenotenotebooksection"
```


### Client Initialization

```go
client := meonenotenotebooksection.NewMeOnenoteNotebookSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOnenoteNotebookSectionClient.CreateMeOnenoteNotebookByIdSection`

```go
ctx := context.TODO()
id := meonenotenotebooksection.NewMeOnenoteNotebookID("notebookIdValue")

payload := meonenotenotebooksection.OnenoteSection{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionClient.CreateMeOnenoteNotebookByIdSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := meonenotenotebooksection.NewMeOnenoteNotebookSectionID("notebookIdValue", "onenoteSectionIdValue")

payload := meonenotenotebooksection.CreateMeOnenoteNotebookByIdSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionClient.CreateMeOnenoteNotebookByIdSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := meonenotenotebooksection.NewMeOnenoteNotebookSectionID("notebookIdValue", "onenoteSectionIdValue")

payload := meonenotenotebooksection.CreateMeOnenoteNotebookByIdSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionClient.DeleteMeOnenoteNotebookByIdSectionById`

```go
ctx := context.TODO()
id := meonenotenotebooksection.NewMeOnenoteNotebookSectionID("notebookIdValue", "onenoteSectionIdValue")

read, err := client.DeleteMeOnenoteNotebookByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionClient.GetMeOnenoteNotebookByIdSectionById`

```go
ctx := context.TODO()
id := meonenotenotebooksection.NewMeOnenoteNotebookSectionID("notebookIdValue", "onenoteSectionIdValue")

read, err := client.GetMeOnenoteNotebookByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionClient.GetMeOnenoteNotebookByIdSectionCount`

```go
ctx := context.TODO()
id := meonenotenotebooksection.NewMeOnenoteNotebookID("notebookIdValue")

read, err := client.GetMeOnenoteNotebookByIdSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionClient.ListMeOnenoteNotebookByIdSections`

```go
ctx := context.TODO()
id := meonenotenotebooksection.NewMeOnenoteNotebookID("notebookIdValue")

// alternatively `client.ListMeOnenoteNotebookByIdSections(ctx, id)` can be used to do batched pagination
items, err := client.ListMeOnenoteNotebookByIdSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeOnenoteNotebookSectionClient.UpdateMeOnenoteNotebookByIdSectionById`

```go
ctx := context.TODO()
id := meonenotenotebooksection.NewMeOnenoteNotebookSectionID("notebookIdValue", "onenoteSectionIdValue")

payload := meonenotenotebooksection.OnenoteSection{
	// ...
}


read, err := client.UpdateMeOnenoteNotebookByIdSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
