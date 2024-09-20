
## `github.com/hashicorp/go-azure-sdk/resource-manager/powerbidedicated/2021-01-01/autoscalevcores` Documentation

The `autoscalevcores` SDK allows for interaction with Azure Resource Manager `powerbidedicated` (API Version `2021-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/powerbidedicated/2021-01-01/autoscalevcores"
```


### Client Initialization

```go
client := autoscalevcores.NewAutoScaleVCoresClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AutoScaleVCoresClient.Create`

```go
ctx := context.TODO()
id := autoscalevcores.NewAutoScaleVCoreID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vcoreName")

payload := autoscalevcores.AutoScaleVCore{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AutoScaleVCoresClient.Delete`

```go
ctx := context.TODO()
id := autoscalevcores.NewAutoScaleVCoreID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vcoreName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AutoScaleVCoresClient.Get`

```go
ctx := context.TODO()
id := autoscalevcores.NewAutoScaleVCoreID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vcoreName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AutoScaleVCoresClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.ListByResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AutoScaleVCoresClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ListBySubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AutoScaleVCoresClient.Update`

```go
ctx := context.TODO()
id := autoscalevcores.NewAutoScaleVCoreID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vcoreName")

payload := autoscalevcores.AutoScaleVCoreUpdateParameters{
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
