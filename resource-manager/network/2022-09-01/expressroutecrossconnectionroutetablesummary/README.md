
## `github.com/hashicorp/go-azure-sdk/resource-manager/network/2022-09-01/expressroutecrossconnectionroutetablesummary` Documentation

The `expressroutecrossconnectionroutetablesummary` SDK allows for interaction with the Azure Resource Manager Service `network` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/network/2022-09-01/expressroutecrossconnectionroutetablesummary"
```


### Client Initialization

```go
client := expressroutecrossconnectionroutetablesummary.NewExpressRouteCrossConnectionRouteTableSummaryClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExpressRouteCrossConnectionRouteTableSummaryClient.ExpressRouteCrossConnectionsListRoutesTableSummary`

```go
ctx := context.TODO()
id := expressroutecrossconnectionroutetablesummary.NewRouteTablesSummaryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "expressRouteCircuitValue", "peeringValue", "routeTablesSummaryValue")

if err := client.ExpressRouteCrossConnectionsListRoutesTableSummaryThenPoll(ctx, id); err != nil {
	// handle the error
}
```
