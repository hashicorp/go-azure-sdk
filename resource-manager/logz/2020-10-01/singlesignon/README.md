
## `github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/singlesignon` Documentation

The `singlesignon` SDK allows for interaction with Azure Resource Manager `logz` (API Version `2020-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/singlesignon"
```


### Client Initialization

```go
client := singlesignon.NewSingleSignOnClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SingleSignOnClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := singlesignon.NewSingleSignOnConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "singleSignOnConfigurationName")

payload := singlesignon.LogzSingleSignOnResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SingleSignOnClient.Get`

```go
ctx := context.TODO()
id := singlesignon.NewSingleSignOnConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "singleSignOnConfigurationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SingleSignOnClient.List`

```go
ctx := context.TODO()
id := singlesignon.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
