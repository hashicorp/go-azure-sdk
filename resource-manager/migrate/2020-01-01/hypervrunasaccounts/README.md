
## `github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervrunasaccounts` Documentation

The `hypervrunasaccounts` SDK allows for interaction with the Azure Resource Manager Service `migrate` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervrunasaccounts"
```


### Client Initialization

```go
client := hypervrunasaccounts.NewHyperVRunAsAccountsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HyperVRunAsAccountsClient.GetAllRunAsAccountsInSite`

```go
ctx := context.TODO()
id := hypervrunasaccounts.NewHyperVSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteValue")

// alternatively `client.GetAllRunAsAccountsInSite(ctx, id)` can be used to do batched pagination
items, err := client.GetAllRunAsAccountsInSiteComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `HyperVRunAsAccountsClient.GetRunAsAccount`

```go
ctx := context.TODO()
id := hypervrunasaccounts.NewHyperVSiteRunAsAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteValue", "runAsAccountValue")

read, err := client.GetRunAsAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
