
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/followedsite` Documentation

The `followedsite` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/followedsite"
```


### Client Initialization

```go
client := followedsite.NewFollowedSiteClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FollowedSiteClient.AddFollowedSites`

```go
ctx := context.TODO()
id := followedsite.NewUserID("userId")

payload := followedsite.AddFollowedSitesRequest{
	// ...
}


// alternatively `client.AddFollowedSites(ctx, id, payload, followedsite.DefaultAddFollowedSitesOperationOptions())` can be used to do batched pagination
items, err := client.AddFollowedSitesComplete(ctx, id, payload, followedsite.DefaultAddFollowedSitesOperationOptions())
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
id := followedsite.NewUserIdFollowedSiteID("userId", "siteId")

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
id := followedsite.NewUserID("userId")

read, err := client.GetFollowedSitesCount(ctx, id, followedsite.DefaultGetFollowedSitesCountOperationOptions())
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
id := followedsite.NewUserID("userId")

// alternatively `client.ListFollowedSites(ctx, id, followedsite.DefaultListFollowedSitesOperationOptions())` can be used to do batched pagination
items, err := client.ListFollowedSitesComplete(ctx, id, followedsite.DefaultListFollowedSitesOperationOptions())
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
id := followedsite.NewUserID("userId")

payload := followedsite.RemoveFollowedSitesRequest{
	// ...
}


// alternatively `client.RemoveFollowedSites(ctx, id, payload, followedsite.DefaultRemoveFollowedSitesOperationOptions())` can be used to do batched pagination
items, err := client.RemoveFollowedSitesComplete(ctx, id, payload, followedsite.DefaultRemoveFollowedSitesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
