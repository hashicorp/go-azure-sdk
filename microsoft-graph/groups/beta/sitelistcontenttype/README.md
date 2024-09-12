
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


### Example Usage: `SiteListContentTypeClient.AddSiteListContentTypesCopy`

```go
ctx := context.TODO()
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupIdValue", "siteIdValue", "listIdValue")

payload := sitelistcontenttype.AddSiteListContentTypesCopyRequest{
	// ...
}


read, err := client.AddSiteListContentTypesCopy(ctx, id, payload)
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
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupIdValue", "siteIdValue", "listIdValue")

payload := sitelistcontenttype.AddSiteListContentTypesCopyFromHubRequest{
	// ...
}


read, err := client.AddSiteListContentTypesCopyFromHub(ctx, id, payload)
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
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

payload := sitelistcontenttype.CopySiteListContentTypeToDefaultContentLocationRequest{
	// ...
}


read, err := client.CopySiteListContentTypeToDefaultContentLocation(ctx, id, payload)
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
id := sitelistcontenttype.NewGroupIdSiteIdListIdContentTypeID("groupIdValue", "siteIdValue", "listIdValue", "contentTypeIdValue")

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
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupIdValue", "siteIdValue", "listIdValue")

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
id := sitelistcontenttype.NewGroupIdSiteIdListID("groupIdValue", "siteIdValue", "listIdValue")

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
