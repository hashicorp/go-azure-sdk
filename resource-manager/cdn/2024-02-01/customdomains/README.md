
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2024-02-01/customdomains` Documentation

The `customdomains` SDK allows for interaction with the Azure Resource Manager Service `cdn` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2024-02-01/customdomains"
```


### Client Initialization

```go
client := customdomains.NewCustomDomainsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CustomDomainsClient.Create`

```go
ctx := context.TODO()
id := customdomains.NewEndpointCustomDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileValue", "endpointValue", "customDomainValue")

payload := customdomains.CustomDomainParameters{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CustomDomainsClient.Delete`

```go
ctx := context.TODO()
id := customdomains.NewEndpointCustomDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileValue", "endpointValue", "customDomainValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `CustomDomainsClient.DisableCustomHTTPS`

```go
ctx := context.TODO()
id := customdomains.NewEndpointCustomDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileValue", "endpointValue", "customDomainValue")

if err := client.DisableCustomHTTPSThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `CustomDomainsClient.EnableCustomHTTPS`

```go
ctx := context.TODO()
id := customdomains.NewEndpointCustomDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileValue", "endpointValue", "customDomainValue")

payload := customdomains.CustomDomainHTTPSParameters{
	// ...
}


if err := client.EnableCustomHTTPSThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CustomDomainsClient.Get`

```go
ctx := context.TODO()
id := customdomains.NewEndpointCustomDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileValue", "endpointValue", "customDomainValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomDomainsClient.ListByEndpoint`

```go
ctx := context.TODO()
id := customdomains.NewEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileValue", "endpointValue")

// alternatively `client.ListByEndpoint(ctx, id)` can be used to do batched pagination
items, err := client.ListByEndpointComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
