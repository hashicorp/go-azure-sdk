
## `github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-02-01/managedhsmkeys` Documentation

The `managedhsmkeys` SDK allows for interaction with Azure Resource Manager `keyvault` (API Version `2023-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-02-01/managedhsmkeys"
```


### Client Initialization

```go
client := managedhsmkeys.NewManagedHsmKeysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedHsmKeysClient.CreateIfNotExist`

```go
ctx := context.TODO()
id := managedhsmkeys.NewKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMName", "keyName")

payload := managedhsmkeys.ManagedHsmKeyCreateParameters{
	// ...
}


read, err := client.CreateIfNotExist(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedHsmKeysClient.Get`

```go
ctx := context.TODO()
id := managedhsmkeys.NewKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMName", "keyName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedHsmKeysClient.GetVersion`

```go
ctx := context.TODO()
id := managedhsmkeys.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMName", "keyName", "versionName")

read, err := client.GetVersion(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedHsmKeysClient.List`

```go
ctx := context.TODO()
id := managedhsmkeys.NewManagedHSMID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedHsmKeysClient.ListVersions`

```go
ctx := context.TODO()
id := managedhsmkeys.NewKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMName", "keyName")

// alternatively `client.ListVersions(ctx, id)` can be used to do batched pagination
items, err := client.ListVersionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
