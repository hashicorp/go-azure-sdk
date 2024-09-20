
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcontenttype` Documentation

The `sitelistcontenttype` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcontenttype"
```


### Client Initialization

```go
client := sitelistcontenttype.NewSiteListContentTypeClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteListContentTypeClient.AddSiteListContentTypesCopy`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupId", "siteId", "listId")

payload := sitelistcontenttype.AddSiteListContentTypesCopyRequest{
	// ...
}


read, err := client.AddSiteListContentTypesCopy(ctx, id, payload, sitelistcontenttype.DefaultAddSiteListContentTypesCopyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.AddSiteListContentTypesCopyFromHub`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupId", "siteId", "listId")

payload := sitelistcontenttype.AddSiteListContentTypesCopyFromHubRequest{
	// ...
}


read, err := client.AddSiteListContentTypesCopyFromHub(ctx, id, payload, sitelistcontenttype.DefaultAddSiteListContentTypesCopyFromHubOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.CopySiteListContentTypeToDefaultContentLocation`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupId", "siteId", "listId", "contentTypeId")

payload := sitelistcontenttype.CopySiteListContentTypeToDefaultContentLocationRequest{
	// ...
}


read, err := client.CopySiteListContentTypeToDefaultContentLocation(ctx, id, payload, sitelistcontenttype.DefaultCopySiteListContentTypeToDefaultContentLocationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.CreateSiteListContentType`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupId", "siteId", "listId")

payload := sitelistcontenttype.ContentType{
	// ...
}


read, err := client.CreateSiteListContentType(ctx, id, payload, sitelistcontenttype.DefaultCreateSiteListContentTypeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.CreateSiteListContentTypeAssociateWithHubSite`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupId", "siteId", "listId", "contentTypeId")

payload := sitelistcontenttype.CreateSiteListContentTypeAssociateWithHubSiteRequest{
	// ...
}


read, err := client.CreateSiteListContentTypeAssociateWithHubSite(ctx, id, payload, sitelistcontenttype.DefaultCreateSiteListContentTypeAssociateWithHubSiteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.CreateSiteListContentTypePublish`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupId", "siteId", "listId", "contentTypeId")

read, err := client.CreateSiteListContentTypePublish(ctx, id, sitelistcontenttype.DefaultCreateSiteListContentTypePublishOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.CreateSiteListContentTypeUnpublish`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupId", "siteId", "listId", "contentTypeId")

read, err := client.CreateSiteListContentTypeUnpublish(ctx, id, sitelistcontenttype.DefaultCreateSiteListContentTypeUnpublishOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.DeleteSiteListContentType`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupId", "siteId", "listId", "contentTypeId")

read, err := client.DeleteSiteListContentType(ctx, id, sitelistcontenttype.DefaultDeleteSiteListContentTypeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.GetSiteListContentType`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupId", "siteId", "listId", "contentTypeId")

read, err := client.GetSiteListContentType(ctx, id, sitelistcontenttype.DefaultGetSiteListContentTypeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.GetSiteListContentTypesCount`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupId", "siteId", "listId")

read, err := client.GetSiteListContentTypesCount(ctx, id, sitelistcontenttype.DefaultGetSiteListContentTypesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.ListSiteListContentTypes`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupId", "siteId", "listId")

// alternatively `client.ListSiteListContentTypes(ctx, id, sitelistcontenttype.DefaultListSiteListContentTypesOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteListContentTypesComplete(ctx, id, sitelistcontenttype.DefaultListSiteListContentTypesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteListContentTypeClient.UpdateSiteListContentType`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupId", "siteId", "listId", "contentTypeId")

payload := sitelistcontenttype.ContentType{
	// ...
}


read, err := client.UpdateSiteListContentType(ctx, id, payload, sitelistcontenttype.DefaultUpdateSiteListContentTypeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
