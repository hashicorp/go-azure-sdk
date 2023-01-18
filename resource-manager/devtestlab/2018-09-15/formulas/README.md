
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/formulas` Documentation

The `formulas` SDK allows for interaction with the Azure Resource Manager Service `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/formulas"
```


### Client Initialization

```go
client := formulas.NewFormulasClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FormulasClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := formulas.NewFormulaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "nameValue")

payload := formulas.Formula{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `FormulasClient.Delete`

```go
ctx := context.TODO()
id := formulas.NewFormulaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "nameValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FormulasClient.Get`

```go
ctx := context.TODO()
id := formulas.NewFormulaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "nameValue")

read, err := client.Get(ctx, id, formulas.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FormulasClient.List`

```go
ctx := context.TODO()
id := formulas.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue")

// alternatively `client.List(ctx, id, formulas.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, formulas.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FormulasClient.Update`

```go
ctx := context.TODO()
id := formulas.NewFormulaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "nameValue")

payload := formulas.UpdateResource{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
