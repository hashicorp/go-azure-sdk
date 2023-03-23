
## `github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/keyvalues` Documentation

The `keyvalues` SDK allows for interaction with the Azure Resource Manager Service `appconfiguration` (API Version `2023-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/keyvalues"
```


### Client Initialization

```go
client := keyvalues.NewKeyValuesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `KeyValuesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := keyvalues.NewKeyValueID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configurationStoreValue", "keyValueValue")

payload := keyvalues.KeyValue{
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


### Example Usage: `KeyValuesClient.Delete`

```go
ctx := context.TODO()
id := keyvalues.NewKeyValueID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configurationStoreValue", "keyValueValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `KeyValuesClient.Get`

```go
ctx := context.TODO()
id := keyvalues.NewKeyValueID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configurationStoreValue", "keyValueValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
