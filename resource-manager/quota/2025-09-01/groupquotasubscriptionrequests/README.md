
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasubscriptionrequests` Documentation

The `groupquotasubscriptionrequests` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasubscriptionrequests"
```


### Client Initialization

```go
client := groupquotasubscriptionrequests.NewGroupQuotaSubscriptionRequestsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupQuotaSubscriptionRequestsClient.Get`

```go
ctx := context.TODO()
id := groupquotasubscriptionrequests.NewSubscriptionRequestID("managementGroupId", "groupQuotaName", "requestId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupQuotaSubscriptionRequestsClient.List`

```go
ctx := context.TODO()
id := groupquotasubscriptionrequests.NewGroupQuotaID("managementGroupId", "groupQuotaName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
