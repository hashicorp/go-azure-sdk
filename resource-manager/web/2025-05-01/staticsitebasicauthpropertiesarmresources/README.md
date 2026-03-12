
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitebasicauthpropertiesarmresources` Documentation

The `staticsitebasicauthpropertiesarmresources` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitebasicauthpropertiesarmresources"
```


### Client Initialization

```go
client := staticsitebasicauthpropertiesarmresources.NewStaticSiteBasicAuthPropertiesARMResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StaticSiteBasicAuthPropertiesARMResourcesClient.StaticSitesCreateOrUpdateBasicAuth`

```go
ctx := context.TODO()
id := staticsitebasicauthpropertiesarmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

payload := staticsitebasicauthpropertiesarmresources.StaticSiteBasicAuthPropertiesARMResource{
	// ...
}


read, err := client.StaticSitesCreateOrUpdateBasicAuth(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteBasicAuthPropertiesARMResourcesClient.StaticSitesGetBasicAuth`

```go
ctx := context.TODO()
id := staticsitebasicauthpropertiesarmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

read, err := client.StaticSitesGetBasicAuth(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteBasicAuthPropertiesARMResourcesClient.StaticSitesListBasicAuth`

```go
ctx := context.TODO()
id := staticsitebasicauthpropertiesarmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

// alternatively `client.StaticSitesListBasicAuth(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesListBasicAuthComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
