
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcontenttype` Documentation

The `sitelistcontenttype` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcontenttype"
```


### Client Initialization

```go
client := sitelistcontenttype.NewSiteListContentTypeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteListContentTypeClient.AddGroupSiteListContentTypesCopy`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupIdValue", "siteIdValue", "listIdValue")

payload := sitelistcontenttype.AddGroupSiteListContentTypesCopyRequest{
	// ...
}


read, err := client.AddGroupSiteListContentTypesCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.AddGroupSiteListContentTypesCopyFromContentTypeHub`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupIdValue", "siteIdValue", "listIdValue")

payload := sitelistcontenttype.AddGroupSiteListContentTypesCopyFromContentTypeHubRequest{
	// ...
}


read, err := client.AddGroupSiteListContentTypesCopyFromContentTypeHub(ctx, id, payload)
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
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupIdValue", "siteIdValue", "listIdValue")

payload := sitelistcontenttype.ContentType{
	// ...
}


read, err := client.CreateSiteListContentType(ctx, id, payload)
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
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

payload := sitelistcontenttype.CreateSiteListContentTypeAssociateWithHubSiteRequest{
	// ...
}


read, err := client.CreateSiteListContentTypeAssociateWithHubSite(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.CreateSiteListContentTypeCopyToDefaultContentLocation`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

payload := sitelistcontenttype.CreateSiteListContentTypeCopyToDefaultContentLocationRequest{
	// ...
}


read, err := client.CreateSiteListContentTypeCopyToDefaultContentLocation(ctx, id, payload)
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
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

read, err := client.CreateSiteListContentTypePublish(ctx, id)
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
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

read, err := client.CreateSiteListContentTypeUnpublish(ctx, id)
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
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

read, err := client.DeleteSiteListContentType(ctx, id)
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
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

read, err := client.GetSiteListContentType(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListContentTypeClient.GetSiteListContentTypeCount`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupIdValue", "siteIdValue", "listIdValue")

read, err := client.GetSiteListContentTypeCount(ctx, id)
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
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupIdValue", "siteIdValue", "listIdValue")

// alternatively `client.ListSiteListContentTypes(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteListContentTypesComplete(ctx, id)
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
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

payload := sitelistcontenttype.ContentType{
	// ...
}


read, err := client.UpdateSiteListContentType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
