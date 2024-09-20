
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-05-01/dppfeaturesupport` Documentation

The `dppfeaturesupport` SDK allows for interaction with Azure Resource Manager `dataprotection` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-05-01/dppfeaturesupport"
```


### Client Initialization

```go
client := dppfeaturesupport.NewDppFeatureSupportClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DppFeatureSupportClient.DataProtectionCheckFeatureSupport`

```go
ctx := context.TODO()
id := dppfeaturesupport.NewLocationID("12345678-1234-9876-4563-123456789012", "location")

payload := dppfeaturesupport.FeatureValidationRequestBase{
	// ...
}


read, err := client.DataProtectionCheckFeatureSupport(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
