
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templatemigratableto` Documentation

The `templatemigratableto` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templatemigratableto"
```


### Client Initialization

```go
client := templatemigratableto.NewTemplateMigratableToClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TemplateMigratableToClient.CreateTemplateMigratableTo`

```go
ctx := context.TODO()
id := templatemigratableto.NewDeviceManagementTemplateID("deviceManagementTemplateIdValue")

payload := templatemigratableto.DeviceManagementTemplate{
	// ...
}


read, err := client.CreateTemplateMigratableTo(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateMigratableToClient.CreateTemplateMigratableToCreateInstance`

```go
ctx := context.TODO()
id := templatemigratableto.NewDeviceManagementTemplateIdMigratableToID("deviceManagementTemplateIdValue", "deviceManagementTemplateId1Value")

payload := templatemigratableto.CreateTemplateMigratableToCreateInstanceRequest{
	// ...
}


read, err := client.CreateTemplateMigratableToCreateInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateMigratableToClient.DeleteTemplateMigratableTo`

```go
ctx := context.TODO()
id := templatemigratableto.NewDeviceManagementTemplateIdMigratableToID("deviceManagementTemplateIdValue", "deviceManagementTemplateId1Value")

read, err := client.DeleteTemplateMigratableTo(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateMigratableToClient.GetTemplateMigratableTo`

```go
ctx := context.TODO()
id := templatemigratableto.NewDeviceManagementTemplateIdMigratableToID("deviceManagementTemplateIdValue", "deviceManagementTemplateId1Value")

read, err := client.GetTemplateMigratableTo(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateMigratableToClient.GetTemplateMigratableToCount`

```go
ctx := context.TODO()
id := templatemigratableto.NewDeviceManagementTemplateID("deviceManagementTemplateIdValue")

read, err := client.GetTemplateMigratableToCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateMigratableToClient.ListTemplateMigratableToImportOffice365DeviceConfigurationPolicies`

```go
ctx := context.TODO()
id := templatemigratableto.NewDeviceManagementTemplateID("deviceManagementTemplateIdValue")

// alternatively `client.ListTemplateMigratableToImportOffice365DeviceConfigurationPolicies(ctx, id)` can be used to do batched pagination
items, err := client.ListTemplateMigratableToImportOffice365DeviceConfigurationPoliciesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TemplateMigratableToClient.ListTemplateMigratableTos`

```go
ctx := context.TODO()
id := templatemigratableto.NewDeviceManagementTemplateID("deviceManagementTemplateIdValue")

// alternatively `client.ListTemplateMigratableTos(ctx, id)` can be used to do batched pagination
items, err := client.ListTemplateMigratableTosComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TemplateMigratableToClient.UpdateTemplateMigratableTo`

```go
ctx := context.TODO()
id := templatemigratableto.NewDeviceManagementTemplateIdMigratableToID("deviceManagementTemplateIdValue", "deviceManagementTemplateId1Value")

payload := templatemigratableto.DeviceManagementTemplate{
	// ...
}


read, err := client.UpdateTemplateMigratableTo(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
