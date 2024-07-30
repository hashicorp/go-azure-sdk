
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitepermission` Documentation

The `sitepermission` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitepermission"
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


### Example Usage: `SitePermissionClient.CreateSitePermissionRevokeGrant`

```go
ctx := context.TODO()
id := sitepermission.NewGroupIdSiteIdPermissionID("groupIdValue", "siteIdValue", "permissionIdValue")

payload := sitepermission.CreateSitePermissionRevokeGrantRequest{
	// ...
}


read, err := client.CreateSitePermissionRevokeGrant(ctx, id, payload)
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

read, err := client.DeleteSitePermission(ctx, id)
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

read, err := client.GetSitePermission(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitePermissionClient.GetSitePermissionCount`

```go
ctx := context.TODO()
id := sitepermission.NewGroupIdSiteID("groupIdValue", "siteIdValue")

read, err := client.GetSitePermissionCount(ctx, id)
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


// alternatively `client.ListSitePermissionGrants(ctx, id, payload)` can be used to do batched pagination
items, err := client.ListSitePermissionGrantsComplete(ctx, id, payload)
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

// alternatively `client.ListSitePermissions(ctx, id)` can be used to do batched pagination
items, err := client.ListSitePermissionsComplete(ctx, id)
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
