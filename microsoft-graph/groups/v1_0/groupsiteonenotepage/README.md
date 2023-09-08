
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupsiteonenotepage` Documentation

The `groupsiteonenotepage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupsiteonenotepage"
```


### Client Initialization

```go
client := groupsiteonenotepage.NewGroupSiteOnenotePageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteOnenotePageClient.CreateGroupByIdSiteByIdOnenotePage`

```go
ctx := context.TODO()
id := groupsiteonenotepage.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteonenotepage.OnenotePage{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenotePage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenotePageClient.CreateGroupByIdSiteByIdOnenotePageByIdCopyToSection`

```go
ctx := context.TODO()
id := groupsiteonenotepage.NewGroupSiteOnenotePageID("groupIdValue", "siteIdValue", "onenotePageIdValue")

payload := groupsiteonenotepage.CreateGroupByIdSiteByIdOnenotePageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenotePageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenotePageClient.CreateGroupByIdSiteByIdOnenotePageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := groupsiteonenotepage.NewGroupSiteOnenotePageID("groupIdValue", "siteIdValue", "onenotePageIdValue")

payload := groupsiteonenotepage.CreateGroupByIdSiteByIdOnenotePageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenotePageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenotePageClient.DeleteGroupByIdSiteByIdOnenotePageById`

```go
ctx := context.TODO()
id := groupsiteonenotepage.NewGroupSiteOnenotePageID("groupIdValue", "siteIdValue", "onenotePageIdValue")

read, err := client.DeleteGroupByIdSiteByIdOnenotePageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenotePageClient.GetGroupByIdSiteByIdOnenotePageById`

```go
ctx := context.TODO()
id := groupsiteonenotepage.NewGroupSiteOnenotePageID("groupIdValue", "siteIdValue", "onenotePageIdValue")

read, err := client.GetGroupByIdSiteByIdOnenotePageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenotePageClient.GetGroupByIdSiteByIdOnenotePageCount`

```go
ctx := context.TODO()
id := groupsiteonenotepage.NewGroupSiteID("groupIdValue", "siteIdValue")

read, err := client.GetGroupByIdSiteByIdOnenotePageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenotePageClient.ListGroupByIdSiteByIdOnenotePages`

```go
ctx := context.TODO()
id := groupsiteonenotepage.NewGroupSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListGroupByIdSiteByIdOnenotePages(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdOnenotePagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteOnenotePageClient.UpdateGroupByIdSiteByIdOnenotePageById`

```go
ctx := context.TODO()
id := groupsiteonenotepage.NewGroupSiteOnenotePageID("groupIdValue", "siteIdValue", "onenotePageIdValue")

payload := groupsiteonenotepage.OnenotePage{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdOnenotePageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
