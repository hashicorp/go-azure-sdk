
## `github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1.0/locks` Documentation

The `locks` SDK allows for interaction with <unknown source data type> `appconfiguration` (API Version `1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1.0/locks"
```


### Client Initialization

```go
client := locks.NewLocksClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `LocksClient.DeleteLock`

```go
ctx := context.TODO()
id := locks.NewLockID("lockName")

read, err := client.DeleteLock(ctx, id, locks.DefaultDeleteLockOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LocksClient.PutLock`

```go
ctx := context.TODO()
id := locks.NewLockID("lockName")

read, err := client.PutLock(ctx, id, locks.DefaultPutLockOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
