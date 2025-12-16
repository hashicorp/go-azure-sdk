
## `github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/dynatracesinglesignonresources` Documentation

The `dynatracesinglesignonresources` SDK allows for interaction with Azure Resource Manager `dynatrace` (API Version `2024-04-24`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/dynatracesinglesignonresources"
```


### Client Initialization

```go
client := dynatracesinglesignonresources.NewDynatraceSingleSignOnResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DynatraceSingleSignOnResourcesClient.SingleSignOnCreateOrUpdate`

```go
ctx := context.TODO()
id := dynatracesinglesignonresources.NewSingleSignOnConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "singleSignOnConfigurationName")

payload := dynatracesinglesignonresources.DynatraceSingleSignOnResource{
	// ...
}


if err := client.SingleSignOnCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DynatraceSingleSignOnResourcesClient.SingleSignOnGet`

```go
ctx := context.TODO()
id := dynatracesinglesignonresources.NewSingleSignOnConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "singleSignOnConfigurationName")

read, err := client.SingleSignOnGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DynatraceSingleSignOnResourcesClient.SingleSignOnList`

```go
ctx := context.TODO()
id := dynatracesinglesignonresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

// alternatively `client.SingleSignOnList(ctx, id)` can be used to do batched pagination
items, err := client.SingleSignOnListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
