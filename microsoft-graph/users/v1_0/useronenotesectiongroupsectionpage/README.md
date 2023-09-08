
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/useronenotesectiongroupsectionpage` Documentation

The `useronenotesectiongroupsectionpage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/useronenotesectiongroupsectionpage"
```


### Client Initialization

```go
client := useronenotesectiongroupsectionpage.NewUserOnenoteSectionGroupSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOnenoteSectionGroupSectionPageClient.CreateUserByIdOnenoteSectionGroupByIdSectionByIdPage`

```go
ctx := context.TODO()
id := useronenotesectiongroupsectionpage.NewUserOnenoteSectionGroupSectionID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := useronenotesectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.CreateUserByIdOnenoteSectionGroupByIdSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionGroupSectionPageClient.CreateUserByIdOnenoteSectionGroupByIdSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := useronenotesectiongroupsectionpage.NewUserOnenoteSectionGroupSectionPageID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := useronenotesectiongroupsectionpage.CreateUserByIdOnenoteSectionGroupByIdSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteSectionGroupByIdSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionGroupSectionPageClient.CreateUserByIdOnenoteSectionGroupByIdSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := useronenotesectiongroupsectionpage.NewUserOnenoteSectionGroupSectionPageID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := useronenotesectiongroupsectionpage.CreateUserByIdOnenoteSectionGroupByIdSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteSectionGroupByIdSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionGroupSectionPageClient.DeleteUserByIdOnenoteSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := useronenotesectiongroupsectionpage.NewUserOnenoteSectionGroupSectionPageID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteUserByIdOnenoteSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionGroupSectionPageClient.GetUserByIdOnenoteSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := useronenotesectiongroupsectionpage.NewUserOnenoteSectionGroupSectionPageID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetUserByIdOnenoteSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionGroupSectionPageClient.GetUserByIdOnenoteSectionGroupByIdSectionByIdPageCount`

```go
ctx := context.TODO()
id := useronenotesectiongroupsectionpage.NewUserOnenoteSectionGroupSectionID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetUserByIdOnenoteSectionGroupByIdSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionGroupSectionPageClient.ListUserByIdOnenoteSectionGroupByIdSectionByIdPages`

```go
ctx := context.TODO()
id := useronenotesectiongroupsectionpage.NewUserOnenoteSectionGroupSectionID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

// alternatively `client.ListUserByIdOnenoteSectionGroupByIdSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOnenoteSectionGroupByIdSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserOnenoteSectionGroupSectionPageClient.UpdateUserByIdOnenoteSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := useronenotesectiongroupsectionpage.NewUserOnenoteSectionGroupSectionPageID("userIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := useronenotesectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateUserByIdOnenoteSectionGroupByIdSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
