
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/site` Documentation

The `site` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/site"
```


### Client Initialization

```go
client := site.NewSiteClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteClient.AddSites`

```go
ctx := context.TODO()
id := site.NewGroupID("groupId")

payload := site.AddSitesRequest{
	// ...
}


// alternatively `client.AddSites(ctx, id, payload, site.DefaultAddSitesOperationOptions())` can be used to do batched pagination
items, err := client.AddSitesComplete(ctx, id, payload, site.DefaultAddSitesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteClient.CreateSiteArchive`

```go
ctx := context.TODO()
id := site.NewGroupIdSiteID("groupId", "siteId")

read, err := client.CreateSiteArchive(ctx, id, site.DefaultCreateSiteArchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteClient.CreateSiteUnarchive`

```go
ctx := context.TODO()
id := site.NewGroupIdSiteID("groupId", "siteId")

read, err := client.CreateSiteUnarchive(ctx, id, site.DefaultCreateSiteUnarchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteClient.GetSite`

```go
ctx := context.TODO()
id := site.NewGroupIdSiteID("groupId", "siteId")

read, err := client.GetSite(ctx, id, site.DefaultGetSiteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteClient.GetSitesCount`

```go
ctx := context.TODO()
id := site.NewGroupID("groupId")

read, err := client.GetSitesCount(ctx, id, site.DefaultGetSitesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteClient.ListSites`

```go
ctx := context.TODO()
id := site.NewGroupID("groupId")

// alternatively `client.ListSites(ctx, id, site.DefaultListSitesOperationOptions())` can be used to do batched pagination
items, err := client.ListSitesComplete(ctx, id, site.DefaultListSitesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteClient.RemoveSites`

```go
ctx := context.TODO()
id := site.NewGroupID("groupId")

payload := site.RemoveSitesRequest{
	// ...
}


// alternatively `client.RemoveSites(ctx, id, payload, site.DefaultRemoveSitesOperationOptions())` can be used to do batched pagination
items, err := client.RemoveSitesComplete(ctx, id, payload, site.DefaultRemoveSitesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteClient.UpdateSite`

```go
ctx := context.TODO()
id := site.NewGroupIdSiteID("groupId", "siteId")

payload := site.Site{
	// ...
}


read, err := client.UpdateSite(ctx, id, payload, site.DefaultUpdateSiteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
