
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupsitecontenttype` Documentation

The `groupsitecontenttype` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupsitecontenttype"
```


### Client Initialization

```go
client := groupsitecontenttype.NewGroupSiteContentTypeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteContentTypeClient.AddGroupByIdSiteByIdContentTypesCopy`

```go
ctx := context.TODO()
id := groupsitecontenttype.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsitecontenttype.AddGroupByIdSiteByIdContentTypesCopyRequest{
	// ...
}


read, err := client.AddGroupByIdSiteByIdContentTypesCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteContentTypeClient.AddGroupByIdSiteByIdContentTypesCopyFromContentTypeHub`

```go
ctx := context.TODO()
id := groupsitecontenttype.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsitecontenttype.AddGroupByIdSiteByIdContentTypesCopyFromContentTypeHubRequest{
	// ...
}


read, err := client.AddGroupByIdSiteByIdContentTypesCopyFromContentTypeHub(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteContentTypeClient.CreateGroupByIdSiteByIdContentType`

```go
ctx := context.TODO()
id := groupsitecontenttype.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsitecontenttype.ContentType{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdContentType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteContentTypeClient.CreateGroupByIdSiteByIdContentTypeByIdAssociateWithHubSite`

```go
ctx := context.TODO()
id := groupsitecontenttype.NewGroupSiteContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

payload := groupsitecontenttype.CreateGroupByIdSiteByIdContentTypeByIdAssociateWithHubSiteRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdContentTypeByIdAssociateWithHubSite(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteContentTypeClient.CreateGroupByIdSiteByIdContentTypeByIdCopyToDefaultContentLocation`

```go
ctx := context.TODO()
id := groupsitecontenttype.NewGroupSiteContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

payload := groupsitecontenttype.CreateGroupByIdSiteByIdContentTypeByIdCopyToDefaultContentLocationRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdContentTypeByIdCopyToDefaultContentLocation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteContentTypeClient.CreateGroupByIdSiteByIdContentTypeByIdPublish`

```go
ctx := context.TODO()
id := groupsitecontenttype.NewGroupSiteContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

read, err := client.CreateGroupByIdSiteByIdContentTypeByIdPublish(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteContentTypeClient.CreateGroupByIdSiteByIdContentTypeByIdUnpublish`

```go
ctx := context.TODO()
id := groupsitecontenttype.NewGroupSiteContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

read, err := client.CreateGroupByIdSiteByIdContentTypeByIdUnpublish(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteContentTypeClient.DeleteGroupByIdSiteByIdContentTypeById`

```go
ctx := context.TODO()
id := groupsitecontenttype.NewGroupSiteContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

read, err := client.DeleteGroupByIdSiteByIdContentTypeById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteContentTypeClient.GetGroupByIdSiteByIdContentTypeById`

```go
ctx := context.TODO()
id := groupsitecontenttype.NewGroupSiteContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

read, err := client.GetGroupByIdSiteByIdContentTypeById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteContentTypeClient.GetGroupByIdSiteByIdContentTypeCount`

```go
ctx := context.TODO()
id := groupsitecontenttype.NewGroupSiteID("groupIdValue", "siteIdValue")

read, err := client.GetGroupByIdSiteByIdContentTypeCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteContentTypeClient.ListGroupByIdSiteByIdContentTypes`

```go
ctx := context.TODO()
id := groupsitecontenttype.NewGroupSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListGroupByIdSiteByIdContentTypes(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdContentTypesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteContentTypeClient.UpdateGroupByIdSiteByIdContentTypeById`

```go
ctx := context.TODO()
id := groupsitecontenttype.NewGroupSiteContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

payload := groupsitecontenttype.ContentType{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdContentTypeById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
