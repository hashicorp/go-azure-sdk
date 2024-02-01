
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2023-05-01/aggregatedcost` Documentation

The `aggregatedcost` SDK allows for interaction with the Azure Resource Manager Service `consumption` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2023-05-01/aggregatedcost"
```


### Client Initialization

```go
client := aggregatedcost.NewAggregatedCostClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AggregatedCostClient.GetByManagementGroup`

```go
ctx := context.TODO()
id := commonids.NewManagementGroupID("groupIdValue")

read, err := client.GetByManagementGroup(ctx, id, aggregatedcost.DefaultGetByManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AggregatedCostClient.GetForBillingPeriodByManagementGroup`

```go
ctx := context.TODO()
id := aggregatedcost.NewProviders2BillingPeriodID("managementGroupIdValue", "billingPeriodValue")

read, err := client.GetForBillingPeriodByManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
