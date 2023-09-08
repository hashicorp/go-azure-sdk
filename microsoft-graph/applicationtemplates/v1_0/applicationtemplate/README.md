
## `github.com/hashicorp/go-azure-sdk/resource-manager/applicationtemplates/v1.0/applicationtemplate` Documentation

The `applicationtemplate` SDK allows for interaction with the Azure Resource Manager Service `applicationtemplates` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/applicationtemplates/v1.0/applicationtemplate"
```


### Client Initialization

```go
client := applicationtemplate.NewApplicationTemplateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationTemplateClient.GetApplicationTemplateById`

```go
ctx := context.TODO()
id := applicationtemplate.NewApplicationTemplateID("applicationTemplateIdValue")

read, err := client.GetApplicationTemplateById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationTemplateClient.GetApplicationTemplateCount`

```go
ctx := context.TODO()


read, err := client.GetApplicationTemplateCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationTemplateClient.InstantiateApplicationTemplateById`

```go
ctx := context.TODO()
id := applicationtemplate.NewApplicationTemplateID("applicationTemplateIdValue")

payload := applicationtemplate.InstantiateApplicationTemplateByIdRequest{
	// ...
}


read, err := client.InstantiateApplicationTemplateById(ctx, id, payload)
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
