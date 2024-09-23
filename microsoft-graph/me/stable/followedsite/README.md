
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/followedsite` Documentation

The `followedsite` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/followedsite"
```


### Client Initialization

```go
client := followedsite.NewFollowedSiteClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FollowedSiteClient.AddFollowedSites`

```go
ctx := context.TODO()

payload := followedsite.AddFollowedSitesRequest{
	// ...
}


// alternatively `client.AddFollowedSites(ctx, payload, followedsite.DefaultAddFollowedSitesOperationOptions())` can be used to do batched pagination
items, err := client.AddFollowedSitesComplete(ctx, payload, followedsite.DefaultAddFollowedSitesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FollowedSiteClient.GetFollowedSite`

```go
ctx := context.TODO()
id := followedsite.NewMeFollowedSiteID("siteId")

read, err := client.GetFollowedSite(ctx, id, followedsite.DefaultGetFollowedSiteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FollowedSiteClient.GetFollowedSitesCount`

```go
ctx := context.TODO()


read, err := client.GetFollowedSitesCount(ctx, followedsite.DefaultGetFollowedSitesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FollowedSiteClient.ListFollowedSites`

```go
ctx := context.TODO()


// alternatively `client.ListFollowedSites(ctx, followedsite.DefaultListFollowedSitesOperationOptions())` can be used to do batched pagination
items, err := client.ListFollowedSitesComplete(ctx, followedsite.DefaultListFollowedSitesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FollowedSiteClient.RemoveFollowedSites`

```go
ctx := context.TODO()

payload := followedsite.RemoveFollowedSitesRequest{
	// ...
}


// alternatively `client.RemoveFollowedSites(ctx, payload, followedsite.DefaultRemoveFollowedSitesOperationOptions())` can be used to do batched pagination
items, err := client.RemoveFollowedSitesComplete(ctx, payload, followedsite.DefaultRemoveFollowedSitesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
