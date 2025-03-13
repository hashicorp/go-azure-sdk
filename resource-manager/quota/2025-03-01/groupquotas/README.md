
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-03-01/groupquotas` Documentation

The `groupquotas` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-03-01/groupquotas"
```


### Client Initialization

```go
client := groupquotas.NewGroupQuotasClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupQuotasClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := groupquotas.NewGroupQuotaID("managementGroupId", "groupQuotaName")

payload := groupquotas.GroupQuotasEntity{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `GroupQuotasClient.Delete`

```go
ctx := context.TODO()
id := groupquotas.NewGroupQuotaID("managementGroupId", "groupQuotaName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `GroupQuotasClient.Get`

```go
ctx := context.TODO()
id := groupquotas.NewGroupQuotaID("managementGroupId", "groupQuotaName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupQuotasClient.List`

```go
ctx := context.TODO()
id := commonids.NewManagementGroupID("groupId")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupQuotasClient.Update`

```go
ctx := context.TODO()
id := groupquotas.NewGroupQuotaID("managementGroupId", "groupQuotaName")

payload := groupquotas.GroupQuotasEntityPatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
