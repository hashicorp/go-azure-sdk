
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/netappresource` Documentation

The `netappresource` SDK allows for interaction with the Azure Resource Manager Service `netapp` (API Version `2022-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/netappresource"
```


### Client Initialization

```go
client := netappresource.NewNetAppResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetAppResourceClient.CheckFilePathAvailability`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := netappresource.FilePathAvailabilityRequest{
	// ...
}


read, err := client.CheckFilePathAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := netappresource.ResourceNameAvailabilityRequest{
	// ...
}


read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.CheckQuotaAvailability`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := netappresource.QuotaAvailabilityRequest{
	// ...
}


read, err := client.CheckQuotaAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.QueryRegionInfo`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.QueryRegionInfo(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.QuotaLimitsGet`

```go
ctx := context.TODO()
id := netappresource.NewQuotaLimitID("12345678-1234-9876-4563-123456789012", "locationValue", "quotaLimitValue")

read, err := client.QuotaLimitsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.QuotaLimitsList`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.QuotaLimitsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
