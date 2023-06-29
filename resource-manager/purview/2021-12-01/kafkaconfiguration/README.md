
## `github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/kafkaconfiguration` Documentation

The `kafkaconfiguration` SDK allows for interaction with the Azure Resource Manager Service `purview` (API Version `2021-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/kafkaconfiguration"
```


### Client Initialization

```go
client := kafkaconfiguration.NewKafkaConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `KafkaConfigurationClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := kafkaconfiguration.NewKafkaConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "kafkaConfigurationValue")

payload := kafkaconfiguration.KafkaConfiguration{
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


### Example Usage: `KafkaConfigurationClient.Delete`

```go
ctx := context.TODO()
id := kafkaconfiguration.NewKafkaConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "kafkaConfigurationValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KafkaConfigurationClient.Get`

```go
ctx := context.TODO()
id := kafkaconfiguration.NewKafkaConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "kafkaConfigurationValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KafkaConfigurationClient.ListByAccount`

```go
ctx := context.TODO()
id := kafkaconfiguration.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

// alternatively `client.ListByAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
