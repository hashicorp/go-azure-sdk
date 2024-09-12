
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/drivelistcontenttype` Documentation

The `drivelistcontenttype` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/drivelistcontenttype"
```


### Client Initialization

```go
client := drivelistcontenttype.NewDriveListContentTypeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveListContentTypeClient.AddDriveListContentTypesCopy`

```go
ctx := context.TODO()
id := drivelistcontenttype.NewMeDriveID("driveIdValue")

payload := drivelistcontenttype.AddDriveListContentTypesCopyRequest{
	// ...
}


read, err := client.AddDriveListContentTypesCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListContentTypeClient.AddDriveListContentTypesCopyFromHub`

```go
ctx := context.TODO()
id := drivelistcontenttype.NewMeDriveID("driveIdValue")

payload := drivelistcontenttype.AddDriveListContentTypesCopyFromHubRequest{
	// ...
}


read, err := client.AddDriveListContentTypesCopyFromHub(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListContentTypeClient.CopyDriveListContentTypeToDefaultContentLocation`

```go
ctx := context.TODO()
id := drivelistcontenttype.NewMeDriveIdListContentTypeID("driveIdValue", "contentTypeIdValue")

payload := drivelistcontenttype.CopyDriveListContentTypeToDefaultContentLocationRequest{
	// ...
}


read, err := client.CopyDriveListContentTypeToDefaultContentLocation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListContentTypeClient.CreateDriveListContentType`

```go
ctx := context.TODO()
id := drivelistcontenttype.NewMeDriveID("driveIdValue")

payload := drivelistcontenttype.ContentType{
	// ...
}


read, err := client.CreateDriveListContentType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListContentTypeClient.CreateDriveListContentTypeAssociateWithHubSite`

```go
ctx := context.TODO()
id := drivelistcontenttype.NewMeDriveIdListContentTypeID("driveIdValue", "contentTypeIdValue")

payload := drivelistcontenttype.CreateDriveListContentTypeAssociateWithHubSiteRequest{
	// ...
}


read, err := client.CreateDriveListContentTypeAssociateWithHubSite(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListContentTypeClient.CreateDriveListContentTypePublish`

```go
ctx := context.TODO()
id := drivelistcontenttype.NewMeDriveIdListContentTypeID("driveIdValue", "contentTypeIdValue")

read, err := client.CreateDriveListContentTypePublish(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListContentTypeClient.CreateDriveListContentTypeUnpublish`

```go
ctx := context.TODO()
id := drivelistcontenttype.NewMeDriveIdListContentTypeID("driveIdValue", "contentTypeIdValue")

read, err := client.CreateDriveListContentTypeUnpublish(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListContentTypeClient.DeleteDriveListContentType`

```go
ctx := context.TODO()
id := drivelistcontenttype.NewMeDriveIdListContentTypeID("driveIdValue", "contentTypeIdValue")

read, err := client.DeleteDriveListContentType(ctx, id, drivelistcontenttype.DefaultDeleteDriveListContentTypeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListContentTypeClient.GetDriveListContentType`

```go
ctx := context.TODO()
id := drivelistcontenttype.NewMeDriveIdListContentTypeID("driveIdValue", "contentTypeIdValue")

read, err := client.GetDriveListContentType(ctx, id, drivelistcontenttype.DefaultGetDriveListContentTypeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListContentTypeClient.GetDriveListContentTypesCount`

```go
ctx := context.TODO()
id := drivelistcontenttype.NewMeDriveID("driveIdValue")

read, err := client.GetDriveListContentTypesCount(ctx, id, drivelistcontenttype.DefaultGetDriveListContentTypesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListContentTypeClient.ListDriveListContentTypes`

```go
ctx := context.TODO()
id := drivelistcontenttype.NewMeDriveID("driveIdValue")

// alternatively `client.ListDriveListContentTypes(ctx, id, drivelistcontenttype.DefaultListDriveListContentTypesOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveListContentTypesComplete(ctx, id, drivelistcontenttype.DefaultListDriveListContentTypesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveListContentTypeClient.UpdateDriveListContentType`

```go
ctx := context.TODO()
id := drivelistcontenttype.NewMeDriveIdListContentTypeID("driveIdValue", "contentTypeIdValue")

payload := drivelistcontenttype.ContentType{
	// ...
}


read, err := client.UpdateDriveListContentType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
