
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/site` Documentation

The `site` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/site"
```


### Client Initialization

```go
client := site.NewSiteClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteClient.AddGroupSites`

```go
ctx := context.TODO()
id := site.NewGroupID("groupIdValue")

payload := site.AddGroupSitesRequest{
	// ...
}


// alternatively `client.AddGroupSites(ctx, id, payload)` can be used to do batched pagination
items, err := client.AddGroupSitesComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteClient.GetSite`

```go
ctx := context.TODO()
id := site.NewGroupIdSiteID("groupIdValue", "siteIdValue")

read, err := client.GetSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteClient.GetSiteCount`

```go
ctx := context.TODO()
id := site.NewGroupID("groupIdValue")

read, err := client.GetSiteCount(ctx, id)
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
id := site.NewGroupID("groupIdValue")

// alternatively `client.ListSites(ctx, id)` can be used to do batched pagination
items, err := client.ListSitesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteClient.RemoveGroupSites`

```go
ctx := context.TODO()
id := site.NewGroupID("groupIdValue")

payload := site.RemoveGroupSitesRequest{
	// ...
}


// alternatively `client.RemoveGroupSites(ctx, id, payload)` can be used to do batched pagination
items, err := client.RemoveGroupSitesComplete(ctx, id, payload)
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
id := site.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := site.Site{
	// ...
}


read, err := client.UpdateSite(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
