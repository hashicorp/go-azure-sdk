
## `github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/migrates` Documentation

The `migrates` SDK allows for interaction with the Azure Resource Manager Service `migrate` (API Version `2020-07-07`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/migrates"
```


### Client Initialization

```go
client := migrates.NewMigratesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MigratesClient.HyperVSitesList`

```go
ctx := context.TODO()
id := migrates.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.HyperVSitesList(ctx, id)` can be used to do batched pagination
items, err := client.HyperVSitesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MigratesClient.HyperVSitesListBySubscription`

```go
ctx := context.TODO()
id := migrates.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.HyperVSitesListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.HyperVSitesListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MigratesClient.VMwareSitesList`

```go
ctx := context.TODO()
id := migrates.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.VMwareSitesList(ctx, id)` can be used to do batched pagination
items, err := client.VMwareSitesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MigratesClient.VMwareSitesListBySubscription`

```go
ctx := context.TODO()
id := migrates.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.VMwareSitesListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.VMwareSitesListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
