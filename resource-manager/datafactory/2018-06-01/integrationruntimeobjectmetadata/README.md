
## `github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/integrationruntimeobjectmetadata` Documentation

The `integrationruntimeobjectmetadata` SDK allows for interaction with the Azure Resource Manager Service `datafactory` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/integrationruntimeobjectmetadata"
```


### Client Initialization

```go
client := integrationruntimeobjectmetadata.NewIntegrationRuntimeObjectMetadataClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IntegrationRuntimeObjectMetadataClient.Get`

```go
ctx := context.TODO()
id := integrationruntimeobjectmetadata.NewIntegrationRuntimeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "integrationRuntimeValue")

payload := integrationruntimeobjectmetadata.GetSsisObjectMetadataRequest{
	// ...
}


// alternatively `client.Get(ctx, id, payload)` can be used to do batched pagination
items, err := client.GetComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IntegrationRuntimeObjectMetadataClient.Refresh`

```go
ctx := context.TODO()
id := integrationruntimeobjectmetadata.NewIntegrationRuntimeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "integrationRuntimeValue")

if err := client.RefreshThenPoll(ctx, id); err != nil {
	// handle the error
}
```
