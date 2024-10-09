
## `github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/favoritesapis` Documentation

The `favoritesapis` SDK allows for interaction with Azure Resource Manager `applicationinsights` (API Version `2015-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/favoritesapis"
```


### Client Initialization

```go
client := favoritesapis.NewFavoritesAPIsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FavoritesAPIsClient.FavoritesAdd`

```go
ctx := context.TODO()
id := favoritesapis.NewFavoriteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "favoriteId")

payload := favoritesapis.ApplicationInsightsComponentFavorite{
	// ...
}


read, err := client.FavoritesAdd(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FavoritesAPIsClient.FavoritesDelete`

```go
ctx := context.TODO()
id := favoritesapis.NewFavoriteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "favoriteId")

read, err := client.FavoritesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FavoritesAPIsClient.FavoritesGet`

```go
ctx := context.TODO()
id := favoritesapis.NewFavoriteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "favoriteId")

read, err := client.FavoritesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FavoritesAPIsClient.FavoritesList`

```go
ctx := context.TODO()
id := favoritesapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

read, err := client.FavoritesList(ctx, id, favoritesapis.DefaultFavoritesListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FavoritesAPIsClient.FavoritesUpdate`

```go
ctx := context.TODO()
id := favoritesapis.NewFavoriteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "favoriteId")

payload := favoritesapis.ApplicationInsightsComponentFavorite{
	// ...
}


read, err := client.FavoritesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
