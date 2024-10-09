
## `github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2023-11-22/secrets` Documentation

The `secrets` SDK allows for interaction with Azure Resource Manager `redhatopenshift` (API Version `2023-11-22`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2023-11-22/secrets"
```


### Client Initialization

```go
client := secrets.NewSecretsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SecretsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := secrets.NewSecretID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterName", "secretName")

payload := secrets.Secret{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecretsClient.Delete`

```go
ctx := context.TODO()
id := secrets.NewSecretID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterName", "secretName")

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
id := secrets.NewSecretID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterName", "secretName")

read, err := client.Get(ctx, id)
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
id := secrets.NewOpenShiftClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
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
id := secrets.NewSecretID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterName", "secretName")

payload := secrets.SecretUpdate{
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
