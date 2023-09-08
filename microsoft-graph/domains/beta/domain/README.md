
## `github.com/hashicorp/go-azure-sdk/resource-manager/domains/beta/domain` Documentation

The `domain` SDK allows for interaction with the Azure Resource Manager Service `domains` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/domains/beta/domain"
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


### Example Usage: `DomainClient.CreateDomainByIdForceDelete`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

payload := domain.CreateDomainByIdForceDeleteRequest{
	// ...
}


read, err := client.CreateDomainByIdForceDelete(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.CreateDomainByIdPromote`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

read, err := client.CreateDomainByIdPromote(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.CreateDomainByIdPromoteToInitial`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

read, err := client.CreateDomainByIdPromoteToInitial(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.CreateDomainByIdVerify`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

read, err := client.CreateDomainByIdVerify(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.DeleteDomainById`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

read, err := client.DeleteDomainById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.GetDomainById`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

read, err := client.GetDomainById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.GetDomainCount`

```go
ctx := context.TODO()


read, err := client.GetDomainCount(ctx)
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


### Example Usage: `DomainClient.UpdateDomainById`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainIdValue")

payload := domain.Domain{
	// ...
}


read, err := client.UpdateDomainById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
