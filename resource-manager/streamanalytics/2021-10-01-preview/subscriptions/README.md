
## `github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/subscriptions` Documentation

The `subscriptions` SDK allows for interaction with the Azure Resource Manager Service `streamanalytics` (API Version `2021-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/streamanalytics/2021-10-01-preview/subscriptions"
```


### Client Initialization

```go
client := subscriptions.NewSubscriptionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SubscriptionsClient.CompileQuery`

```go
ctx := context.TODO()
id := subscriptions.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := subscriptions.CompileQuery{
	// ...
}


read, err := client.CompileQuery(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionsClient.ListQuotas`

```go
ctx := context.TODO()
id := subscriptions.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.ListQuotas(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionsClient.SampleInput`

```go
ctx := context.TODO()
id := subscriptions.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := subscriptions.SampleInput{
	// ...
}


if err := client.SampleInputThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SubscriptionsClient.TestInput`

```go
ctx := context.TODO()
id := subscriptions.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := subscriptions.TestInput{
	// ...
}


if err := client.TestInputThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SubscriptionsClient.TestOutput`

```go
ctx := context.TODO()
id := subscriptions.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := subscriptions.TestOutput{
	// ...
}


if err := client.TestOutputThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SubscriptionsClient.TestQuery`

```go
ctx := context.TODO()
id := subscriptions.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := subscriptions.TestQuery{
	// ...
}


if err := client.TestQueryThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
