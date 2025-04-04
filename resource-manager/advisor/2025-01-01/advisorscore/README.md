
## `github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2025-01-01/advisorscore` Documentation

The `advisorscore` SDK allows for interaction with Azure Resource Manager `advisor` (API Version `2025-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2025-01-01/advisorscore"
```


### Client Initialization

```go
client := advisorscore.NewAdvisorScoreClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AdvisorScoreClient.Get`

```go
ctx := context.TODO()
id := advisorscore.NewAdvisorScoreID("12345678-1234-9876-4563-123456789012", "advisorScoreName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdvisorScoreClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
