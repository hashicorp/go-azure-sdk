
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/embeddedsimactivationcodepool` Documentation

The `embeddedsimactivationcodepool` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/embeddedsimactivationcodepool"
```


### Client Initialization

```go
client := embeddedsimactivationcodepool.NewEmbeddedSIMActivationCodePoolClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EmbeddedSIMActivationCodePoolClient.AssignEmbeddedSIMActivationCodePools`

```go
ctx := context.TODO()
id := embeddedsimactivationcodepool.NewDeviceManagementEmbeddedSIMActivationCodePoolID("embeddedSIMActivationCodePoolIdValue")

payload := embeddedsimactivationcodepool.AssignEmbeddedSIMActivationCodePoolsRequest{
	// ...
}


// alternatively `client.AssignEmbeddedSIMActivationCodePools(ctx, id, payload, embeddedsimactivationcodepool.DefaultAssignEmbeddedSIMActivationCodePoolsOperationOptions())` can be used to do batched pagination
items, err := client.AssignEmbeddedSIMActivationCodePoolsComplete(ctx, id, payload, embeddedsimactivationcodepool.DefaultAssignEmbeddedSIMActivationCodePoolsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EmbeddedSIMActivationCodePoolClient.CreateEmbeddedSIMActivationCodePool`

```go
ctx := context.TODO()

payload := embeddedsimactivationcodepool.EmbeddedSIMActivationCodePool{
	// ...
}


read, err := client.CreateEmbeddedSIMActivationCodePool(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EmbeddedSIMActivationCodePoolClient.DeleteEmbeddedSIMActivationCodePool`

```go
ctx := context.TODO()
id := embeddedsimactivationcodepool.NewDeviceManagementEmbeddedSIMActivationCodePoolID("embeddedSIMActivationCodePoolIdValue")

read, err := client.DeleteEmbeddedSIMActivationCodePool(ctx, id, embeddedsimactivationcodepool.DefaultDeleteEmbeddedSIMActivationCodePoolOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EmbeddedSIMActivationCodePoolClient.GetEmbeddedSIMActivationCodePool`

```go
ctx := context.TODO()
id := embeddedsimactivationcodepool.NewDeviceManagementEmbeddedSIMActivationCodePoolID("embeddedSIMActivationCodePoolIdValue")

read, err := client.GetEmbeddedSIMActivationCodePool(ctx, id, embeddedsimactivationcodepool.DefaultGetEmbeddedSIMActivationCodePoolOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EmbeddedSIMActivationCodePoolClient.GetEmbeddedSIMActivationCodePoolsCount`

```go
ctx := context.TODO()


read, err := client.GetEmbeddedSIMActivationCodePoolsCount(ctx, embeddedsimactivationcodepool.DefaultGetEmbeddedSIMActivationCodePoolsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EmbeddedSIMActivationCodePoolClient.ListEmbeddedSIMActivationCodePools`

```go
ctx := context.TODO()


// alternatively `client.ListEmbeddedSIMActivationCodePools(ctx, embeddedsimactivationcodepool.DefaultListEmbeddedSIMActivationCodePoolsOperationOptions())` can be used to do batched pagination
items, err := client.ListEmbeddedSIMActivationCodePoolsComplete(ctx, embeddedsimactivationcodepool.DefaultListEmbeddedSIMActivationCodePoolsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EmbeddedSIMActivationCodePoolClient.UpdateEmbeddedSIMActivationCodePool`

```go
ctx := context.TODO()
id := embeddedsimactivationcodepool.NewDeviceManagementEmbeddedSIMActivationCodePoolID("embeddedSIMActivationCodePoolIdValue")

payload := embeddedsimactivationcodepool.EmbeddedSIMActivationCodePool{
	// ...
}


read, err := client.UpdateEmbeddedSIMActivationCodePool(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
