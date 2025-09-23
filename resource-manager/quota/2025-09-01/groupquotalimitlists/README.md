
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotalimitlists` Documentation

The `groupquotalimitlists` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotalimitlists"
```


### Client Initialization

```go
client := groupquotalimitlists.NewGroupQuotaLimitListsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupQuotaLimitListsClient.GroupQuotaLimitsList`

```go
ctx := context.TODO()
id := groupquotalimitlists.NewGroupQuotaLimitID("managementGroupId", "groupQuotaName", "resourceProviderName", "groupQuotaLimitName")

read, err := client.GroupQuotaLimitsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupQuotaLimitListsClient.GroupQuotaLimitsRequestUpdate`

```go
ctx := context.TODO()
id := groupquotalimitlists.NewGroupQuotaLimitID("managementGroupId", "groupQuotaName", "resourceProviderName", "groupQuotaLimitName")

payload := groupquotalimitlists.GroupQuotaLimitList{
	// ...
}


if err := client.GroupQuotaLimitsRequestUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
