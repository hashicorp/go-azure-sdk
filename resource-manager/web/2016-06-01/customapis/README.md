
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2016-06-01/customapis` Documentation

The `customapis` SDK allows for interaction with the Azure Resource Manager Service `web` (API Version `2016-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2016-06-01/customapis"
```


### Client Initialization

```go
client := customapis.NewCustomAPIsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `CustomAPIsClient.CustomApisCreateOrUpdate`

```go
ctx := context.TODO()
id := customapis.NewCustomApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "apiValue")

payload := customapis.CustomApiDefinition{
	// ...
}

read, err := client.CustomApisCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.CustomApisDelete`

```go
ctx := context.TODO()
id := customapis.NewCustomApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "apiValue")
read, err := client.CustomApisDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.CustomApisExtractApiDefinitionFromWsdl`

```go
ctx := context.TODO()
id := customapis.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := customapis.WsdlDefinition{
	// ...
}

read, err := client.CustomApisExtractApiDefinitionFromWsdl(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.CustomApisGet`

```go
ctx := context.TODO()
id := customapis.NewCustomApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "apiValue")
read, err := client.CustomApisGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.CustomApisList`

```go
ctx := context.TODO()
id := customapis.NewSubscriptionID()
read, err := client.CustomApisList(ctx, id, customapis.DefaultCustomApisListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.CustomApisListByResourceGroup`

```go
ctx := context.TODO()
id := customapis.NewResourceGroupID()
read, err := client.CustomApisListByResourceGroup(ctx, id, customapis.DefaultCustomApisListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.CustomApisListWsdlInterfaces`

```go
ctx := context.TODO()
id := customapis.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := customapis.WsdlDefinition{
	// ...
}

read, err := client.CustomApisListWsdlInterfaces(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.CustomApisMove`

```go
ctx := context.TODO()
id := customapis.NewCustomApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "apiValue")

payload := customapis.CustomApiReference{
	// ...
}

read, err := client.CustomApisMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAPIsClient.CustomApisUpdate`

```go
ctx := context.TODO()
id := customapis.NewCustomApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "apiValue")

payload := customapis.CustomApiDefinition{
	// ...
}

read, err := client.CustomApisUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
