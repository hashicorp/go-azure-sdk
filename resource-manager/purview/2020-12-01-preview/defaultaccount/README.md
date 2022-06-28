
## `github.com/hashicorp/go-azure-sdk/resource-manager/purview/2020-12-01-preview/defaultaccount` Documentation

The `defaultaccount` SDK allows for interaction with the Azure Resource Manager Service `purview` (API Version `2020-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/purview/2020-12-01-preview/defaultaccount"
```


### Client Initialization

```go
client := defaultaccount.NewDefaultAccountClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DefaultAccountClient.Get`

```go
ctx := context.TODO()


read, err := client.Get(ctx, defaultaccount.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DefaultAccountClient.Remove`

```go
ctx := context.TODO()


read, err := client.Remove(ctx, defaultaccount.DefaultRemoveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DefaultAccountClient.Set`

```go
ctx := context.TODO()

payload := defaultaccount.DefaultAccountPayload{
	// ...
}


read, err := client.Set(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
