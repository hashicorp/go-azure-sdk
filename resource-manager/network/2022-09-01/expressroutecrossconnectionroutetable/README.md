
## `github.com/hashicorp/go-azure-sdk/resource-manager/network/2022-09-01/expressroutecrossconnectionroutetable` Documentation

The `expressroutecrossconnectionroutetable` SDK allows for interaction with the Azure Resource Manager Service `network` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/network/2022-09-01/expressroutecrossconnectionroutetable"
```


### Client Initialization

```go
client := expressroutecrossconnectionroutetable.NewExpressRouteCrossConnectionRouteTableClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExpressRouteCrossConnectionRouteTableClient.ExpressRouteCrossConnectionsListRoutesTable`

```go
ctx := context.TODO()
id := expressroutecrossconnectionroutetable.NewRouteTableID("12345678-1234-9876-4563-123456789012", "example-resource-group", "routeTableValue")

if err := client.ExpressRouteCrossConnectionsListRoutesTableThenPoll(ctx, id); err != nil {
	// handle the error
}
```
