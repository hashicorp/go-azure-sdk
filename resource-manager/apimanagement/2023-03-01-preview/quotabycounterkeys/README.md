
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-03-01-preview/quotabycounterkeys` Documentation

The `quotabycounterkeys` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2023-03-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-03-01-preview/quotabycounterkeys"
```


### Client Initialization

```go
client := quotabycounterkeys.NewQuotaByCounterKeysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `QuotaByCounterKeysClient.ListByService`

```go
ctx := context.TODO()
id := quotabycounterkeys.NewQuotaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "quotaCounterKeyValue")

read, err := client.ListByService(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `QuotaByCounterKeysClient.Update`

```go
ctx := context.TODO()
id := quotabycounterkeys.NewQuotaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "quotaCounterKeyValue")

payload := quotabycounterkeys.QuotaCounterValueUpdateContract{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
