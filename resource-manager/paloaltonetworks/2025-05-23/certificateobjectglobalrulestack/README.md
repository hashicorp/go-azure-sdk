
## `github.com/hashicorp/go-azure-sdk/resource-manager/paloaltonetworks/2025-05-23/certificateobjectglobalrulestack` Documentation

The `certificateobjectglobalrulestack` SDK allows for interaction with Azure Resource Manager `paloaltonetworks` (API Version `2025-05-23`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/paloaltonetworks/2025-05-23/certificateobjectglobalrulestack"
```


### Client Initialization

```go
client := certificateobjectglobalrulestack.NewCertificateObjectGlobalRulestackClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CertificateObjectGlobalRulestackClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := certificateobjectglobalrulestack.NewCertificateID("globalRulestackName", "certificateName")

payload := certificateobjectglobalrulestack.CertificateObjectGlobalRulestackResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CertificateObjectGlobalRulestackClient.Delete`

```go
ctx := context.TODO()
id := certificateobjectglobalrulestack.NewCertificateID("globalRulestackName", "certificateName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `CertificateObjectGlobalRulestackClient.Get`

```go
ctx := context.TODO()
id := certificateobjectglobalrulestack.NewCertificateID("globalRulestackName", "certificateName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateObjectGlobalRulestackClient.List`

```go
ctx := context.TODO()
id := certificateobjectglobalrulestack.NewGlobalRulestackID("globalRulestackName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
