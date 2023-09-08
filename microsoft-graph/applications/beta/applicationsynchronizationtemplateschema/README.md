
## `github.com/hashicorp/go-azure-sdk/resource-manager/applications/beta/applicationsynchronizationtemplateschema` Documentation

The `applicationsynchronizationtemplateschema` SDK allows for interaction with the Azure Resource Manager Service `applications` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/applications/beta/applicationsynchronizationtemplateschema"
```


### Client Initialization

```go
client := applicationsynchronizationtemplateschema.NewApplicationSynchronizationTemplateSchemaClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationSynchronizationTemplateSchemaClient.DeleteApplicationByIdSynchronizationTemplateByIdSchema`

```go
ctx := context.TODO()
id := applicationsynchronizationtemplateschema.NewApplicationSynchronizationTemplateID("applicationIdValue", "synchronizationTemplateIdValue")

read, err := client.DeleteApplicationByIdSynchronizationTemplateByIdSchema(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationTemplateSchemaClient.GetApplicationByIdSynchronizationTemplateByIdSchema`

```go
ctx := context.TODO()
id := applicationsynchronizationtemplateschema.NewApplicationSynchronizationTemplateID("applicationIdValue", "synchronizationTemplateIdValue")

read, err := client.GetApplicationByIdSynchronizationTemplateByIdSchema(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationTemplateSchemaClient.ParseApplicationByIdSynchronizationTemplateByIdSchemaExpression`

```go
ctx := context.TODO()
id := applicationsynchronizationtemplateschema.NewApplicationSynchronizationTemplateID("applicationIdValue", "synchronizationTemplateIdValue")

payload := applicationsynchronizationtemplateschema.ParseApplicationByIdSynchronizationTemplateByIdSchemaExpressionRequest{
	// ...
}


read, err := client.ParseApplicationByIdSynchronizationTemplateByIdSchemaExpression(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationTemplateSchemaClient.UpdateApplicationByIdSynchronizationTemplateByIdSchema`

```go
ctx := context.TODO()
id := applicationsynchronizationtemplateschema.NewApplicationSynchronizationTemplateID("applicationIdValue", "synchronizationTemplateIdValue")

payload := applicationsynchronizationtemplateschema.SynchronizationSchema{
	// ...
}


read, err := client.UpdateApplicationByIdSynchronizationTemplateByIdSchema(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
