
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/topleveldomains` Documentation

The `topleveldomains` SDK allows for interaction with the Azure Resource Manager Service `web` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/topleveldomains"
```


### Client Initialization

```go
client := topleveldomains.NewTopLevelDomainsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TopLevelDomainsClient.Get`

```go
ctx := context.TODO()
id := topleveldomains.NewTopLevelDomainID("12345678-1234-9876-4563-123456789012", "topLevelDomainValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TopLevelDomainsClient.List`

```go
ctx := context.TODO()
id := topleveldomains.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TopLevelDomainsClient.ListAgreements`

```go
ctx := context.TODO()
id := topleveldomains.NewTopLevelDomainID("12345678-1234-9876-4563-123456789012", "topLevelDomainValue")

payload := topleveldomains.TopLevelDomainAgreementOption{
	// ...
}


// alternatively `client.ListAgreements(ctx, id, payload)` can be used to do batched pagination
items, err := client.ListAgreementsComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
