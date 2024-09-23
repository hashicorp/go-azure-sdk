
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitem` Documentation

The `sitelistitem` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitem"
```


### Client Initialization

```go
client := sitelistitem.NewSiteListItemClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteListItemClient.CreateSiteListItem`

```go
ctx := context.TODO()
id := sitelistitem.NewGroupIdSiteIdListID("groupId", "siteId", "listId")

payload := sitelistitem.ListItem{
	// ...
}


read, err := client.CreateSiteListItem(ctx, id, payload, sitelistitem.DefaultCreateSiteListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListItemClient.CreateSiteListItemLink`

```go
ctx := context.TODO()
id := sitelistitem.NewGroupIdSiteIdListIdItemID("groupId", "siteId", "listId", "listItemId")

payload := sitelistitem.CreateSiteListItemLinkRequest{
	// ...
}


read, err := client.CreateSiteListItemLink(ctx, id, payload, sitelistitem.DefaultCreateSiteListItemLinkOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListItemClient.DeleteSiteListItem`

```go
ctx := context.TODO()
id := sitelistitem.NewGroupIdSiteIdListIdItemID("groupId", "siteId", "listId", "listItemId")

read, err := client.DeleteSiteListItem(ctx, id, sitelistitem.DefaultDeleteSiteListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListItemClient.GetSiteListItem`

```go
ctx := context.TODO()
id := sitelistitem.NewGroupIdSiteIdListIdItemID("groupId", "siteId", "listId", "listItemId")

read, err := client.GetSiteListItem(ctx, id, sitelistitem.DefaultGetSiteListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListItemClient.ListSiteListItems`

```go
ctx := context.TODO()
id := sitelistitem.NewGroupIdSiteIdListID("groupId", "siteId", "listId")

// alternatively `client.ListSiteListItems(ctx, id, sitelistitem.DefaultListSiteListItemsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteListItemsComplete(ctx, id, sitelistitem.DefaultListSiteListItemsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteListItemClient.UpdateSiteListItem`

```go
ctx := context.TODO()
id := sitelistitem.NewGroupIdSiteIdListIdItemID("groupId", "siteId", "listId", "listItemId")

payload := sitelistitem.ListItem{
	// ...
}


read, err := client.UpdateSiteListItem(ctx, id, payload, sitelistitem.DefaultUpdateSiteListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
