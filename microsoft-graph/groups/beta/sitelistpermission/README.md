
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistpermission` Documentation

The `sitelistpermission` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistpermission"
```


### Client Initialization

```go
client := sitelistpermission.NewSiteListPermissionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteListPermissionClient.CreateSiteListPermission`

```go
ctx := context.TODO()
id := sitelistpermission.NewGroupIdSiteIdListID("groupId", "siteId", "listId")

payload := sitelistpermission.Permission{
	// ...
}


read, err := client.CreateSiteListPermission(ctx, id, payload, sitelistpermission.DefaultCreateSiteListPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListPermissionClient.DeleteSiteListPermission`

```go
ctx := context.TODO()
id := sitelistpermission.NewGroupIdSiteIdListIdPermissionID("groupId", "siteId", "listId", "permissionId")

read, err := client.DeleteSiteListPermission(ctx, id, sitelistpermission.DefaultDeleteSiteListPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListPermissionClient.GetSiteListPermission`

```go
ctx := context.TODO()
id := sitelistpermission.NewGroupIdSiteIdListIdPermissionID("groupId", "siteId", "listId", "permissionId")

read, err := client.GetSiteListPermission(ctx, id, sitelistpermission.DefaultGetSiteListPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListPermissionClient.GetSiteListPermissionsCount`

```go
ctx := context.TODO()
id := sitelistpermission.NewGroupIdSiteIdListID("groupId", "siteId", "listId")

read, err := client.GetSiteListPermissionsCount(ctx, id, sitelistpermission.DefaultGetSiteListPermissionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListPermissionClient.ListSiteListPermissionGrants`

```go
ctx := context.TODO()
id := sitelistpermission.NewGroupIdSiteIdListIdPermissionID("groupId", "siteId", "listId", "permissionId")

payload := sitelistpermission.ListSiteListPermissionGrantsRequest{
	// ...
}


// alternatively `client.ListSiteListPermissionGrants(ctx, id, payload, sitelistpermission.DefaultListSiteListPermissionGrantsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteListPermissionGrantsComplete(ctx, id, payload, sitelistpermission.DefaultListSiteListPermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteListPermissionClient.ListSiteListPermissions`

```go
ctx := context.TODO()
id := sitelistpermission.NewGroupIdSiteIdListID("groupId", "siteId", "listId")

// alternatively `client.ListSiteListPermissions(ctx, id, sitelistpermission.DefaultListSiteListPermissionsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteListPermissionsComplete(ctx, id, sitelistpermission.DefaultListSiteListPermissionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteListPermissionClient.RevokeSiteListPermissionGrants`

```go
ctx := context.TODO()
id := sitelistpermission.NewGroupIdSiteIdListIdPermissionID("groupId", "siteId", "listId", "permissionId")

payload := sitelistpermission.RevokeSiteListPermissionGrantsRequest{
	// ...
}


read, err := client.RevokeSiteListPermissionGrants(ctx, id, payload, sitelistpermission.DefaultRevokeSiteListPermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListPermissionClient.UpdateSiteListPermission`

```go
ctx := context.TODO()
id := sitelistpermission.NewGroupIdSiteIdListIdPermissionID("groupId", "siteId", "listId", "permissionId")

payload := sitelistpermission.Permission{
	// ...
}


read, err := client.UpdateSiteListPermission(ctx, id, payload, sitelistpermission.DefaultUpdateSiteListPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
