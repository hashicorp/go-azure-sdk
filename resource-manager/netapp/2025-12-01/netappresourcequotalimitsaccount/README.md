
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/netappresourcequotalimitsaccount` Documentation

The `netappresourcequotalimitsaccount` SDK allows for interaction with Azure Resource Manager `netapp` (API Version `2025-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-12-01/netappresourcequotalimitsaccount"
```


### Client Initialization

```go
client := netappresourcequotalimitsaccount.NewNetAppResourceQuotaLimitsAccountClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetAppResourceQuotaLimitsAccountClient.Get`

```go
ctx := context.TODO()
id := netappresourcequotalimitsaccount.NewNetAppAccountQuotaLimitID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "quotaLimitName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceQuotaLimitsAccountClient.List`

```go
ctx := context.TODO()
id := netappresourcequotalimitsaccount.NewNetAppAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
