
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/apikey` Documentation

The `apikey` SDK allows for interaction with Azure Resource Manager `elastic` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/apikey"
```


### Client Initialization

```go
client := apikey.NewApiKeyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApiKeyClient.OrganizationsGetApiKey`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := apikey.UserEmailId{
	// ...
}


read, err := client.OrganizationsGetApiKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
