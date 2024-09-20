
## `github.com/hashicorp/go-azure-sdk/resource-manager/insights/2018-03-01/metricalertsstatus` Documentation

The `metricalertsstatus` SDK allows for interaction with Azure Resource Manager `insights` (API Version `2018-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2018-03-01/metricalertsstatus"
```


### Client Initialization

```go
client := metricalertsstatus.NewMetricAlertsStatusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MetricAlertsStatusClient.List`

```go
ctx := context.TODO()
id := metricalertsstatus.NewMetricAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "ruleName")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MetricAlertsStatusClient.ListByName`

```go
ctx := context.TODO()
id := metricalertsstatus.NewStatusID("12345678-1234-9876-4563-123456789012", "example-resource-group", "ruleName", "statusName")

read, err := client.ListByName(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
