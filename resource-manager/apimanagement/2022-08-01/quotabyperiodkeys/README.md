
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/quotabyperiodkeys` Documentation

The `quotabyperiodkeys` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/quotabyperiodkeys"
```


### Client Initialization

```go
client := quotabyperiodkeys.NewQuotaByPeriodKeysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `QuotaByPeriodKeysClient.Get`

```go
ctx := context.TODO()
id := quotabyperiodkeys.NewPeriodID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "quotaCounterKeyValue", "quotaPeriodKeyValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `QuotaByPeriodKeysClient.Update`

```go
ctx := context.TODO()
id := quotabyperiodkeys.NewPeriodID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "quotaCounterKeyValue", "quotaPeriodKeyValue")

payload := quotabyperiodkeys.QuotaCounterValueUpdateContract{
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
