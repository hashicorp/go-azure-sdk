
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/customdnssuffixconfigurations` Documentation

The `customdnssuffixconfigurations` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/customdnssuffixconfigurations"
```


### Client Initialization

```go
client := customdnssuffixconfigurations.NewCustomDnsSuffixConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CustomDnsSuffixConfigurationsClient.AppServiceEnvironmentsDeleteAseCustomDnsSuffixConfiguration`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.AppServiceEnvironmentsDeleteAseCustomDnsSuffixConfiguration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomDnsSuffixConfigurationsClient.AppServiceEnvironmentsGetAseCustomDnsSuffixConfiguration`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.AppServiceEnvironmentsGetAseCustomDnsSuffixConfiguration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomDnsSuffixConfigurationsClient.AppServiceEnvironmentsUpdateAseCustomDnsSuffixConfiguration`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

payload := customdnssuffixconfigurations.CustomDnsSuffixConfiguration{
	// ...
}


read, err := client.AppServiceEnvironmentsUpdateAseCustomDnsSuffixConfiguration(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
