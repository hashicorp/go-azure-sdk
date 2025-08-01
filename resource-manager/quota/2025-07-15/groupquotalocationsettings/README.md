
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-07-15/groupquotalocationsettings` Documentation

The `groupquotalocationsettings` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-07-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-07-15/groupquotalocationsettings"
```


### Client Initialization

```go
client := groupquotalocationsettings.NewGroupQuotaLocationSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupQuotaLocationSettingsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := groupquotalocationsettings.NewLocationSettingID("managementGroupId", "groupQuotaName", "resourceProviderName", "locationSettingName")

payload := groupquotalocationsettings.GroupQuotasEnforcementStatus{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `GroupQuotaLocationSettingsClient.Get`

```go
ctx := context.TODO()
id := groupquotalocationsettings.NewLocationSettingID("managementGroupId", "groupQuotaName", "resourceProviderName", "locationSettingName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupQuotaLocationSettingsClient.List`

```go
ctx := context.TODO()
id := groupquotalocationsettings.NewResourceProviderID("managementGroupId", "groupQuotaName", "resourceProviderName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupQuotaLocationSettingsClient.Update`

```go
ctx := context.TODO()
id := groupquotalocationsettings.NewLocationSettingID("managementGroupId", "groupQuotaName", "resourceProviderName", "locationSettingName")

payload := groupquotalocationsettings.GroupQuotasEnforcementStatus{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
