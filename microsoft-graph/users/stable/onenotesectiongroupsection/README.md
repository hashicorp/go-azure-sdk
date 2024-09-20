
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onenotesectiongroupsection` Documentation

The `onenotesectiongroupsection` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onenotesectiongroupsection"
```


### Client Initialization

```go
client := onenotesectiongroupsection.NewOnenoteSectionGroupSectionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnenoteSectionGroupSectionClient.CopyOnenoteSectionGroupSectionToNotebook`

```go
ctx := context.TODO()
id := onenotesectiongroupsection.NewUserIdOnenoteSectionGroupIdSectionID("userId", "sectionGroupId", "onenoteSectionId")

payload := onenotesectiongroupsection.CopyOnenoteSectionGroupSectionToNotebookRequest{
	// ...
}


read, err := client.CopyOnenoteSectionGroupSectionToNotebook(ctx, id, payload, onenotesectiongroupsection.DefaultCopyOnenoteSectionGroupSectionToNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionGroupSectionClient.CopyOnenoteSectionGroupSectionToSectionGroup`

```go
ctx := context.TODO()
id := onenotesectiongroupsection.NewUserIdOnenoteSectionGroupIdSectionID("userId", "sectionGroupId", "onenoteSectionId")

payload := onenotesectiongroupsection.CopyOnenoteSectionGroupSectionToSectionGroupRequest{
	// ...
}


read, err := client.CopyOnenoteSectionGroupSectionToSectionGroup(ctx, id, payload, onenotesectiongroupsection.DefaultCopyOnenoteSectionGroupSectionToSectionGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionGroupSectionClient.CreateOnenoteSectionGroupSection`

```go
ctx := context.TODO()
id := onenotesectiongroupsection.NewUserIdOnenoteSectionGroupID("userId", "sectionGroupId")

payload := onenotesectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.CreateOnenoteSectionGroupSection(ctx, id, payload, onenotesectiongroupsection.DefaultCreateOnenoteSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionGroupSectionClient.DeleteOnenoteSectionGroupSection`

```go
ctx := context.TODO()
id := onenotesectiongroupsection.NewUserIdOnenoteSectionGroupIdSectionID("userId", "sectionGroupId", "onenoteSectionId")

read, err := client.DeleteOnenoteSectionGroupSection(ctx, id, onenotesectiongroupsection.DefaultDeleteOnenoteSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionGroupSectionClient.GetOnenoteSectionGroupSection`

```go
ctx := context.TODO()
id := onenotesectiongroupsection.NewUserIdOnenoteSectionGroupIdSectionID("userId", "sectionGroupId", "onenoteSectionId")

read, err := client.GetOnenoteSectionGroupSection(ctx, id, onenotesectiongroupsection.DefaultGetOnenoteSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionGroupSectionClient.GetOnenoteSectionGroupSectionsCount`

```go
ctx := context.TODO()
id := onenotesectiongroupsection.NewUserIdOnenoteSectionGroupID("userId", "sectionGroupId")

read, err := client.GetOnenoteSectionGroupSectionsCount(ctx, id, onenotesectiongroupsection.DefaultGetOnenoteSectionGroupSectionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionGroupSectionClient.ListOnenoteSectionGroupSections`

```go
ctx := context.TODO()
id := onenotesectiongroupsection.NewUserIdOnenoteSectionGroupID("userId", "sectionGroupId")

// alternatively `client.ListOnenoteSectionGroupSections(ctx, id, onenotesectiongroupsection.DefaultListOnenoteSectionGroupSectionsOperationOptions())` can be used to do batched pagination
items, err := client.ListOnenoteSectionGroupSectionsComplete(ctx, id, onenotesectiongroupsection.DefaultListOnenoteSectionGroupSectionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnenoteSectionGroupSectionClient.UpdateOnenoteSectionGroupSection`

```go
ctx := context.TODO()
id := onenotesectiongroupsection.NewUserIdOnenoteSectionGroupIdSectionID("userId", "sectionGroupId", "onenoteSectionId")

payload := onenotesectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.UpdateOnenoteSectionGroupSection(ctx, id, payload, onenotesectiongroupsection.DefaultUpdateOnenoteSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
