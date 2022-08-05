
## `github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/recommendedactions` Documentation

The `recommendedactions` SDK allows for interaction with the Azure Resource Manager Service `mariadb` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/recommendedactions"
```


### Client Initialization

```go
client := recommendedactions.NewRecommendedActionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecommendedActionsClient.Get`

```go
ctx := context.TODO()
id := recommendedactions.NewRecommendedActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "advisorValue", "recommendedActionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendedActionsClient.ListByServer`

```go
ctx := context.TODO()
id := recommendedactions.NewAdvisorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "advisorValue")

// alternatively `client.ListByServer(ctx, id, recommendedactions.DefaultListByServerOperationOptions())` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id, recommendedactions.DefaultListByServerOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
