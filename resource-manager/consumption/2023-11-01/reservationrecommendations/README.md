
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2023-11-01/reservationrecommendations` Documentation

The `reservationrecommendations` SDK allows for interaction with the Azure Resource Manager Service `consumption` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2023-11-01/reservationrecommendations"
```


### Client Initialization

```go
client := reservationrecommendations.NewReservationRecommendationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReservationRecommendationsClient.List`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.List(ctx, id, reservationrecommendations.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, reservationrecommendations.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
