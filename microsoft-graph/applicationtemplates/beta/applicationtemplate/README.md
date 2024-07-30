
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/applicationtemplates/beta/applicationtemplate` Documentation

The `applicationtemplate` SDK allows for interaction with the Azure Resource Manager Service `applicationtemplates` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/applicationtemplates/beta/applicationtemplate"
```


### Client Initialization

```go
client := applicationtemplate.NewApplicationTemplateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationTemplateClient.GetApplicationTemplate`

```go
ctx := context.TODO()
id := applicationtemplate.NewApplicationTemplateID("applicationTemplateIdValue")

read, err := client.GetApplicationTemplate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationTemplateClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationTemplateClient.InstantiateApplicationTemplate`

```go
ctx := context.TODO()
id := applicationtemplate.NewApplicationTemplateID("applicationTemplateIdValue")

payload := applicationtemplate.InstantiateApplicationTemplateRequest{
	// ...
}


read, err := client.InstantiateApplicationTemplate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationTemplateClient.ListApplicationTemplates`

```go
ctx := context.TODO()


// alternatively `client.ListApplicationTemplates(ctx)` can be used to do batched pagination
items, err := client.ListApplicationTemplatesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
