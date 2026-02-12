
## `github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1.0/keyvalues` Documentation

The `keyvalues` SDK allows for interaction with <unknown source data type> `appconfiguration` (API Version `1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1.0/keyvalues"
```


### Client Initialization

```go
client := keyvalues.NewKeyValuesClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `KeyValuesClient.CheckKeyValue`

```go
ctx := context.TODO()
id := keyvalues.NewKvID("kvName")

read, err := client.CheckKeyValue(ctx, id, keyvalues.DefaultCheckKeyValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeyValuesClient.CheckKeyValues`

```go
ctx := context.TODO()


read, err := client.CheckKeyValues(ctx, keyvalues.DefaultCheckKeyValuesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeyValuesClient.DeleteKeyValue`

```go
ctx := context.TODO()
id := keyvalues.NewKvID("kvName")

read, err := client.DeleteKeyValue(ctx, id, keyvalues.DefaultDeleteKeyValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeyValuesClient.GetKeyValue`

```go
ctx := context.TODO()
id := keyvalues.NewKvID("kvName")

read, err := client.GetKeyValue(ctx, id, keyvalues.DefaultGetKeyValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeyValuesClient.GetKeyValues`

```go
ctx := context.TODO()


// alternatively `client.GetKeyValues(ctx, keyvalues.DefaultGetKeyValuesOperationOptions())` can be used to do batched pagination
items, err := client.GetKeyValuesComplete(ctx, keyvalues.DefaultGetKeyValuesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `KeyValuesClient.PutKeyValue`

```go
ctx := context.TODO()
id := keyvalues.NewKvID("kvName")

payload := keyvalues.KeyValue{
	// ...
}


read, err := client.PutKeyValue(ctx, id, payload, keyvalues.DefaultPutKeyValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
