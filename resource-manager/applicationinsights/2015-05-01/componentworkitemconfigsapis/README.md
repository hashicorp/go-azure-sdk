
## `github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentworkitemconfigsapis` Documentation

The `componentworkitemconfigsapis` SDK allows for interaction with Azure Resource Manager `applicationinsights` (API Version `2015-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentworkitemconfigsapis"
```


### Client Initialization

```go
client := componentworkitemconfigsapis.NewComponentWorkItemConfigsAPIsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComponentWorkItemConfigsAPIsClient.WorkItemConfigurationsCreate`

```go
ctx := context.TODO()
id := componentworkitemconfigsapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

payload := componentworkitemconfigsapis.WorkItemCreateConfiguration{
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


### Example Usage: `ComponentWorkItemConfigsAPIsClient.WorkItemConfigurationsDelete`

```go
ctx := context.TODO()
id := componentworkitemconfigsapis.NewWorkItemConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "workItemConfigId")

read, err := client.WorkItemConfigurationsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentWorkItemConfigsAPIsClient.WorkItemConfigurationsGetDefault`

```go
ctx := context.TODO()
id := componentworkitemconfigsapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

read, err := client.WorkItemConfigurationsGetDefault(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentWorkItemConfigsAPIsClient.WorkItemConfigurationsGetItem`

```go
ctx := context.TODO()
id := componentworkitemconfigsapis.NewWorkItemConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "workItemConfigId")

read, err := client.WorkItemConfigurationsGetItem(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentWorkItemConfigsAPIsClient.WorkItemConfigurationsList`

```go
ctx := context.TODO()
id := componentworkitemconfigsapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

read, err := client.WorkItemConfigurationsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentWorkItemConfigsAPIsClient.WorkItemConfigurationsUpdateItem`

```go
ctx := context.TODO()
id := componentworkitemconfigsapis.NewWorkItemConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "workItemConfigId")

payload := componentworkitemconfigsapis.WorkItemCreateConfiguration{
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
