
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupsitepermission` Documentation

The `groupsitepermission` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupsitepermission"
```


### Client Initialization

```go
client := groupsitepermission.NewGroupSitePermissionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSitePermissionClient.CreateGroupByIdSiteByIdPermission`

```go
ctx := context.TODO()
id := groupsitepermission.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsitepermission.Permission{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdPermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSitePermissionClient.DeleteGroupByIdSiteByIdPermissionById`

```go
ctx := context.TODO()
id := groupsitepermission.NewGroupSitePermissionID("groupIdValue", "siteIdValue", "permissionIdValue")

read, err := client.DeleteGroupByIdSiteByIdPermissionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSitePermissionClient.GetGroupByIdSiteByIdPermissionById`

```go
ctx := context.TODO()
id := groupsitepermission.NewGroupSitePermissionID("groupIdValue", "siteIdValue", "permissionIdValue")

read, err := client.GetGroupByIdSiteByIdPermissionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSitePermissionClient.GetGroupByIdSiteByIdPermissionCount`

```go
ctx := context.TODO()
id := groupsitepermission.NewGroupSiteID("groupIdValue", "siteIdValue")

read, err := client.GetGroupByIdSiteByIdPermissionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSitePermissionClient.ListGroupByIdSiteByIdPermissionByIdGrants`

```go
ctx := context.TODO()
id := groupsitepermission.NewGroupSitePermissionID("groupIdValue", "siteIdValue", "permissionIdValue")

// alternatively `client.ListGroupByIdSiteByIdPermissionByIdGrants(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdPermissionByIdGrantsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSitePermissionClient.ListGroupByIdSiteByIdPermissions`

```go
ctx := context.TODO()
id := groupsitepermission.NewGroupSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListGroupByIdSiteByIdPermissions(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdPermissionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSitePermissionClient.UpdateGroupByIdSiteByIdPermissionById`

```go
ctx := context.TODO()
id := groupsitepermission.NewGroupSitePermissionID("groupIdValue", "siteIdValue", "permissionIdValue")

payload := groupsitepermission.Permission{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdPermissionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
