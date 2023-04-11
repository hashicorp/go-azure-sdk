
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/iotsecuritysolutions` Documentation

The `iotsecuritysolutions` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2017-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/iotsecuritysolutions"
```


### Client Initialization

```go
client := iotsecuritysolutions.NewIotSecuritySolutionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IotSecuritySolutionsClient.IoTSecuritySolutionsList`

```go
ctx := context.TODO()
id := iotsecuritysolutions.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.IoTSecuritySolutionsList(ctx, id, iotsecuritysolutions.DefaultIoTSecuritySolutionsListOperationOptions())` can be used to do batched pagination
items, err := client.IoTSecuritySolutionsListComplete(ctx, id, iotsecuritysolutions.DefaultIoTSecuritySolutionsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotSecuritySolutionsClient.IoTSecuritySolutionsResourceGroupList`

```go
ctx := context.TODO()
id := iotsecuritysolutions.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.IoTSecuritySolutionsResourceGroupList(ctx, id, iotsecuritysolutions.DefaultIoTSecuritySolutionsResourceGroupListOperationOptions())` can be used to do batched pagination
items, err := client.IoTSecuritySolutionsResourceGroupListComplete(ctx, id, iotsecuritysolutions.DefaultIoTSecuritySolutionsResourceGroupListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotSecuritySolutionsClient.IotSecuritySolutionCreate`

```go
ctx := context.TODO()
id := iotsecuritysolutions.NewIotSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue")

payload := iotsecuritysolutions.IoTSecuritySolutionModel{
	// ...
}


read, err := client.IotSecuritySolutionCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotSecuritySolutionsClient.IotSecuritySolutionDelete`

```go
ctx := context.TODO()
id := iotsecuritysolutions.NewIotSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue")

read, err := client.IotSecuritySolutionDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotSecuritySolutionsClient.IotSecuritySolutionGet`

```go
ctx := context.TODO()
id := iotsecuritysolutions.NewIotSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue")

read, err := client.IotSecuritySolutionGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotSecuritySolutionsClient.IotSecuritySolutionUpdate`

```go
ctx := context.TODO()
id := iotsecuritysolutions.NewIotSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue")

payload := iotsecuritysolutions.UpdateIotSecuritySolutionData{
	// ...
}


read, err := client.IotSecuritySolutionUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
