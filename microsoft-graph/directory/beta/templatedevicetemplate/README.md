
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/templatedevicetemplate` Documentation

The `templatedevicetemplate` SDK allows for interaction with Microsoft Graph `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/templatedevicetemplate"
```


### Client Initialization

```go
client := templatedevicetemplate.NewTemplateDeviceTemplateClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TemplateDeviceTemplateClient.CreateTemplateDeviceFromTemplate`

```go
ctx := context.TODO()
id := templatedevicetemplate.NewDirectoryTemplateDeviceTemplateID("deviceTemplateId")

payload := templatedevicetemplate.CreateTemplateDeviceFromTemplateRequest{
	// ...
}


read, err := client.CreateTemplateDeviceFromTemplate(ctx, id, payload, templatedevicetemplate.DefaultCreateTemplateDeviceFromTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateDeviceTemplateClient.CreateTemplateDeviceTemplate`

```go
ctx := context.TODO()

payload := templatedevicetemplate.DeviceTemplate{
	// ...
}


read, err := client.CreateTemplateDeviceTemplate(ctx, payload, templatedevicetemplate.DefaultCreateTemplateDeviceTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateDeviceTemplateClient.DeleteTemplateDeviceTemplate`

```go
ctx := context.TODO()
id := templatedevicetemplate.NewDirectoryTemplateDeviceTemplateID("deviceTemplateId")

read, err := client.DeleteTemplateDeviceTemplate(ctx, id, templatedevicetemplate.DefaultDeleteTemplateDeviceTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateDeviceTemplateClient.GetTemplateDeviceTemplate`

```go
ctx := context.TODO()
id := templatedevicetemplate.NewDirectoryTemplateDeviceTemplateID("deviceTemplateId")

read, err := client.GetTemplateDeviceTemplate(ctx, id, templatedevicetemplate.DefaultGetTemplateDeviceTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateDeviceTemplateClient.GetTemplateDeviceTemplatesCount`

```go
ctx := context.TODO()


read, err := client.GetTemplateDeviceTemplatesCount(ctx, templatedevicetemplate.DefaultGetTemplateDeviceTemplatesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateDeviceTemplateClient.ListTemplateDeviceTemplates`

```go
ctx := context.TODO()


// alternatively `client.ListTemplateDeviceTemplates(ctx, templatedevicetemplate.DefaultListTemplateDeviceTemplatesOperationOptions())` can be used to do batched pagination
items, err := client.ListTemplateDeviceTemplatesComplete(ctx, templatedevicetemplate.DefaultListTemplateDeviceTemplatesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TemplateDeviceTemplateClient.UpdateTemplateDeviceTemplate`

```go
ctx := context.TODO()
id := templatedevicetemplate.NewDirectoryTemplateDeviceTemplateID("deviceTemplateId")

payload := templatedevicetemplate.DeviceTemplate{
	// ...
}


read, err := client.UpdateTemplateDeviceTemplate(ctx, id, payload, templatedevicetemplate.DefaultUpdateTemplateDeviceTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
