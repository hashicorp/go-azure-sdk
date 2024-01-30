
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-05-01-preview/graphqlapiresolver` Documentation

The `graphqlapiresolver` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2023-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-05-01-preview/graphqlapiresolver"
```


### Client Initialization

```go
client := graphqlapiresolver.NewGraphQLApiResolverClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GraphQLApiResolverClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := graphqlapiresolver.NewResolverID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue", "resolverIdValue")

payload := graphqlapiresolver.ResolverContract{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, graphqlapiresolver.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GraphQLApiResolverClient.Delete`

```go
ctx := context.TODO()
id := graphqlapiresolver.NewResolverID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue", "resolverIdValue")

read, err := client.Delete(ctx, id, graphqlapiresolver.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GraphQLApiResolverClient.Get`

```go
ctx := context.TODO()
id := graphqlapiresolver.NewResolverID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue", "resolverIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GraphQLApiResolverClient.GetEntityTag`

```go
ctx := context.TODO()
id := graphqlapiresolver.NewResolverID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue", "resolverIdValue")

read, err := client.GetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GraphQLApiResolverClient.ListByApi`

```go
ctx := context.TODO()
id := graphqlapiresolver.NewApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue")

// alternatively `client.ListByApi(ctx, id, graphqlapiresolver.DefaultListByApiOperationOptions())` can be used to do batched pagination
items, err := client.ListByApiComplete(ctx, id, graphqlapiresolver.DefaultListByApiOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GraphQLApiResolverClient.Update`

```go
ctx := context.TODO()
id := graphqlapiresolver.NewResolverID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue", "resolverIdValue")

payload := graphqlapiresolver.ResolverUpdateContract{
	// ...
}


read, err := client.Update(ctx, id, payload, graphqlapiresolver.DefaultUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
