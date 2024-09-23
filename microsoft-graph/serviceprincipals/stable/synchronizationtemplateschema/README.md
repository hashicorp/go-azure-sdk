
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationtemplateschema` Documentation

The `synchronizationtemplateschema` SDK allows for interaction with Microsoft Graph `serviceprincipals` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationtemplateschema"
```


### Client Initialization

```go
client := synchronizationtemplateschema.NewSynchronizationTemplateSchemaClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SynchronizationTemplateSchemaClient.DeleteSynchronizationTemplateSchema`

```go
ctx := context.TODO()
id := synchronizationtemplateschema.NewServicePrincipalIdSynchronizationTemplateID("servicePrincipalId", "synchronizationTemplateId")

read, err := client.DeleteSynchronizationTemplateSchema(ctx, id, synchronizationtemplateschema.DefaultDeleteSynchronizationTemplateSchemaOperationOptions())
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
id := synchronizationtemplateschema.NewServicePrincipalIdSynchronizationTemplateID("servicePrincipalId", "synchronizationTemplateId")

read, err := client.GetSynchronizationTemplateSchema(ctx, id, synchronizationtemplateschema.DefaultGetSynchronizationTemplateSchemaOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationTemplateSchemaClient.ParseSynchronizationTemplateSchemaExpression`

```go
ctx := context.TODO()
id := synchronizationtemplateschema.NewServicePrincipalIdSynchronizationTemplateID("servicePrincipalId", "synchronizationTemplateId")

payload := synchronizationtemplateschema.ParseSynchronizationTemplateSchemaExpressionRequest{
	// ...
}


read, err := client.ParseSynchronizationTemplateSchemaExpression(ctx, id, payload, synchronizationtemplateschema.DefaultParseSynchronizationTemplateSchemaExpressionOperationOptions())
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
id := synchronizationtemplateschema.NewServicePrincipalIdSynchronizationTemplateID("servicePrincipalId", "synchronizationTemplateId")

payload := synchronizationtemplateschema.SynchronizationSchema{
	// ...
}


read, err := client.UpdateSynchronizationTemplateSchema(ctx, id, payload, synchronizationtemplateschema.DefaultUpdateSynchronizationTemplateSchemaOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
