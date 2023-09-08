
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meonenotesectiongroupsection` Documentation

The `meonenotesectiongroupsection` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meonenotesectiongroupsection"
```


### Client Initialization

```go
client := meonenotesectiongroupsection.NewMeOnenoteSectionGroupSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOnenoteSectionGroupSectionClient.CreateMeOnenoteSectionGroupByIdSection`

```go
ctx := context.TODO()
id := meonenotesectiongroupsection.NewMeOnenoteSectionGroupID("sectionGroupIdValue")

payload := meonenotesectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.CreateMeOnenoteSectionGroupByIdSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionGroupSectionClient.CreateMeOnenoteSectionGroupByIdSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := meonenotesectiongroupsection.NewMeOnenoteSectionGroupSectionID("sectionGroupIdValue", "onenoteSectionIdValue")

payload := meonenotesectiongroupsection.CreateMeOnenoteSectionGroupByIdSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateMeOnenoteSectionGroupByIdSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionGroupSectionClient.CreateMeOnenoteSectionGroupByIdSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := meonenotesectiongroupsection.NewMeOnenoteSectionGroupSectionID("sectionGroupIdValue", "onenoteSectionIdValue")

payload := meonenotesectiongroupsection.CreateMeOnenoteSectionGroupByIdSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateMeOnenoteSectionGroupByIdSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionGroupSectionClient.DeleteMeOnenoteSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := meonenotesectiongroupsection.NewMeOnenoteSectionGroupSectionID("sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.DeleteMeOnenoteSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionGroupSectionClient.GetMeOnenoteSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := meonenotesectiongroupsection.NewMeOnenoteSectionGroupSectionID("sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetMeOnenoteSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionGroupSectionClient.GetMeOnenoteSectionGroupByIdSectionCount`

```go
ctx := context.TODO()
id := meonenotesectiongroupsection.NewMeOnenoteSectionGroupID("sectionGroupIdValue")

read, err := client.GetMeOnenoteSectionGroupByIdSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionGroupSectionClient.ListMeOnenoteSectionGroupByIdSections`

```go
ctx := context.TODO()
id := meonenotesectiongroupsection.NewMeOnenoteSectionGroupID("sectionGroupIdValue")

// alternatively `client.ListMeOnenoteSectionGroupByIdSections(ctx, id)` can be used to do batched pagination
items, err := client.ListMeOnenoteSectionGroupByIdSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeOnenoteSectionGroupSectionClient.UpdateMeOnenoteSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := meonenotesectiongroupsection.NewMeOnenoteSectionGroupSectionID("sectionGroupIdValue", "onenoteSectionIdValue")

payload := meonenotesectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.UpdateMeOnenoteSectionGroupByIdSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
