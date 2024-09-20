
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecontenttype` Documentation

The `sitecontenttype` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecontenttype"
```


### Client Initialization

```go
client := sitecontenttype.NewSiteContentTypeClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteContentTypeClient.AddSiteContentTypesCopy`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteID("groupId", "siteId")

payload := sitecontenttype.AddSiteContentTypesCopyRequest{
	// ...
}


read, err := client.AddSiteContentTypesCopy(ctx, id, payload, sitecontenttype.DefaultAddSiteContentTypesCopyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.AddSiteContentTypesCopyFromHub`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteID("groupId", "siteId")

payload := sitecontenttype.AddSiteContentTypesCopyFromHubRequest{
	// ...
}


read, err := client.AddSiteContentTypesCopyFromHub(ctx, id, payload, sitecontenttype.DefaultAddSiteContentTypesCopyFromHubOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.CopySiteContentTypeToDefaultContentLocation`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupId", "siteId", "contentTypeId")

payload := sitecontenttype.CopySiteContentTypeToDefaultContentLocationRequest{
	// ...
}


read, err := client.CopySiteContentTypeToDefaultContentLocation(ctx, id, payload, sitecontenttype.DefaultCopySiteContentTypeToDefaultContentLocationOperationOptions())
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
id := sitecontenttype.NewGroupIdSiteID("groupId", "siteId")

payload := sitecontenttype.ContentType{
	// ...
}


read, err := client.CreateSiteContentType(ctx, id, payload, sitecontenttype.DefaultCreateSiteContentTypeOperationOptions())
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
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupId", "siteId", "contentTypeId")

payload := sitecontenttype.CreateSiteContentTypeAssociateWithHubSiteRequest{
	// ...
}


read, err := client.CreateSiteContentTypeAssociateWithHubSite(ctx, id, payload, sitecontenttype.DefaultCreateSiteContentTypeAssociateWithHubSiteOperationOptions())
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
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupId", "siteId", "contentTypeId")

read, err := client.CreateSiteContentTypePublish(ctx, id, sitecontenttype.DefaultCreateSiteContentTypePublishOperationOptions())
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
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupId", "siteId", "contentTypeId")

read, err := client.CreateSiteContentTypeUnpublish(ctx, id, sitecontenttype.DefaultCreateSiteContentTypeUnpublishOperationOptions())
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
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupId", "siteId", "contentTypeId")

read, err := client.DeleteSiteContentType(ctx, id, sitecontenttype.DefaultDeleteSiteContentTypeOperationOptions())
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
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupId", "siteId", "contentTypeId")

read, err := client.GetSiteContentType(ctx, id, sitecontenttype.DefaultGetSiteContentTypeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentTypeClient.GetSiteContentTypesCount`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteID("groupId", "siteId")

read, err := client.GetSiteContentTypesCount(ctx, id, sitecontenttype.DefaultGetSiteContentTypesCountOperationOptions())
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
id := sitecontenttype.NewGroupIdSiteID("groupId", "siteId")

// alternatively `client.ListSiteContentTypes(ctx, id, sitecontenttype.DefaultListSiteContentTypesOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteContentTypesComplete(ctx, id, sitecontenttype.DefaultListSiteContentTypesOperationOptions())
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
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupId", "siteId", "contentTypeId")

payload := sitecontenttype.ContentType{
	// ...
}


read, err := client.UpdateSiteContentType(ctx, id, payload, sitecontenttype.DefaultUpdateSiteContentTypeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
