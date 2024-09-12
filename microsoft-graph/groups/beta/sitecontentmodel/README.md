
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecontentmodel` Documentation

The `sitecontentmodel` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecontentmodel"
```


### Client Initialization

```go
client := sitecontentmodel.NewSiteContentModelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteContentModelClient.AddSiteContentModelToDrive`

```go
ctx := context.TODO()
id := sitecontentmodel.NewGroupIdSiteIdContentModelID("groupIdValue", "siteIdValue", "contentModelIdValue")

payload := sitecontentmodel.AddSiteContentModelToDriveRequest{
	// ...
}


read, err := client.AddSiteContentModelToDrive(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentModelClient.CreateSiteContentModel`

```go
ctx := context.TODO()
id := sitecontentmodel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := sitecontentmodel.ContentModel{
	// ...
}


read, err := client.CreateSiteContentModel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentModelClient.DeleteSiteContentModel`

```go
ctx := context.TODO()
id := sitecontentmodel.NewGroupIdSiteIdContentModelID("groupIdValue", "siteIdValue", "contentModelIdValue")

read, err := client.DeleteSiteContentModel(ctx, id, sitecontentmodel.DefaultDeleteSiteContentModelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentModelClient.GetSiteContentModel`

```go
ctx := context.TODO()
id := sitecontentmodel.NewGroupIdSiteIdContentModelID("groupIdValue", "siteIdValue", "contentModelIdValue")

read, err := client.GetSiteContentModel(ctx, id, sitecontentmodel.DefaultGetSiteContentModelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentModelClient.GetSiteContentModelsCount`

```go
ctx := context.TODO()
id := sitecontentmodel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

read, err := client.GetSiteContentModelsCount(ctx, id, sitecontentmodel.DefaultGetSiteContentModelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentModelClient.ListSiteContentModels`

```go
ctx := context.TODO()
id := sitecontentmodel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListSiteContentModels(ctx, id, sitecontentmodel.DefaultListSiteContentModelsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteContentModelsComplete(ctx, id, sitecontentmodel.DefaultListSiteContentModelsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteContentModelClient.RemoveSiteContentModelFromDrive`

```go
ctx := context.TODO()
id := sitecontentmodel.NewGroupIdSiteIdContentModelID("groupIdValue", "siteIdValue", "contentModelIdValue")

payload := sitecontentmodel.RemoveSiteContentModelFromDriveRequest{
	// ...
}


read, err := client.RemoveSiteContentModelFromDrive(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContentModelClient.UpdateSiteContentModel`

```go
ctx := context.TODO()
id := sitecontentmodel.NewGroupIdSiteIdContentModelID("groupIdValue", "siteIdValue", "contentModelIdValue")

payload := sitecontentmodel.ContentModel{
	// ...
}


read, err := client.UpdateSiteContentModel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
