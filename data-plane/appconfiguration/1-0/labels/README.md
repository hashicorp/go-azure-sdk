
## `github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1.0/labels` Documentation

The `labels` SDK allows for interaction with <unknown source data type> `appconfiguration` (API Version `1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1.0/labels"
```


### Client Initialization

```go
client := labels.NewLabelsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `LabelsClient.CheckLabels`

```go
ctx := context.TODO()


read, err := client.CheckLabels(ctx, labels.DefaultCheckLabelsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LabelsClient.GetLabels`

```go
ctx := context.TODO()


// alternatively `client.GetLabels(ctx, labels.DefaultGetLabelsOperationOptions())` can be used to do batched pagination
items, err := client.GetLabelsComplete(ctx, labels.DefaultGetLabelsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
