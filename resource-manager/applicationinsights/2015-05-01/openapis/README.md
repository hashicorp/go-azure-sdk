
## `github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/openapis` Documentation

The `openapis` SDK allows for interaction with Azure Resource Manager `applicationinsights` (API Version `2015-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/openapis"
```


### Client Initialization

```go
client := openapis.NewOpenapisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenapisClient.APIKeysCreate`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

payload := openapis.APIKeyRequest{
	// ...
}


read, err := client.APIKeysCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.APIKeysDelete`

```go
ctx := context.TODO()
id := openapis.NewApiKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "keyId")

read, err := client.APIKeysDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.APIKeysGet`

```go
ctx := context.TODO()
id := openapis.NewApiKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "keyId")

read, err := client.APIKeysGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.APIKeysList`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

// alternatively `client.APIKeysList(ctx, id)` can be used to do batched pagination
items, err := client.APIKeysListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.AnnotationsCreate`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

payload := openapis.Annotation{
	// ...
}


read, err := client.AnnotationsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.AnnotationsDelete`

```go
ctx := context.TODO()
id := openapis.NewAnnotationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "annotationId")

read, err := client.AnnotationsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.AnnotationsGet`

```go
ctx := context.TODO()
id := openapis.NewAnnotationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "annotationId")

read, err := client.AnnotationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.AnnotationsList`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

// alternatively `client.AnnotationsList(ctx, id, openapis.DefaultAnnotationsListOperationOptions())` can be used to do batched pagination
items, err := client.AnnotationsListComplete(ctx, id, openapis.DefaultAnnotationsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ComponentAvailableFeaturesGet`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

read, err := client.ComponentAvailableFeaturesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ComponentCurrentBillingFeaturesGet`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

read, err := client.ComponentCurrentBillingFeaturesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ComponentCurrentBillingFeaturesUpdate`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

payload := openapis.ApplicationInsightsComponentBillingFeatures{
	// ...
}


read, err := client.ComponentCurrentBillingFeaturesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ComponentFeatureCapabilitiesGet`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

read, err := client.ComponentFeatureCapabilitiesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ComponentQuotaStatusGet`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

read, err := client.ComponentQuotaStatusGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ExportConfigurationsCreate`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

payload := openapis.ApplicationInsightsComponentExportRequest{
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


### Example Usage: `OpenapisClient.ExportConfigurationsDelete`

```go
ctx := context.TODO()
id := openapis.NewExportConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "exportId")

read, err := client.ExportConfigurationsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ExportConfigurationsGet`

```go
ctx := context.TODO()
id := openapis.NewExportConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "exportId")

read, err := client.ExportConfigurationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ExportConfigurationsList`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

read, err := client.ExportConfigurationsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ExportConfigurationsUpdate`

```go
ctx := context.TODO()
id := openapis.NewExportConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "exportId")

payload := openapis.ApplicationInsightsComponentExportRequest{
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


### Example Usage: `OpenapisClient.ProactiveDetectionConfigurationsGet`

```go
ctx := context.TODO()
id := openapis.NewProactiveDetectionConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "configurationId")

read, err := client.ProactiveDetectionConfigurationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ProactiveDetectionConfigurationsList`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

read, err := client.ProactiveDetectionConfigurationsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ProactiveDetectionConfigurationsUpdate`

```go
ctx := context.TODO()
id := openapis.NewProactiveDetectionConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "configurationId")

payload := openapis.ApplicationInsightsComponentProactiveDetectionConfiguration{
	// ...
}


read, err := client.ProactiveDetectionConfigurationsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.WorkItemConfigurationsCreate`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

payload := openapis.WorkItemCreateConfiguration{
	// ...
}


read, err := client.WorkItemConfigurationsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.WorkItemConfigurationsDelete`

```go
ctx := context.TODO()
id := openapis.NewWorkItemConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "workItemConfigId")

read, err := client.WorkItemConfigurationsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.WorkItemConfigurationsGetDefault`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

read, err := client.WorkItemConfigurationsGetDefault(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.WorkItemConfigurationsGetItem`

```go
ctx := context.TODO()
id := openapis.NewWorkItemConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "workItemConfigId")

read, err := client.WorkItemConfigurationsGetItem(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.WorkItemConfigurationsList`

```go
ctx := context.TODO()
id := openapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

// alternatively `client.WorkItemConfigurationsList(ctx, id)` can be used to do batched pagination
items, err := client.WorkItemConfigurationsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.WorkItemConfigurationsUpdateItem`

```go
ctx := context.TODO()
id := openapis.NewWorkItemConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "workItemConfigId")

payload := openapis.WorkItemCreateConfiguration{
	// ...
}


read, err := client.WorkItemConfigurationsUpdateItem(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
