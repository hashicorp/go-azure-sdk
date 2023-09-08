
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupsite` Documentation

The `groupsite` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupsite"
```


### Client Initialization

```go
client := groupsite.NewGroupSiteClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteClient.AddGroupByIdSites`

```go
ctx := context.TODO()
id := groupsite.NewGroupID("groupIdValue")

// alternatively `client.AddGroupByIdSites(ctx, id)` can be used to do batched pagination
items, err := client.AddGroupByIdSitesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteClient.GetGroupByIdSiteById`

```go
ctx := context.TODO()
id := groupsite.NewGroupSiteID("groupIdValue", "siteIdValue")

read, err := client.GetGroupByIdSiteById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteClient.GetGroupByIdSiteCount`

```go
ctx := context.TODO()
id := groupsite.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdSiteCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteClient.ListGroupByIdSites`

```go
ctx := context.TODO()
id := groupsite.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdSites(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSitesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteClient.RemoveGroupByIdSites`

```go
ctx := context.TODO()
id := groupsite.NewGroupID("groupIdValue")

// alternatively `client.RemoveGroupByIdSites(ctx, id)` can be used to do batched pagination
items, err := client.RemoveGroupByIdSitesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteClient.UpdateGroupByIdSiteById`

```go
ctx := context.TODO()
id := groupsite.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsite.Site{
	// ...
}


read, err := client.UpdateGroupByIdSiteById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
