
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationtemplateschema` Documentation

The `synchronizationtemplateschema` SDK allows for interaction with the Azure Resource Manager Service `applications` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationtemplateschema"
```


### Client Initialization

```go
client := synchronizationtemplateschema.NewSynchronizationTemplateSchemaClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SynchronizationTemplateSchemaClient.DeleteSynchronizationTemplateSchema`

```go
ctx := context.TODO()
id := synchronizationtemplateschema.NewApplicationIdSynchronizationTemplateID("applicationIdValue", "synchronizationTemplateIdValue")

read, err := client.DeleteSynchronizationTemplateSchema(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationTemplateSchemaClient.GetSynchronizationTemplateSchema`

```go
ctx := context.TODO()
id := synchronizationtemplateschema.NewApplicationIdSynchronizationTemplateID("applicationIdValue", "synchronizationTemplateIdValue")

read, err := client.GetSynchronizationTemplateSchema(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationTemplateSchemaClient.ParseApplicationSynchronizationTemplateSchemaExpression`

```go
ctx := context.TODO()
id := synchronizationtemplateschema.NewApplicationIdSynchronizationTemplateID("applicationIdValue", "synchronizationTemplateIdValue")

payload := synchronizationtemplateschema.ParseApplicationSynchronizationTemplateSchemaExpressionRequest{
	// ...
}


read, err := client.ParseApplicationSynchronizationTemplateSchemaExpression(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationTemplateSchemaClient.UpdateSynchronizationTemplateSchema`

```go
ctx := context.TODO()
id := synchronizationtemplateschema.NewApplicationIdSynchronizationTemplateID("applicationIdValue", "synchronizationTemplateIdValue")

payload := synchronizationtemplateschema.SynchronizationSchema{
	// ...
}


read, err := client.UpdateSynchronizationTemplateSchema(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
