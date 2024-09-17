
## `github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/consumerinvitation` Documentation

The `consumerinvitation` SDK allows for interaction with Azure Resource Manager `datashare` (API Version `2021-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/consumerinvitation"
```


### Client Initialization

```go
client := consumerinvitation.NewConsumerInvitationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConsumerInvitationClient.Get`

```go
ctx := context.TODO()
id := consumerinvitation.NewConsumerInvitationID("locationValue", "invitationIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConsumerInvitationClient.ListInvitations`

```go
ctx := context.TODO()


// alternatively `client.ListInvitations(ctx)` can be used to do batched pagination
items, err := client.ListInvitationsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConsumerInvitationClient.RejectInvitation`

```go
ctx := context.TODO()
id := consumerinvitation.NewLocationID("locationValue")

payload := consumerinvitation.ConsumerInvitation{
	// ...
}


read, err := client.RejectInvitation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
