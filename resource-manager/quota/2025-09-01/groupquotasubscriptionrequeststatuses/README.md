
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasubscriptionrequeststatuses` Documentation

The `groupquotasubscriptionrequeststatuses` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasubscriptionrequeststatuses"
```


### Client Initialization

```go
client := groupquotasubscriptionrequeststatuses.NewGroupQuotaSubscriptionRequestStatusesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupQuotaSubscriptionRequestStatusesClient.GroupQuotaSubscriptionRequestsGet`

```go
ctx := context.TODO()
id := groupquotasubscriptionrequeststatuses.NewSubscriptionRequestID("managementGroupId", "groupQuotaName", "requestId")

read, err := client.GroupQuotaSubscriptionRequestsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupQuotaSubscriptionRequestStatusesClient.GroupQuotaSubscriptionRequestsList`

```go
ctx := context.TODO()
id := groupquotasubscriptionrequeststatuses.NewGroupQuotaID("managementGroupId", "groupQuotaName")

// alternatively `client.GroupQuotaSubscriptionRequestsList(ctx, id)` can be used to do batched pagination
items, err := client.GroupQuotaSubscriptionRequestsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
