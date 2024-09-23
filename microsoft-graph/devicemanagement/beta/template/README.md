
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/template` Documentation

The `template` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/template"
```


### Client Initialization

```go
client := template.NewTemplateClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TemplateClient.CreateTemplate`

```go
ctx := context.TODO()

payload := template.DeviceManagementTemplate{
	// ...
}


read, err := client.CreateTemplate(ctx, payload, template.DefaultCreateTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateClient.CreateTemplateInstance`

```go
ctx := context.TODO()
id := template.NewDeviceManagementTemplateID("deviceManagementTemplateId")

payload := template.CreateTemplateInstanceRequest{
	// ...
}


read, err := client.CreateTemplateInstance(ctx, id, payload, template.DefaultCreateTemplateInstanceOperationOptions())
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
id := template.NewDeviceManagementTemplateID("deviceManagementTemplateId")

read, err := client.DeleteTemplate(ctx, id, template.DefaultDeleteTemplateOperationOptions())
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
id := template.NewDeviceManagementTemplateID("deviceManagementTemplateId")

read, err := client.GetTemplate(ctx, id, template.DefaultGetTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateClient.GetTemplatesCount`

```go
ctx := context.TODO()


read, err := client.GetTemplatesCount(ctx, template.DefaultGetTemplatesCountOperationOptions())
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


// alternatively `client.ListTemplateImportOffice365DeviceConfigurationPolicies(ctx, template.DefaultListTemplateImportOffice365DeviceConfigurationPoliciesOperationOptions())` can be used to do batched pagination
items, err := client.ListTemplateImportOffice365DeviceConfigurationPoliciesComplete(ctx, template.DefaultListTemplateImportOffice365DeviceConfigurationPoliciesOperationOptions())
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


// alternatively `client.ListTemplates(ctx, template.DefaultListTemplatesOperationOptions())` can be used to do batched pagination
items, err := client.ListTemplatesComplete(ctx, template.DefaultListTemplatesOperationOptions())
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
id := template.NewDeviceManagementTemplateID("deviceManagementTemplateId")

payload := template.DeviceManagementTemplate{
	// ...
}


read, err := client.UpdateTemplate(ctx, id, payload, template.DefaultUpdateTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
