
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitepermission` Documentation

The `sitepermission` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitepermission"
```


### Client Initialization

```go
client := sitepermission.NewSitePermissionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SitePermissionClient.CreateSitePermission`

```go
ctx := context.TODO()
id := sitepermission.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := sitepermission.Permission{
	// ...
}


read, err := client.CreateSitePermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitePermissionClient.DeleteSitePermission`

```go
ctx := context.TODO()
id := sitepermission.NewGroupIdSiteIdPermissionID("groupIdValue", "siteIdValue", "permissionIdValue")

read, err := client.DeleteSitePermission(ctx, id, sitepermission.DefaultDeleteSitePermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitePermissionClient.GetSitePermission`

```go
ctx := context.TODO()
id := sitepermission.NewGroupIdSiteIdPermissionID("groupIdValue", "siteIdValue", "permissionIdValue")

read, err := client.GetSitePermission(ctx, id, sitepermission.DefaultGetSitePermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitePermissionClient.GetSitePermissionsCount`

```go
ctx := context.TODO()
id := sitepermission.NewGroupIdSiteID("groupIdValue", "siteIdValue")

read, err := client.GetSitePermissionsCount(ctx, id, sitepermission.DefaultGetSitePermissionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitePermissionClient.ListSitePermissionGrants`

```go
ctx := context.TODO()
id := sitepermission.NewGroupIdSiteIdPermissionID("groupIdValue", "siteIdValue", "permissionIdValue")

payload := sitepermission.ListSitePermissionGrantsRequest{
	// ...
}


// alternatively `client.ListSitePermissionGrants(ctx, id, payload, sitepermission.DefaultListSitePermissionGrantsOperationOptions())` can be used to do batched pagination
items, err := client.ListSitePermissionGrantsComplete(ctx, id, payload, sitepermission.DefaultListSitePermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SitePermissionClient.ListSitePermissions`

```go
ctx := context.TODO()
id := sitepermission.NewGroupIdSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListSitePermissions(ctx, id, sitepermission.DefaultListSitePermissionsOperationOptions())` can be used to do batched pagination
items, err := client.ListSitePermissionsComplete(ctx, id, sitepermission.DefaultListSitePermissionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SitePermissionClient.UpdateSitePermission`

```go
ctx := context.TODO()
id := sitepermission.NewGroupIdSiteIdPermissionID("groupIdValue", "siteIdValue", "permissionIdValue")

payload := sitepermission.Permission{
	// ...
}


read, err := client.UpdateSitePermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
