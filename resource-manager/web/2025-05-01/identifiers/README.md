
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/identifiers` Documentation

The `identifiers` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/identifiers"
```


### Client Initialization

```go
client := identifiers.NewIdentifiersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IdentifiersClient.WebAppsCreateOrUpdateDomainOwnershipIdentifier`

```go
ctx := context.TODO()
id := identifiers.NewDomainOwnershipIdentifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "domainOwnershipIdentifierName")

payload := identifiers.Identifier{
	// ...
}


read, err := client.WebAppsCreateOrUpdateDomainOwnershipIdentifier(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentifiersClient.WebAppsDeleteDomainOwnershipIdentifier`

```go
ctx := context.TODO()
id := identifiers.NewDomainOwnershipIdentifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "domainOwnershipIdentifierName")

read, err := client.WebAppsDeleteDomainOwnershipIdentifier(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentifiersClient.WebAppsGetDomainOwnershipIdentifier`

```go
ctx := context.TODO()
id := identifiers.NewDomainOwnershipIdentifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "domainOwnershipIdentifierName")

read, err := client.WebAppsGetDomainOwnershipIdentifier(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentifiersClient.WebAppsListDomainOwnershipIdentifiers`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListDomainOwnershipIdentifiers(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListDomainOwnershipIdentifiersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IdentifiersClient.WebAppsUpdateDomainOwnershipIdentifier`

```go
ctx := context.TODO()
id := identifiers.NewDomainOwnershipIdentifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "domainOwnershipIdentifierName")

payload := identifiers.Identifier{
	// ...
}


read, err := client.WebAppsUpdateDomainOwnershipIdentifier(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
