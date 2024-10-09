
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/devicecapacitycheck` Documentation

The `devicecapacitycheck` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2023-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/devicecapacitycheck"
```


### Client Initialization

```go
client := devicecapacitycheck.NewDeviceCapacityCheckClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceCapacityCheckClient.CheckResourceCreationFeasibility`

```go
ctx := context.TODO()
id := devicecapacitycheck.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

payload := devicecapacitycheck.DeviceCapacityRequestInfo{
	// ...
}


if err := client.CheckResourceCreationFeasibilityThenPoll(ctx, id, payload, devicecapacitycheck.DefaultCheckResourceCreationFeasibilityOperationOptions()); err != nil {
	// handle the error
}
```
