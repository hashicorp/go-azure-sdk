
## `github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/sessionhostconfiguration` Documentation

The `sessionhostconfiguration` SDK allows for interaction with Azure Resource Manager `desktopvirtualization` (API Version `2024-04-08-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/sessionhostconfiguration"
```


### Client Initialization

```go
client := sessionhostconfiguration.NewSessionHostConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SessionHostConfigurationClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := sessionhostconfiguration.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName")

payload := sessionhostconfiguration.SessionHostConfiguration{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SessionHostConfigurationClient.Get`

```go
ctx := context.TODO()
id := sessionhostconfiguration.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SessionHostConfigurationClient.ListByHostPool`

```go
ctx := context.TODO()
id := sessionhostconfiguration.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName")

// alternatively `client.ListByHostPool(ctx, id)` can be used to do batched pagination
items, err := client.ListByHostPoolComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SessionHostConfigurationClient.Update`

```go
ctx := context.TODO()
id := sessionhostconfiguration.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName")

payload := sessionhostconfiguration.SessionHostConfigurationPatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
