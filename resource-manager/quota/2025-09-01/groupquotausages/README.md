
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotausages` Documentation

The `groupquotausages` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotausages"
```


### Client Initialization

```go
client := groupquotausages.NewGroupQuotaUsagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupQuotaUsagesClient.List`

```go
ctx := context.TODO()
id := groupquotausages.NewLocationUsageID("managementGroupId", "groupQuotaName", "resourceProviderName", "locationUsageName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
