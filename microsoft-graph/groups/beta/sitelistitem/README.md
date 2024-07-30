
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitem` Documentation

The `sitelistitem` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitem"
```


### Client Initialization

```go
client := sitelistitem.NewSiteListItemClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteListItemClient.CreateSiteListItem`

```go
ctx := context.TODO()
id := sitelistitem.NewGroupIdSiteIdListID("groupIdValue", "siteIdValue", "listIdValue")

payload := sitelistitem.ListItem{
	// ...
}


read, err := client.CreateSiteListItem(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteListItemClient.CreateSiteListItemCreateLink`

```go
ctx := context.TODO()
id := sitelistitem.NewGroupIdSiteIdListIdItemID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue")

payload := sitelistitem.CreateSiteListItemCreateLinkRequest{
	// ...
}


read, err := client.CreateSiteListItemCreateLink(ctx, id, payload)
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
id := sitelistitem.NewGroupIdSiteIdListIdItemID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue")

read, err := client.DeleteSiteListItem(ctx, id)
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
id := sitelistitem.NewGroupIdSiteIdListIdItemID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue")

read, err := client.GetSiteListItem(ctx, id)
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
id := sitelistitem.NewGroupIdSiteIdListID("groupIdValue", "siteIdValue", "listIdValue")

// alternatively `client.ListSiteListItems(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteListItemsComplete(ctx, id)
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
id := sitelistitem.NewGroupIdSiteIdListIdItemID("groupIdValue", "siteIdValue", "listIdValue", "listItemIdValue")

payload := sitelistitem.ListItem{
	// ...
}


read, err := client.UpdateSiteListItem(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
