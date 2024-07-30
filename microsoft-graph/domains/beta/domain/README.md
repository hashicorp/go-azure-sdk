
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/beta/domain` Documentation

The `domain` SDK allows for interaction with the Azure Resource Manager Service `domains` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/beta/domain"
```


### Client Initialization

```go
client := domain.NewDomainClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DomainClient.CreateDomain`

```go
ctx := context.TODO()

payload := domain.Domain{
	// ...
}


read, err := client.CreateDomain(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.CreateForceDelete`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

payload := domain.CreateForceDeleteRequest{
	// ...
}


read, err := client.CreateForceDelete(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.CreatePromote`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

read, err := client.CreatePromote(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.CreatePromoteToInitial`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

read, err := client.CreatePromoteToInitial(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.CreateVerify`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

read, err := client.CreateVerify(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.DeleteDomain`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

read, err := client.DeleteDomain(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.GetDomain`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

read, err := client.GetDomain(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.ListDomains`

```go
ctx := context.TODO()


// alternatively `client.ListDomains(ctx)` can be used to do batched pagination
items, err := client.ListDomainsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DomainClient.UpdateDomain`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

payload := domain.Domain{
	// ...
}


read, err := client.UpdateDomain(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
