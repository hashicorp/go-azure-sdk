
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecontenttype` Documentation

The `sitecontenttype` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecontenttype"
```


### Client Initialization

```go
client := sitecontenttype.NewSiteContentTypeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteContentTypeClient.AddSiteContentTypesCopy`

```go
ctx := context.TODO()
id := sitecontenttype.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := sitecontenttype.AddSiteContentTypesCopyRequest{
	// ...
}


read, err := client.AddSiteContentTypesCopy(ctx, id, payload)
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
id := sitecontenttype.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := sitecontenttype.AddSiteContentTypesCopyFromHubRequest{
	// ...
}


read, err := client.AddSiteContentTypesCopyFromHub(ctx, id, payload)
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
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

payload := sitecontenttype.CopySiteContentTypeToDefaultContentLocationRequest{
	// ...
}


read, err := client.CopySiteContentTypeToDefaultContentLocation(ctx, id, payload)
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
id := sitecontenttype.NewGroupIdSiteIdContentTypeID("groupIdValue", "siteIdValue", "contentTypeIdValue")

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
id := sitecontenttype.NewGroupIdSiteID("groupIdValue", "siteIdValue")

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
id := sitecontenttype.NewGroupIdSiteID("groupIdValue", "siteIdValue")

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
