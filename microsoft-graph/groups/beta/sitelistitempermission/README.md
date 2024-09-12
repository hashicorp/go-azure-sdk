
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitempermission` Documentation

The `sitelistitempermission` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitempermission"
```


### Client Initialization

```go
client := sitelistitempermission.NewSiteListItemPermissionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteListItemPermissionClient.CreateSiteListItemPermission`

```go
ctx := context.TODO()
id := sitelistitempermission.NewGroupIdSiteIdListIdItemID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue")

payload := sitelistitempermission.Permission{
	// ...
}


read, err := client.CreateSiteListItemPermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListItemPermissionClient.DeleteSiteListItemPermission`

```go
ctx := context.TODO()
id := sitelistitempermission.NewGroupIdSiteIdListIdItemIdPermissionID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue", "permissionIdValue")

read, err := client.DeleteSiteListItemPermission(ctx, id, sitelistitempermission.DefaultDeleteSiteListItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListItemPermissionClient.GetSiteListItemPermission`

```go
ctx := context.TODO()
id := sitelistitempermission.NewGroupIdSiteIdListIdItemIdPermissionID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue", "permissionIdValue")

read, err := client.GetSiteListItemPermission(ctx, id, sitelistitempermission.DefaultGetSiteListItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListItemPermissionClient.GetSiteListItemPermissionsCount`

```go
ctx := context.TODO()
id := sitelistitempermission.NewGroupIdSiteIdListIdItemID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue")

read, err := client.GetSiteListItemPermissionsCount(ctx, id, sitelistitempermission.DefaultGetSiteListItemPermissionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListItemPermissionClient.ListSiteListItemPermissionGrants`

```go
ctx := context.TODO()
id := sitelistitempermission.NewGroupIdSiteIdListIdItemIdPermissionID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue", "permissionIdValue")

payload := sitelistitempermission.ListSiteListItemPermissionGrantsRequest{
	// ...
}


// alternatively `client.ListSiteListItemPermissionGrants(ctx, id, payload, sitelistitempermission.DefaultListSiteListItemPermissionGrantsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteListItemPermissionGrantsComplete(ctx, id, payload, sitelistitempermission.DefaultListSiteListItemPermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteListItemPermissionClient.ListSiteListItemPermissions`

```go
ctx := context.TODO()
id := sitelistitempermission.NewGroupIdSiteIdListIdItemID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue")

// alternatively `client.ListSiteListItemPermissions(ctx, id, sitelistitempermission.DefaultListSiteListItemPermissionsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteListItemPermissionsComplete(ctx, id, sitelistitempermission.DefaultListSiteListItemPermissionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteListItemPermissionClient.RevokeSiteListItemPermissionGrant`

```go
ctx := context.TODO()
id := sitelistitempermission.NewGroupIdSiteIdListIdItemIdPermissionID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue", "permissionIdValue")

payload := sitelistitempermission.RevokeSiteListItemPermissionGrantRequest{
	// ...
}


read, err := client.RevokeSiteListItemPermissionGrant(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListItemPermissionClient.UpdateSiteListItemPermission`

```go
ctx := context.TODO()
id := sitelistitempermission.NewGroupIdSiteIdListIdItemIdPermissionID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue", "permissionIdValue")

payload := sitelistitempermission.Permission{
	// ...
}


read, err := client.UpdateSiteListItemPermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
