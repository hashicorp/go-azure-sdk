
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebinitem` Documentation

The `siterecyclebinitem` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebinitem"
```


### Client Initialization

```go
client := siterecyclebinitem.NewSiteRecycleBinItemClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteRecycleBinItemClient.CreateSiteRecycleBinItem`

```go
ctx := context.TODO()
id := siterecyclebinitem.NewGroupIdSiteID("groupId", "siteId")

payload := siterecyclebinitem.RecycleBinItem{
	// ...
}


read, err := client.CreateSiteRecycleBinItem(ctx, id, payload, siterecyclebinitem.DefaultCreateSiteRecycleBinItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteRecycleBinItemClient.DeleteSiteRecycleBinItem`

```go
ctx := context.TODO()
id := siterecyclebinitem.NewGroupIdSiteIdRecycleBinItemID("groupId", "siteId", "recycleBinItemId")

read, err := client.DeleteSiteRecycleBinItem(ctx, id, siterecyclebinitem.DefaultDeleteSiteRecycleBinItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteRecycleBinItemClient.DeleteSiteRecycleBinItems`

```go
ctx := context.TODO()
id := siterecyclebinitem.NewGroupIdSiteID("groupId", "siteId")

payload := siterecyclebinitem.DeleteSiteRecycleBinItemsRequest{
	// ...
}


read, err := client.DeleteSiteRecycleBinItems(ctx, id, payload, siterecyclebinitem.DefaultDeleteSiteRecycleBinItemsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteRecycleBinItemClient.GetSiteRecycleBinItem`

```go
ctx := context.TODO()
id := siterecyclebinitem.NewGroupIdSiteIdRecycleBinItemID("groupId", "siteId", "recycleBinItemId")

read, err := client.GetSiteRecycleBinItem(ctx, id, siterecyclebinitem.DefaultGetSiteRecycleBinItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteRecycleBinItemClient.GetSiteRecycleBinItemsCount`

```go
ctx := context.TODO()
id := siterecyclebinitem.NewGroupIdSiteID("groupId", "siteId")

read, err := client.GetSiteRecycleBinItemsCount(ctx, id, siterecyclebinitem.DefaultGetSiteRecycleBinItemsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteRecycleBinItemClient.ListSiteRecycleBinItems`

```go
ctx := context.TODO()
id := siterecyclebinitem.NewGroupIdSiteID("groupId", "siteId")

// alternatively `client.ListSiteRecycleBinItems(ctx, id, siterecyclebinitem.DefaultListSiteRecycleBinItemsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteRecycleBinItemsComplete(ctx, id, siterecyclebinitem.DefaultListSiteRecycleBinItemsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteRecycleBinItemClient.RestoreSiteRecycleBinItems`

```go
ctx := context.TODO()
id := siterecyclebinitem.NewGroupIdSiteID("groupId", "siteId")

payload := siterecyclebinitem.RestoreSiteRecycleBinItemsRequest{
	// ...
}


// alternatively `client.RestoreSiteRecycleBinItems(ctx, id, payload, siterecyclebinitem.DefaultRestoreSiteRecycleBinItemsOperationOptions())` can be used to do batched pagination
items, err := client.RestoreSiteRecycleBinItemsComplete(ctx, id, payload, siterecyclebinitem.DefaultRestoreSiteRecycleBinItemsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteRecycleBinItemClient.UpdateSiteRecycleBinItem`

```go
ctx := context.TODO()
id := siterecyclebinitem.NewGroupIdSiteIdRecycleBinItemID("groupId", "siteId", "recycleBinItemId")

payload := siterecyclebinitem.RecycleBinItem{
	// ...
}


read, err := client.UpdateSiteRecycleBinItem(ctx, id, payload, siterecyclebinitem.DefaultUpdateSiteRecycleBinItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
