
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-06-01-preview/quotausagesforflexibleservers` Documentation

The `quotausagesforflexibleservers` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2023-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-06-01-preview/quotausagesforflexibleservers"
```


### Client Initialization

```go
client := quotausagesforflexibleservers.NewQuotaUsagesForFlexibleServersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `QuotaUsagesForFlexibleServersClient.QuotaUsagesList`

```go
ctx := context.TODO()
id := quotausagesforflexibleservers.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.QuotaUsagesList(ctx, id)` can be used to do batched pagination
items, err := client.QuotaUsagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
