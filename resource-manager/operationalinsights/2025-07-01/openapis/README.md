
## `github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-07-01/openapis` Documentation

The `openapis` SDK allows for interaction with Azure Resource Manager `operationalinsights` (API Version `2025-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-07-01/openapis"
```


### Client Initialization

```go
client := openapis.NewOpenapisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenapisClient.QueriesDelete`

```go
ctx := context.TODO()
id := openapis.NewQueryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackName", "queryName")

read, err := client.QueriesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.QueriesGet`

```go
ctx := context.TODO()
id := openapis.NewQueryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackName", "queryName")

read, err := client.QueriesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.QueriesList`

```go
ctx := context.TODO()
id := openapis.NewQueryPackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackName")

// alternatively `client.QueriesList(ctx, id, openapis.DefaultQueriesListOperationOptions())` can be used to do batched pagination
items, err := client.QueriesListComplete(ctx, id, openapis.DefaultQueriesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.QueriesPut`

```go
ctx := context.TODO()
id := openapis.NewQueryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackName", "queryName")

payload := openapis.LogAnalyticsQueryPackQuery{
	// ...
}


read, err := client.QueriesPut(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.QueriesSearch`

```go
ctx := context.TODO()
id := openapis.NewQueryPackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackName")

payload := openapis.LogAnalyticsQueryPackQuerySearchProperties{
	// ...
}


// alternatively `client.QueriesSearch(ctx, id, payload, openapis.DefaultQueriesSearchOperationOptions())` can be used to do batched pagination
items, err := client.QueriesSearchComplete(ctx, id, payload, openapis.DefaultQueriesSearchOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.QueriesUpdate`

```go
ctx := context.TODO()
id := openapis.NewQueryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackName", "queryName")

payload := openapis.LogAnalyticsQueryPackQuery{
	// ...
}


read, err := client.QueriesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.QueryPacksCreateOrUpdate`

```go
ctx := context.TODO()
id := openapis.NewQueryPackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackName")

payload := openapis.LogAnalyticsQueryPack{
	// ...
}


read, err := client.QueryPacksCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.QueryPacksCreateOrUpdateWithoutName`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

payload := openapis.LogAnalyticsQueryPack{
	// ...
}


read, err := client.QueryPacksCreateOrUpdateWithoutName(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.QueryPacksDelete`

```go
ctx := context.TODO()
id := openapis.NewQueryPackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackName")

read, err := client.QueryPacksDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.QueryPacksGet`

```go
ctx := context.TODO()
id := openapis.NewQueryPackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackName")

read, err := client.QueryPacksGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.QueryPacksList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.QueryPacksList(ctx, id)` can be used to do batched pagination
items, err := client.QueryPacksListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.QueryPacksListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.QueryPacksListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.QueryPacksListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.QueryPacksUpdateTags`

```go
ctx := context.TODO()
id := openapis.NewQueryPackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackName")

payload := openapis.TagsResource{
	// ...
}


read, err := client.QueryPacksUpdateTags(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
