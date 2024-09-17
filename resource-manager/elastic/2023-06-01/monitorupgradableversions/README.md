
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/monitorupgradableversions` Documentation

The `monitorupgradableversions` SDK allows for interaction with Azure Resource Manager `elastic` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/monitorupgradableversions"
```


### Client Initialization

```go
client := monitorupgradableversions.NewMonitorUpgradableVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MonitorUpgradableVersionsClient.UpgradableVersionsDetails`

```go
ctx := context.TODO()
id := monitorupgradableversions.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorValue")

read, err := client.UpgradableVersionsDetails(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
