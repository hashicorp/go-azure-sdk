
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/connectedresources` Documentation

The `connectedresources` SDK allows for interaction with Azure Resource Manager `elastic` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/connectedresources"
```


### Client Initialization

```go
client := connectedresources.NewConnectedResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConnectedResourcesClient.BillingInfoGet`

```go
ctx := context.TODO()
id := connectedresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

read, err := client.BillingInfoGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectedResourcesClient.ConnectedPartnerResourcesList`

```go
ctx := context.TODO()
id := connectedresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

// alternatively `client.ConnectedPartnerResourcesList(ctx, id)` can be used to do batched pagination
items, err := client.ConnectedPartnerResourcesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
