
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/netappresource` Documentation

The `netappresource` SDK allows for interaction with the Azure Resource Manager Service `netapp` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/netappresource"
```


### Client Initialization

```go
client := netappresource.NewNetAppResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetAppResourceClient.NetAppResourceCheckFilePathAvailability`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := netappresource.FilePathAvailabilityRequest{
	// ...
}


read, err := client.NetAppResourceCheckFilePathAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.NetAppResourceCheckNameAvailability`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := netappresource.ResourceNameAvailabilityRequest{
	// ...
}


read, err := client.NetAppResourceCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.NetAppResourceCheckQuotaAvailability`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := netappresource.QuotaAvailabilityRequest{
	// ...
}


read, err := client.NetAppResourceCheckQuotaAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.NetAppResourceQueryRegionInfo`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.NetAppResourceQueryRegionInfo(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.NetAppResourceQuotaLimitsGet`

```go
ctx := context.TODO()
id := netappresource.NewQuotaLimitID("12345678-1234-9876-4563-123456789012", "locationValue", "quotaLimitValue")

read, err := client.NetAppResourceQuotaLimitsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetAppResourceClient.NetAppResourceQuotaLimitsList`

```go
ctx := context.TODO()
id := netappresource.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.NetAppResourceQuotaLimitsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
