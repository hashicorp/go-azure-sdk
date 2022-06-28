
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-04-01/dppfeaturesupport` Documentation

The `dppfeaturesupport` SDK allows for interaction with the Azure Resource Manager Service `dataprotection` (API Version `2022-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-04-01/dppfeaturesupport"
```


### Client Initialization

```go
client := dppfeaturesupport.NewDppFeatureSupportClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `DppFeatureSupportClient.DataProtectionCheckFeatureSupport`

```go
ctx := context.TODO()
id := dppfeaturesupport.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

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
