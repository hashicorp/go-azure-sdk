
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2016-06-01/customapis` Documentation

The `customapis` SDK allows for interaction with Azure Resource Manager `web` (API Version `2016-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2016-06-01/customapis"
```


### Client Initialization

```go
client := customapis.NewCustomAPIsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CustomAPIsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := customapis.NewCustomApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "apiName")

payload := customapis.CustomApiDefinition{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.Delete`

```go
ctx := context.TODO()
id := customapis.NewCustomApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "apiName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.ExtractApiDefinitionFromWsdl`

```go
ctx := context.TODO()
id := customapis.NewLocationID("12345678-1234-9876-4563-123456789012", "location")

payload := customapis.WsdlDefinition{
	// ...
}


read, err := client.ExtractApiDefinitionFromWsdl(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.Get`

```go
ctx := context.TODO()
id := customapis.NewCustomApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "apiName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.List(ctx, id, customapis.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.ListByResourceGroup(ctx, id, customapis.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.ListWsdlInterfaces`

```go
ctx := context.TODO()
id := customapis.NewLocationID("12345678-1234-9876-4563-123456789012", "location")

payload := customapis.WsdlDefinition{
	// ...
}


read, err := client.ListWsdlInterfaces(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.Move`

```go
ctx := context.TODO()
id := customapis.NewCustomApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "apiName")

payload := customapis.CustomApiReference{
	// ...
}


read, err := client.Move(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.Update`

```go
ctx := context.TODO()
id := customapis.NewCustomApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "apiName")

payload := customapis.CustomApiDefinition{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
