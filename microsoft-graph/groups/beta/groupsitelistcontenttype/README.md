
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsitelistcontenttype` Documentation

The `groupsitelistcontenttype` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsitelistcontenttype"
```


### Client Initialization

```go
client := groupsitelistcontenttype.NewGroupSiteListContentTypeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteListContentTypeClient.AddGroupByIdSiteByIdListByIdContentTypesCopy`

```go
ctx := context.TODO()
id := groupsitelistcontenttype.NewGroupSiteListID("groupIdValue", "siteIdValue", "listIdValue")

payload := groupsitelistcontenttype.AddGroupByIdSiteByIdListByIdContentTypesCopyRequest{
	// ...
}


read, err := client.AddGroupByIdSiteByIdListByIdContentTypesCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListContentTypeClient.AddGroupByIdSiteByIdListByIdContentTypesCopyFromContentTypeHub`

```go
ctx := context.TODO()
id := groupsitelistcontenttype.NewGroupSiteListID("groupIdValue", "siteIdValue", "listIdValue")

payload := groupsitelistcontenttype.AddGroupByIdSiteByIdListByIdContentTypesCopyFromContentTypeHubRequest{
	// ...
}


read, err := client.AddGroupByIdSiteByIdListByIdContentTypesCopyFromContentTypeHub(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListContentTypeClient.CreateGroupByIdSiteByIdListByIdContentType`

```go
ctx := context.TODO()
id := groupsitelistcontenttype.NewGroupSiteListID("groupIdValue", "siteIdValue", "listIdValue")

payload := groupsitelistcontenttype.ContentType{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdListByIdContentType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListContentTypeClient.CreateGroupByIdSiteByIdListByIdContentTypeByIdAssociateWithHubSite`

```go
ctx := context.TODO()
id := groupsitelistcontenttype.NewGroupSiteListContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

payload := groupsitelistcontenttype.CreateGroupByIdSiteByIdListByIdContentTypeByIdAssociateWithHubSiteRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdListByIdContentTypeByIdAssociateWithHubSite(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListContentTypeClient.CreateGroupByIdSiteByIdListByIdContentTypeByIdCopyToDefaultContentLocation`

```go
ctx := context.TODO()
id := groupsitelistcontenttype.NewGroupSiteListContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

payload := groupsitelistcontenttype.CreateGroupByIdSiteByIdListByIdContentTypeByIdCopyToDefaultContentLocationRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdListByIdContentTypeByIdCopyToDefaultContentLocation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListContentTypeClient.CreateGroupByIdSiteByIdListByIdContentTypeByIdPublish`

```go
ctx := context.TODO()
id := groupsitelistcontenttype.NewGroupSiteListContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

read, err := client.CreateGroupByIdSiteByIdListByIdContentTypeByIdPublish(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListContentTypeClient.CreateGroupByIdSiteByIdListByIdContentTypeByIdUnpublish`

```go
ctx := context.TODO()
id := groupsitelistcontenttype.NewGroupSiteListContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

read, err := client.CreateGroupByIdSiteByIdListByIdContentTypeByIdUnpublish(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListContentTypeClient.DeleteGroupByIdSiteByIdListByIdContentTypeById`

```go
ctx := context.TODO()
id := groupsitelistcontenttype.NewGroupSiteListContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

read, err := client.DeleteGroupByIdSiteByIdListByIdContentTypeById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListContentTypeClient.GetGroupByIdSiteByIdListByIdContentTypeById`

```go
ctx := context.TODO()
id := groupsitelistcontenttype.NewGroupSiteListContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

read, err := client.GetGroupByIdSiteByIdListByIdContentTypeById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListContentTypeClient.GetGroupByIdSiteByIdListByIdContentTypeCount`

```go
ctx := context.TODO()
id := groupsitelistcontenttype.NewGroupSiteListID("groupIdValue", "siteIdValue", "listIdValue")

read, err := client.GetGroupByIdSiteByIdListByIdContentTypeCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListContentTypeClient.ListGroupByIdSiteByIdListByIdContentTypes`

```go
ctx := context.TODO()
id := groupsitelistcontenttype.NewGroupSiteListID("groupIdValue", "siteIdValue", "listIdValue")

// alternatively `client.ListGroupByIdSiteByIdListByIdContentTypes(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdListByIdContentTypesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteListContentTypeClient.UpdateGroupByIdSiteByIdListByIdContentTypeById`

```go
ctx := context.TODO()
id := groupsitelistcontenttype.NewGroupSiteListContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

payload := groupsitelistcontenttype.ContentType{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdListByIdContentTypeById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
