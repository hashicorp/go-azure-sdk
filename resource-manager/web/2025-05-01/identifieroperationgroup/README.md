
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/identifieroperationgroup` Documentation

The `identifieroperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/identifieroperationgroup"
```


### Client Initialization

```go
client := identifieroperationgroup.NewIdentifierOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IdentifierOperationGroupClient.WebAppsCreateOrUpdateDomainOwnershipIdentifierSlot`

```go
ctx := context.TODO()
id := identifieroperationgroup.NewSlotDomainOwnershipIdentifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "domainOwnershipIdentifierName")

payload := identifieroperationgroup.Identifier{
	// ...
}


read, err := client.WebAppsCreateOrUpdateDomainOwnershipIdentifierSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentifierOperationGroupClient.WebAppsDeleteDomainOwnershipIdentifierSlot`

```go
ctx := context.TODO()
id := identifieroperationgroup.NewSlotDomainOwnershipIdentifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "domainOwnershipIdentifierName")

read, err := client.WebAppsDeleteDomainOwnershipIdentifierSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentifierOperationGroupClient.WebAppsGetDomainOwnershipIdentifierSlot`

```go
ctx := context.TODO()
id := identifieroperationgroup.NewSlotDomainOwnershipIdentifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "domainOwnershipIdentifierName")

read, err := client.WebAppsGetDomainOwnershipIdentifierSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentifierOperationGroupClient.WebAppsListDomainOwnershipIdentifiersSlot`

```go
ctx := context.TODO()
id := identifieroperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListDomainOwnershipIdentifiersSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListDomainOwnershipIdentifiersSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IdentifierOperationGroupClient.WebAppsUpdateDomainOwnershipIdentifierSlot`

```go
ctx := context.TODO()
id := identifieroperationgroup.NewSlotDomainOwnershipIdentifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "domainOwnershipIdentifierName")

payload := identifieroperationgroup.Identifier{
	// ...
}


read, err := client.WebAppsUpdateDomainOwnershipIdentifierSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
