
## `github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2023-05-01/sourcecontrolconfiguration` Documentation

The `sourcecontrolconfiguration` SDK allows for interaction with Azure Resource Manager `kubernetesconfiguration` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2023-05-01/sourcecontrolconfiguration"
```


### Client Initialization

```go
client := sourcecontrolconfiguration.NewSourceControlConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SourceControlConfigurationClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := sourcecontrolconfiguration.NewScopedSourceControlConfigurationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "sourceControlConfigurationName")

payload := sourcecontrolconfiguration.SourceControlConfiguration{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SourceControlConfigurationClient.Delete`

```go
ctx := context.TODO()
id := sourcecontrolconfiguration.NewScopedSourceControlConfigurationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "sourceControlConfigurationName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SourceControlConfigurationClient.Get`

```go
ctx := context.TODO()
id := sourcecontrolconfiguration.NewScopedSourceControlConfigurationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "sourceControlConfigurationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SourceControlConfigurationClient.List`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
