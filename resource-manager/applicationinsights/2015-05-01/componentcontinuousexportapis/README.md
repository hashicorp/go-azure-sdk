
## `github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentcontinuousexportapis` Documentation

The `componentcontinuousexportapis` SDK allows for interaction with Azure Resource Manager `applicationinsights` (API Version `2015-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentcontinuousexportapis"
```


### Client Initialization

```go
client := componentcontinuousexportapis.NewComponentContinuousExportAPIsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComponentContinuousExportAPIsClient.ExportConfigurationsCreate`

```go
ctx := context.TODO()
id := componentcontinuousexportapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

payload := componentcontinuousexportapis.ApplicationInsightsComponentExportRequest{
	// ...
}


read, err := client.ExportConfigurationsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentContinuousExportAPIsClient.ExportConfigurationsDelete`

```go
ctx := context.TODO()
id := componentcontinuousexportapis.NewExportConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "exportId")

read, err := client.ExportConfigurationsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentContinuousExportAPIsClient.ExportConfigurationsGet`

```go
ctx := context.TODO()
id := componentcontinuousexportapis.NewExportConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "exportId")

read, err := client.ExportConfigurationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentContinuousExportAPIsClient.ExportConfigurationsList`

```go
ctx := context.TODO()
id := componentcontinuousexportapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

read, err := client.ExportConfigurationsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentContinuousExportAPIsClient.ExportConfigurationsUpdate`

```go
ctx := context.TODO()
id := componentcontinuousexportapis.NewExportConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "exportId")

payload := componentcontinuousexportapis.ApplicationInsightsComponentExportRequest{
	// ...
}


read, err := client.ExportConfigurationsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
