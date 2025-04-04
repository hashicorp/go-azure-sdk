
## `github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/networkconfigurations` Documentation

The `networkconfigurations` SDK allows for interaction with Azure Resource Manager `hybridcompute` (API Version `2025-01-13`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/networkconfigurations"
```


### Client Initialization

```go
client := networkconfigurations.NewNetworkConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetworkConfigurationsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := networkconfigurations.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineName")

payload := networkconfigurations.NetworkConfiguration{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkConfigurationsClient.Get`

```go
ctx := context.TODO()
id := networkconfigurations.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkConfigurationsClient.Update`

```go
ctx := context.TODO()
id := networkconfigurations.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineName")

payload := networkconfigurations.NetworkConfiguration{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
