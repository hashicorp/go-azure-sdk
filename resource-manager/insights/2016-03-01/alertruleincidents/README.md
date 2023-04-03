
## `github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/alertruleincidents` Documentation

The `alertruleincidents` SDK allows for interaction with the Azure Resource Manager Service `insights` (API Version `2016-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/alertruleincidents"
```


### Client Initialization

```go
client := alertruleincidents.NewAlertRuleIncidentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AlertRuleIncidentsClient.Get`

```go
ctx := context.TODO()
id := alertruleincidents.NewIncidentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "alertRuleValue", "incidentValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertRuleIncidentsClient.ListByAlertRule`

```go
ctx := context.TODO()
id := alertruleincidents.NewAlertRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "alertRuleValue")

read, err := client.ListByAlertRule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
