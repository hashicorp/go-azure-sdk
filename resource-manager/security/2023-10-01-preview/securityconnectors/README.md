
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-10-01-preview/securityconnectors` Documentation

The `securityconnectors` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2023-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-10-01-preview/securityconnectors"
```


### Client Initialization

```go
client := securityconnectors.NewSecurityConnectorsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SecurityConnectorsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := securityconnectors.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorValue")

payload := securityconnectors.SecurityConnector{
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


### Example Usage: `SecurityConnectorsClient.Delete`

```go
ctx := context.TODO()
id := securityconnectors.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityConnectorsClient.Get`

```go
ctx := context.TODO()
id := securityconnectors.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityConnectorsClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SecurityConnectorsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SecurityConnectorsClient.Update`

```go
ctx := context.TODO()
id := securityconnectors.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorValue")

payload := securityconnectors.SecurityConnector{
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
