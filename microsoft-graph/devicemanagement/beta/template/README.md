
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/template` Documentation

The `template` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/template"
```


### Client Initialization

```go
client := template.NewTemplateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TemplateClient.CreateTemplate`

```go
ctx := context.TODO()

payload := template.DeviceManagementTemplate{
	// ...
}


read, err := client.CreateTemplate(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateClient.CreateTemplateCreateInstance`

```go
ctx := context.TODO()
id := template.NewDeviceManagementTemplateID("deviceManagementTemplateIdValue")

payload := template.CreateTemplateCreateInstanceRequest{
	// ...
}


read, err := client.CreateTemplateCreateInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateClient.DeleteTemplate`

```go
ctx := context.TODO()
id := template.NewDeviceManagementTemplateID("deviceManagementTemplateIdValue")

read, err := client.DeleteTemplate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateClient.GetTemplate`

```go
ctx := context.TODO()
id := template.NewDeviceManagementTemplateID("deviceManagementTemplateIdValue")

read, err := client.GetTemplate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateClient.GetTemplateCount`

```go
ctx := context.TODO()


read, err := client.GetTemplateCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateClient.ListTemplateImportOffice365DeviceConfigurationPolicies`

```go
ctx := context.TODO()


// alternatively `client.ListTemplateImportOffice365DeviceConfigurationPolicies(ctx)` can be used to do batched pagination
items, err := client.ListTemplateImportOffice365DeviceConfigurationPoliciesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TemplateClient.ListTemplates`

```go
ctx := context.TODO()


// alternatively `client.ListTemplates(ctx)` can be used to do batched pagination
items, err := client.ListTemplatesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TemplateClient.UpdateTemplate`

```go
ctx := context.TODO()
id := template.NewDeviceManagementTemplateID("deviceManagementTemplateIdValue")

payload := template.DeviceManagementTemplate{
	// ...
}


read, err := client.UpdateTemplate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
