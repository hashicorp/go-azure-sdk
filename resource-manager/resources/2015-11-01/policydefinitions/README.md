
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/policydefinitions` Documentation

The `policydefinitions` SDK allows for interaction with Azure Resource Manager `resources` (API Version `2015-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/policydefinitions"
```


### Client Initialization

```go
client := policydefinitions.NewPolicyDefinitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyDefinitionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := policydefinitions.NewPolicydefinitionID("12345678-1234-9876-4563-123456789012", "policydefinitionValue")

payload := policydefinitions.PolicyDefinition{
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


### Example Usage: `PolicyDefinitionsClient.Delete`

```go
ctx := context.TODO()
id := policydefinitions.NewPolicydefinitionID("12345678-1234-9876-4563-123456789012", "policydefinitionValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyDefinitionsClient.Get`

```go
ctx := context.TODO()
id := policydefinitions.NewPolicydefinitionID("12345678-1234-9876-4563-123456789012", "policydefinitionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
