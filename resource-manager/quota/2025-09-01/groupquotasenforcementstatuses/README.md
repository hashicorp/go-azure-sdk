
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasenforcementstatuses` Documentation

The `groupquotasenforcementstatuses` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasenforcementstatuses"
```


### Client Initialization

```go
client := groupquotasenforcementstatuses.NewGroupQuotasEnforcementStatusesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupQuotasEnforcementStatusesClient.GroupQuotaLocationSettingsCreateOrUpdate`

```go
ctx := context.TODO()
id := groupquotasenforcementstatuses.NewLocationSettingID("managementGroupId", "groupQuotaName", "resourceProviderName", "locationSettingName")

payload := groupquotasenforcementstatuses.GroupQuotasEnforcementStatus{
	// ...
}


if err := client.GroupQuotaLocationSettingsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `GroupQuotasEnforcementStatusesClient.GroupQuotaLocationSettingsGet`

```go
ctx := context.TODO()
id := groupquotasenforcementstatuses.NewLocationSettingID("managementGroupId", "groupQuotaName", "resourceProviderName", "locationSettingName")

read, err := client.GroupQuotaLocationSettingsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupQuotasEnforcementStatusesClient.GroupQuotaLocationSettingsUpdate`

```go
ctx := context.TODO()
id := groupquotasenforcementstatuses.NewLocationSettingID("managementGroupId", "groupQuotaName", "resourceProviderName", "locationSettingName")

payload := groupquotasenforcementstatuses.GroupQuotasEnforcementStatus{
	// ...
}


if err := client.GroupQuotaLocationSettingsUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
