
## `github.com/hashicorp/go-azure-sdk/resource-manager/cloudhealth/2026-11-01/entities` Documentation

The `entities` SDK allows for interaction with Azure Resource Manager `cloudhealth` (API Version `2026-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cloudhealth/2026-11-01/entities"
```


### Client Initialization

```go
client := entities.NewEntitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntitiesClient.AddDataAnnotation`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "entityName")

payload := entities.AddDataAnnotationRequest{
	// ...
}


read, err := client.AddDataAnnotation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitiesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "entityName")

payload := entities.Entity{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `EntitiesClient.Delete`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "entityName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `EntitiesClient.Get`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "entityName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitiesClient.GetDataAnnotations`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "entityName")

payload := entities.GetDataAnnotationsRequest{
	// ...
}


read, err := client.GetDataAnnotations(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitiesClient.GetHistory`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "entityName")

payload := entities.EntityHistoryRequest{
	// ...
}


read, err := client.GetHistory(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitiesClient.GetSignalHistory`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "entityName")

payload := entities.SignalHistoryRequest{
	// ...
}


read, err := client.GetSignalHistory(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitiesClient.GetSignalRecommendations`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "entityName")

read, err := client.GetSignalRecommendations(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitiesClient.IngestHealthReport`

```go
ctx := context.TODO()
id := entities.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "entityName")

payload := entities.HealthReportRequest{
	// ...
}


read, err := client.IngestHealthReport(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitiesClient.ListByHealthModel`

```go
ctx := context.TODO()
id := entities.NewHealthmodelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName")

// alternatively `client.ListByHealthModel(ctx, id, entities.DefaultListByHealthModelOperationOptions())` can be used to do batched pagination
items, err := client.ListByHealthModelComplete(ctx, id, entities.DefaultListByHealthModelOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
