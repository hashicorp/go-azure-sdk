
## `github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/subscriptions` Documentation

The `subscriptions` SDK allows for interaction with the Azure Resource Manager Service `streamanalytics` (API Version `2020-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2020-03-01/subscriptions"
```


### Client Initialization

```go
client := subscriptions.NewSubscriptionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SubscriptionsClient.ListQuotas`

```go
ctx := context.TODO()
id := subscriptions.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.ListQuotas(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
