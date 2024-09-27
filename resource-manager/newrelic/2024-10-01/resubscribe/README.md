
## `github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2024-10-01/resubscribe` Documentation

The `resubscribe` SDK allows for interaction with Azure Resource Manager `newrelic` (API Version `2024-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2024-10-01/resubscribe"
```


### Client Initialization

```go
client := resubscribe.NewResubscribeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ResubscribeClient.MonitorsResubscribe`

```go
ctx := context.TODO()
id := resubscribe.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := resubscribe.ResubscribeProperties{
	// ...
}


if err := client.MonitorsResubscribeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
