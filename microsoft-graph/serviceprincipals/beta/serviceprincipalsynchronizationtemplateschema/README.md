
## `github.com/hashicorp/go-azure-sdk/resource-manager/serviceprincipals/beta/serviceprincipalsynchronizationtemplateschema` Documentation

The `serviceprincipalsynchronizationtemplateschema` SDK allows for interaction with the Azure Resource Manager Service `serviceprincipals` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/serviceprincipals/beta/serviceprincipalsynchronizationtemplateschema"
```


### Client Initialization

```go
client := serviceprincipalsynchronizationtemplateschema.NewServicePrincipalSynchronizationTemplateSchemaClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicePrincipalSynchronizationTemplateSchemaClient.DeleteServicePrincipalByIdSynchronizationTemplateByIdSchema`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationtemplateschema.NewServicePrincipalSynchronizationTemplateID("servicePrincipalIdValue", "synchronizationTemplateIdValue")

read, err := client.DeleteServicePrincipalByIdSynchronizationTemplateByIdSchema(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationTemplateSchemaClient.GetServicePrincipalByIdSynchronizationTemplateByIdSchema`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationtemplateschema.NewServicePrincipalSynchronizationTemplateID("servicePrincipalIdValue", "synchronizationTemplateIdValue")

read, err := client.GetServicePrincipalByIdSynchronizationTemplateByIdSchema(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationTemplateSchemaClient.ParseServicePrincipalByIdSynchronizationTemplateByIdSchemaExpression`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationtemplateschema.NewServicePrincipalSynchronizationTemplateID("servicePrincipalIdValue", "synchronizationTemplateIdValue")

payload := serviceprincipalsynchronizationtemplateschema.ParseServicePrincipalByIdSynchronizationTemplateByIdSchemaExpressionRequest{
	// ...
}


read, err := client.ParseServicePrincipalByIdSynchronizationTemplateByIdSchemaExpression(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationTemplateSchemaClient.UpdateServicePrincipalByIdSynchronizationTemplateByIdSchema`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationtemplateschema.NewServicePrincipalSynchronizationTemplateID("servicePrincipalIdValue", "synchronizationTemplateIdValue")

payload := serviceprincipalsynchronizationtemplateschema.SynchronizationSchema{
	// ...
}


read, err := client.UpdateServicePrincipalByIdSynchronizationTemplateByIdSchema(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
