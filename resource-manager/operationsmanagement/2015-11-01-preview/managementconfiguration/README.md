
## `github.com/hashicorp/go-azure-sdk/resource-manager/operationsmanagement/2015-11-01-preview/managementconfiguration` Documentation

The `managementconfiguration` SDK allows for interaction with Azure Resource Manager `operationsmanagement` (API Version `2015-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/operationsmanagement/2015-11-01-preview/managementconfiguration"
```


### Client Initialization

```go
client := managementconfiguration.NewManagementConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagementConfigurationClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := managementconfiguration.NewManagementConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managementConfigurationName")

payload := managementconfiguration.ManagementConfiguration{
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


### Example Usage: `ManagementConfigurationClient.Delete`

```go
ctx := context.TODO()
id := managementconfiguration.NewManagementConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managementConfigurationName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementConfigurationClient.Get`

```go
ctx := context.TODO()
id := managementconfiguration.NewManagementConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managementConfigurationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementConfigurationClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ListBySubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
