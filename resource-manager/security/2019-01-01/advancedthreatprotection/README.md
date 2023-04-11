
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01/advancedthreatprotection` Documentation

The `advancedthreatprotection` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2019-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01/advancedthreatprotection"
```


### Client Initialization

```go
client := advancedthreatprotection.NewAdvancedThreatProtectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AdvancedThreatProtectionClient.Create`

```go
ctx := context.TODO()
id := advancedthreatprotection.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

payload := advancedthreatprotection.AdvancedThreatProtectionSetting{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdvancedThreatProtectionClient.Get`

```go
ctx := context.TODO()
id := advancedthreatprotection.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
