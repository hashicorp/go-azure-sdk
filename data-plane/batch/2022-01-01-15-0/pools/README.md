
## `github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/pools` Documentation

The `pools` SDK allows for interaction with <unknown source data type> `batch` (API Version `2022-01-01.15.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/pools"
```


### Client Initialization

```go
client := pools.NewPoolsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `PoolsClient.PoolAdd`

```go
ctx := context.TODO()

payload := pools.PoolAddParameter{
	// ...
}


read, err := client.PoolAdd(ctx, payload, pools.DefaultPoolAddOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoolsClient.PoolDelete`

```go
ctx := context.TODO()
id := pools.NewPoolID("poolId")

read, err := client.PoolDelete(ctx, id, pools.DefaultPoolDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoolsClient.PoolDisableAutoScale`

```go
ctx := context.TODO()
id := pools.NewPoolID("poolId")

read, err := client.PoolDisableAutoScale(ctx, id, pools.DefaultPoolDisableAutoScaleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoolsClient.PoolEnableAutoScale`

```go
ctx := context.TODO()
id := pools.NewPoolID("poolId")

payload := pools.PoolEnableAutoScaleParameter{
	// ...
}


read, err := client.PoolEnableAutoScale(ctx, id, payload, pools.DefaultPoolEnableAutoScaleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoolsClient.PoolEvaluateAutoScale`

```go
ctx := context.TODO()
id := pools.NewPoolID("poolId")

payload := pools.PoolEvaluateAutoScaleParameter{
	// ...
}


read, err := client.PoolEvaluateAutoScale(ctx, id, payload, pools.DefaultPoolEvaluateAutoScaleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoolsClient.PoolExists`

```go
ctx := context.TODO()
id := pools.NewPoolID("poolId")

read, err := client.PoolExists(ctx, id, pools.DefaultPoolExistsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoolsClient.PoolGet`

```go
ctx := context.TODO()
id := pools.NewPoolID("poolId")

read, err := client.PoolGet(ctx, id, pools.DefaultPoolGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoolsClient.PoolGetAllLifetimeStatistics`

```go
ctx := context.TODO()


read, err := client.PoolGetAllLifetimeStatistics(ctx, pools.DefaultPoolGetAllLifetimeStatisticsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoolsClient.PoolList`

```go
ctx := context.TODO()


// alternatively `client.PoolList(ctx, pools.DefaultPoolListOperationOptions())` can be used to do batched pagination
items, err := client.PoolListComplete(ctx, pools.DefaultPoolListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PoolsClient.PoolListUsageMetrics`

```go
ctx := context.TODO()


// alternatively `client.PoolListUsageMetrics(ctx, pools.DefaultPoolListUsageMetricsOperationOptions())` can be used to do batched pagination
items, err := client.PoolListUsageMetricsComplete(ctx, pools.DefaultPoolListUsageMetricsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PoolsClient.PoolPatch`

```go
ctx := context.TODO()
id := pools.NewPoolID("poolId")

payload := pools.PoolPatchParameter{
	// ...
}


read, err := client.PoolPatch(ctx, id, payload, pools.DefaultPoolPatchOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoolsClient.PoolResize`

```go
ctx := context.TODO()
id := pools.NewPoolID("poolId")

payload := pools.PoolResizeParameter{
	// ...
}


read, err := client.PoolResize(ctx, id, payload, pools.DefaultPoolResizeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoolsClient.PoolStopResize`

```go
ctx := context.TODO()
id := pools.NewPoolID("poolId")

read, err := client.PoolStopResize(ctx, id, pools.DefaultPoolStopResizeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoolsClient.PoolUpdateProperties`

```go
ctx := context.TODO()
id := pools.NewPoolID("poolId")

payload := pools.PoolUpdatePropertiesParameter{
	// ...
}


read, err := client.PoolUpdateProperties(ctx, id, payload, pools.DefaultPoolUpdatePropertiesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
