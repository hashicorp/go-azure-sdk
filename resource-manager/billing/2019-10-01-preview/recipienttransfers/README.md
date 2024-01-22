
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/recipienttransfers` Documentation

The `recipienttransfers` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/recipienttransfers"
```


### Client Initialization

```go
client := recipienttransfers.NewRecipientTransfersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecipientTransfersClient.Accept`

```go
ctx := context.TODO()
id := recipienttransfers.NewTransferID("transferValue")

payload := recipienttransfers.AcceptTransferRequest{
	// ...
}


read, err := client.Accept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecipientTransfersClient.Decline`

```go
ctx := context.TODO()
id := recipienttransfers.NewTransferID("transferValue")

read, err := client.Decline(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecipientTransfersClient.Get`

```go
ctx := context.TODO()
id := recipienttransfers.NewTransferID("transferValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecipientTransfersClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx)` can be used to do batched pagination
items, err := client.ListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RecipientTransfersClient.Validate`

```go
ctx := context.TODO()
id := recipienttransfers.NewTransferID("transferValue")

payload := recipienttransfers.AcceptTransferRequest{
	// ...
}


read, err := client.Validate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
