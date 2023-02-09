
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/changedetection` Documentation

The `changedetection` SDK allows for interaction with the Azure Resource Manager Service `storagesync` (API Version `2020-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/changedetection"
```


### Client Initialization

```go
client := changedetection.NewChangeDetectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChangeDetectionClient.CloudEndpointsTriggerChangeDetection`

```go
ctx := context.TODO()
id := changedetection.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

payload := changedetection.TriggerChangeDetectionParameters{
	// ...
}


if err := client.CloudEndpointsTriggerChangeDetectionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
