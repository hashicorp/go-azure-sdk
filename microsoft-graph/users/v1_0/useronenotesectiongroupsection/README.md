
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/useronenotesectiongroupsection` Documentation

The `useronenotesectiongroupsection` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/useronenotesectiongroupsection"
```


### Client Initialization

```go
client := useronenotesectiongroupsection.NewUserOnenoteSectionGroupSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOnenoteSectionGroupSectionClient.CreateUserByIdOnenoteSectionGroupByIdSection`

```go
ctx := context.TODO()
id := useronenotesectiongroupsection.NewUserOnenoteSectionGroupID("userIdValue", "sectionGroupIdValue")

payload := useronenotesectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.CreateUserByIdOnenoteSectionGroupByIdSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionGroupSectionClient.CreateUserByIdOnenoteSectionGroupByIdSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := useronenotesectiongroupsection.NewUserOnenoteSectionGroupSectionID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := useronenotesectiongroupsection.CreateUserByIdOnenoteSectionGroupByIdSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteSectionGroupByIdSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionGroupSectionClient.CreateUserByIdOnenoteSectionGroupByIdSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := useronenotesectiongroupsection.NewUserOnenoteSectionGroupSectionID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := useronenotesectiongroupsection.CreateUserByIdOnenoteSectionGroupByIdSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteSectionGroupByIdSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionGroupSectionClient.DeleteUserByIdOnenoteSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := useronenotesectiongroupsection.NewUserOnenoteSectionGroupSectionID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.DeleteUserByIdOnenoteSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionGroupSectionClient.GetUserByIdOnenoteSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := useronenotesectiongroupsection.NewUserOnenoteSectionGroupSectionID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetUserByIdOnenoteSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionGroupSectionClient.GetUserByIdOnenoteSectionGroupByIdSectionCount`

```go
ctx := context.TODO()
id := useronenotesectiongroupsection.NewUserOnenoteSectionGroupID("userIdValue", "sectionGroupIdValue")

read, err := client.GetUserByIdOnenoteSectionGroupByIdSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionGroupSectionClient.ListUserByIdOnenoteSectionGroupByIdSections`

```go
ctx := context.TODO()
id := useronenotesectiongroupsection.NewUserOnenoteSectionGroupID("userIdValue", "sectionGroupIdValue")

// alternatively `client.ListUserByIdOnenoteSectionGroupByIdSections(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOnenoteSectionGroupByIdSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserOnenoteSectionGroupSectionClient.UpdateUserByIdOnenoteSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := useronenotesectiongroupsection.NewUserOnenoteSectionGroupSectionID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := useronenotesectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.UpdateUserByIdOnenoteSectionGroupByIdSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
