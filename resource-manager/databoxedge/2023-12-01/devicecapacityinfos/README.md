
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/devicecapacityinfos` Documentation

The `devicecapacityinfos` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/devicecapacityinfos"
```


### Client Initialization

```go
client := devicecapacityinfos.NewDeviceCapacityInfosClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceCapacityInfosClient.DeviceCapacityInfoGetDeviceCapacityInfo`

```go
ctx := context.TODO()
id := devicecapacityinfos.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

read, err := client.DeviceCapacityInfoGetDeviceCapacityInfo(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
