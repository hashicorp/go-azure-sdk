
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasentities` Documentation

The `groupquotasentities` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasentities"
```


### Client Initialization

```go
client := groupquotasentities.NewGroupQuotasEntitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupQuotasEntitiesClient.GroupQuotaLimitsRequestList`

```go
ctx := context.TODO()
id := groupquotasentities.NewResourceProviderID("managementGroupId", "groupQuotaName", "resourceProviderName")

// alternatively `client.GroupQuotaLimitsRequestList(ctx, id, groupquotasentities.DefaultGroupQuotaLimitsRequestListOperationOptions())` can be used to do batched pagination
items, err := client.GroupQuotaLimitsRequestListComplete(ctx, id, groupquotasentities.DefaultGroupQuotaLimitsRequestListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupQuotasEntitiesClient.GroupQuotaUsagesList`

```go
ctx := context.TODO()
id := groupquotasentities.NewLocationUsageID("managementGroupId", "groupQuotaName", "resourceProviderName", "locationUsageName")

// alternatively `client.GroupQuotaUsagesList(ctx, id)` can be used to do batched pagination
items, err := client.GroupQuotaUsagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupQuotasEntitiesClient.GroupQuotasCreateOrUpdate`

```go
ctx := context.TODO()
id := groupquotasentities.NewGroupQuotaID("managementGroupId", "groupQuotaName")

payload := groupquotasentities.GroupQuotasEntity{
	// ...
}


if err := client.GroupQuotasCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `GroupQuotasEntitiesClient.GroupQuotasDelete`

```go
ctx := context.TODO()
id := groupquotasentities.NewGroupQuotaID("managementGroupId", "groupQuotaName")

if err := client.GroupQuotasDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `GroupQuotasEntitiesClient.GroupQuotasGet`

```go
ctx := context.TODO()
id := groupquotasentities.NewGroupQuotaID("managementGroupId", "groupQuotaName")

read, err := client.GroupQuotasGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupQuotasEntitiesClient.GroupQuotasList`

```go
ctx := context.TODO()
id := commonids.NewManagementGroupID("groupId")

// alternatively `client.GroupQuotasList(ctx, id)` can be used to do batched pagination
items, err := client.GroupQuotasListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupQuotasEntitiesClient.GroupQuotasUpdate`

```go
ctx := context.TODO()
id := groupquotasentities.NewGroupQuotaID("managementGroupId", "groupQuotaName")

payload := groupquotasentities.GroupQuotasEntityPatch{
	// ...
}


if err := client.GroupQuotasUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
