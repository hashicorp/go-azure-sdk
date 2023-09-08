
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsitelistitem` Documentation

The `groupsitelistitem` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsitelistitem"
```


### Client Initialization

```go
client := groupsitelistitem.NewGroupSiteListItemClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteListItemClient.CreateGroupByIdSiteByIdListByIdItem`

```go
ctx := context.TODO()
id := groupsitelistitem.NewGroupSiteListID("groupIdValue", "siteIdValue", "listIdValue")

payload := groupsitelistitem.ListItem{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdListByIdItem(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListItemClient.CreateGroupByIdSiteByIdListByIdItemByIdCreateLink`

```go
ctx := context.TODO()
id := groupsitelistitem.NewGroupSiteListItemID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue")

payload := groupsitelistitem.CreateGroupByIdSiteByIdListByIdItemByIdCreateLinkRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdListByIdItemByIdCreateLink(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListItemClient.DeleteGroupByIdSiteByIdListByIdItemById`

```go
ctx := context.TODO()
id := groupsitelistitem.NewGroupSiteListItemID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue")

read, err := client.DeleteGroupByIdSiteByIdListByIdItemById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListItemClient.GetGroupByIdSiteByIdListByIdItemById`

```go
ctx := context.TODO()
id := groupsitelistitem.NewGroupSiteListItemID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue")

read, err := client.GetGroupByIdSiteByIdListByIdItemById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteListItemClient.ListGroupByIdSiteByIdListByIdItems`

```go
ctx := context.TODO()
id := groupsitelistitem.NewGroupSiteListID("groupIdValue", "siteIdValue", "listIdValue")

// alternatively `client.ListGroupByIdSiteByIdListByIdItems(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdListByIdItemsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteListItemClient.UpdateGroupByIdSiteByIdListByIdItemById`

```go
ctx := context.TODO()
id := groupsitelistitem.NewGroupSiteListItemID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue")

payload := groupsitelistitem.ListItem{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdListByIdItemById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
