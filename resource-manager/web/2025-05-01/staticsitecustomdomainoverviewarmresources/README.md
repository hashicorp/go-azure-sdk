
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitecustomdomainoverviewarmresources` Documentation

The `staticsitecustomdomainoverviewarmresources` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/staticsitecustomdomainoverviewarmresources"
```


### Client Initialization

```go
client := staticsitecustomdomainoverviewarmresources.NewStaticSiteCustomDomainOverviewARMResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StaticSiteCustomDomainOverviewARMResourcesClient.StaticSitesCreateOrUpdateStaticSiteCustomDomain`

```go
ctx := context.TODO()
id := staticsitecustomdomainoverviewarmresources.NewCustomDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "customDomainName")

payload := staticsitecustomdomainoverviewarmresources.StaticSiteCustomDomainRequestPropertiesARMResource{
	// ...
}


if err := client.StaticSitesCreateOrUpdateStaticSiteCustomDomainThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StaticSiteCustomDomainOverviewARMResourcesClient.StaticSitesDeleteStaticSiteCustomDomain`

```go
ctx := context.TODO()
id := staticsitecustomdomainoverviewarmresources.NewCustomDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "customDomainName")

if err := client.StaticSitesDeleteStaticSiteCustomDomainThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `StaticSiteCustomDomainOverviewARMResourcesClient.StaticSitesGetStaticSiteCustomDomain`

```go
ctx := context.TODO()
id := staticsitecustomdomainoverviewarmresources.NewCustomDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "customDomainName")

read, err := client.StaticSitesGetStaticSiteCustomDomain(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StaticSiteCustomDomainOverviewARMResourcesClient.StaticSitesListStaticSiteCustomDomains`

```go
ctx := context.TODO()
id := staticsitecustomdomainoverviewarmresources.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

// alternatively `client.StaticSitesListStaticSiteCustomDomains(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesListStaticSiteCustomDomainsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StaticSiteCustomDomainOverviewARMResourcesClient.StaticSitesValidateCustomDomainCanBeAddedToStaticSite`

```go
ctx := context.TODO()
id := staticsitecustomdomainoverviewarmresources.NewCustomDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "customDomainName")

payload := staticsitecustomdomainoverviewarmresources.StaticSiteCustomDomainRequestPropertiesARMResource{
	// ...
}


if err := client.StaticSitesValidateCustomDomainCanBeAddedToStaticSiteThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
