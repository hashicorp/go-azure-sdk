
## `github.com/hashicorp/go-azure-sdk/resource-manager/oracledatabase/2025-03-01/dbsystemshapes` Documentation

The `dbsystemshapes` SDK allows for interaction with Azure Resource Manager `oracledatabase` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/oracledatabase/2025-03-01/dbsystemshapes"
```


### Client Initialization

```go
client := dbsystemshapes.NewDbSystemShapesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DbSystemShapesClient.Get`

```go
ctx := context.TODO()
id := dbsystemshapes.NewDbSystemShapeID("12345678-1234-9876-4563-123456789012", "locationName", "dbSystemShapeName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DbSystemShapesClient.ListByLocation`

```go
ctx := context.TODO()
id := dbsystemshapes.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.ListByLocation(ctx, id, dbsystemshapes.DefaultListByLocationOperationOptions())` can be used to do batched pagination
items, err := client.ListByLocationComplete(ctx, id, dbsystemshapes.DefaultListByLocationOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
