
## `github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/keyvalues` Documentation

The `keyvalues` SDK allows for interaction with the Azure Resource Manager Service `appconfiguration` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/keyvalues"
```


### Client Initialization

```go
client := keyvalues.NewKeyValuesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `KeyValuesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := keyvalues.NewKeyValueID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configStoreValue", "keyValueValue")

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
id := keyvalues.NewKeyValueID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configStoreValue", "keyValueValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `KeyValuesClient.Get`

```go
ctx := context.TODO()
id := keyvalues.NewKeyValueID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configStoreValue", "keyValueValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeyValuesClient.ListByConfigurationStore`

```go
ctx := context.TODO()
id := keyvalues.NewConfigurationStoreID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configStoreValue")
// alternatively `client.ListByConfigurationStore(ctx, id)` can be used to do batched pagination
items, err := client.ListByConfigurationStoreComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
