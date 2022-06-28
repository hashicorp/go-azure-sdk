
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/forecast` Documentation

The `forecast` SDK allows for interaction with the Azure Resource Manager Service `costmanagement` (API Version `2021-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/forecast"
```


### Client Initialization

```go
client := forecast.NewForecastClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ForecastClient.ExternalCloudProviderUsage`

```go
ctx := context.TODO()
id := forecast.NewExternalCloudProviderTypeID("externalBillingAccounts", "externalCloudProviderIdValue")

payload := forecast.ForecastDefinition{
	// ...
}


read, err := client.ExternalCloudProviderUsage(ctx, id, payload, forecast.DefaultExternalCloudProviderUsageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ForecastClient.Usage`

```go
ctx := context.TODO()
id := forecast.NewScopeID()

payload := forecast.ForecastDefinition{
	// ...
}


read, err := client.Usage(ctx, id, payload, forecast.DefaultUsageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
