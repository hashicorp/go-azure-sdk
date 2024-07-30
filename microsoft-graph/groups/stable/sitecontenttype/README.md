
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecontenttype` Documentation

The `sitecontenttype` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecontenttype"
```


### Client Initialization

```go
client := sitecontenttype.NewSiteContentTypeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteContentTypeClient.AddGroupSiteContentTypesCopy`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := sitecontenttype.AddGroupSiteContentTypesCopyRequest{
	// ...
}


read, err := client.AddGroupSiteContentTypesCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.AddGroupSiteContentTypesCopyFromContentTypeHub`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := sitecontenttype.AddGroupSiteContentTypesCopyFromContentTypeHubRequest{
	// ...
}


read, err := client.AddGroupSiteContentTypesCopyFromContentTypeHub(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.CreateSiteContentType`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := sitecontenttype.ContentType{
	// ...
}


read, err := client.CreateSiteContentType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.CreateSiteContentTypeAssociateWithHubSite`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

payload := sitecontenttype.CreateSiteContentTypeAssociateWithHubSiteRequest{
	// ...
}


read, err := client.CreateSiteContentTypeAssociateWithHubSite(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.CreateSiteContentTypeCopyToDefaultContentLocation`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

payload := sitecontenttype.CreateSiteContentTypeCopyToDefaultContentLocationRequest{
	// ...
}


read, err := client.CreateSiteContentTypeCopyToDefaultContentLocation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.CreateSiteContentTypePublish`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

read, err := client.CreateSiteContentTypePublish(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.CreateSiteContentTypeUnpublish`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

read, err := client.CreateSiteContentTypeUnpublish(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.DeleteSiteContentType`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

read, err := client.DeleteSiteContentType(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.GetSiteContentType`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

read, err := client.GetSiteContentType(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.GetSiteContentTypeCount`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteID("groupIdValue", "siteIdValue")

read, err := client.GetSiteContentTypeCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.ListSiteContentTypes`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListSiteContentTypes(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteContentTypesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteContentTypeClient.UpdateSiteContentType`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

payload := sitecontenttype.ContentType{
	// ...
}


read, err := client.UpdateSiteContentType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
