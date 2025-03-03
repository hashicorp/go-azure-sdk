
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/policytokens` Documentation

The `policytokens` SDK allows for interaction with Azure Resource Manager `resources` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/policytokens"
```


### Client Initialization

```go
client := policytokens.NewPolicyTokensClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyTokensClient.Acquire`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := policytokens.PolicyTokenRequest{
	// ...
}


if err := client.AcquireThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
