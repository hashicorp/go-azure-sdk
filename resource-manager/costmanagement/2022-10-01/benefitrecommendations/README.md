
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/benefitrecommendations` Documentation

The `benefitrecommendations` SDK allows for interaction with Azure Resource Manager `costmanagement` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-10-01/benefitrecommendations"
```


### Client Initialization

```go
client := benefitrecommendations.NewBenefitRecommendationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BenefitRecommendationsClient.List`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.List(ctx, id, benefitrecommendations.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, benefitrecommendations.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
