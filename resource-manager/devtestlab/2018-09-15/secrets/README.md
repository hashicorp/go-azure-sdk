
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/secrets` Documentation

The `secrets` SDK allows for interaction with the Azure Resource Manager Service `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/secrets"
```


### Client Initialization

```go
client := secrets.NewSecretsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SecretsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := secrets.NewSecretID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "nameValue")

payload := secrets.Secret{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SecretsClient.Delete`

```go
ctx := context.TODO()
id := secrets.NewSecretID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "nameValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecretsClient.Get`

```go
ctx := context.TODO()
id := secrets.NewSecretID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "nameValue")

read, err := client.Get(ctx, id, secrets.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecretsClient.List`

```go
ctx := context.TODO()
id := secrets.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "nameValue")

// alternatively `client.List(ctx, id, secrets.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, secrets.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SecretsClient.Update`

```go
ctx := context.TODO()
id := secrets.NewSecretID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "nameValue")

payload := secrets.UpdateResource{
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
