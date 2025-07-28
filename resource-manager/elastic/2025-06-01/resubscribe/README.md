
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/resubscribe` Documentation

The `resubscribe` SDK allows for interaction with Azure Resource Manager `elastic` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/resubscribe"
```


### Client Initialization

```go
client := resubscribe.NewResubscribeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ResubscribeClient.OrganizationsResubscribe`

```go
ctx := context.TODO()
id := resubscribe.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := resubscribe.ResubscribeProperties{
	// ...
}


if err := client.OrganizationsResubscribeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
