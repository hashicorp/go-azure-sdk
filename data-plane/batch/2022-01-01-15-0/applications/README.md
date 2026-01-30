
## `github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/applications` Documentation

The `applications` SDK allows for interaction with <unknown source data type> `batch` (API Version `2022-01-01.15.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/applications"
```


### Client Initialization

```go
client := applications.NewApplicationsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationsClient.ApplicationGet`

```go
ctx := context.TODO()
id := applications.NewApplicationID("applicationId")

read, err := client.ApplicationGet(ctx, id, applications.DefaultApplicationGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationsClient.ApplicationList`

```go
ctx := context.TODO()


// alternatively `client.ApplicationList(ctx, applications.DefaultApplicationListOperationOptions())` can be used to do batched pagination
items, err := client.ApplicationListComplete(ctx, applications.DefaultApplicationListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
