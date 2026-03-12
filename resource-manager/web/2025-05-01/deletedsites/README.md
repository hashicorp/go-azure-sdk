
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/deletedsites` Documentation

The `deletedsites` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/deletedsites"
```


### Client Initialization

```go
client := deletedsites.NewDeletedSitesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeletedSitesClient.DeletedWebAppsGetDeletedWebAppByLocation`

```go
ctx := context.TODO()
id := deletedsites.NewLocationDeletedSiteID("12345678-1234-9876-4563-123456789012", "locationName", "deletedSiteId")

read, err := client.DeletedWebAppsGetDeletedWebAppByLocation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedSitesClient.DeletedWebAppsListByLocation`

```go
ctx := context.TODO()
id := deletedsites.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.DeletedWebAppsListByLocation(ctx, id)` can be used to do batched pagination
items, err := client.DeletedWebAppsListByLocationComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
