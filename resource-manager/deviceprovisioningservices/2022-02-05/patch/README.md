
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/patch` Documentation

The `patch` SDK allows for interaction with the Azure Resource Manager Service `deviceprovisioningservices` (API Version `2022-02-05`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/patch"
```


### Client Initialization

```go
client := patch.NewPATCHClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PATCHClient.IotDpsResourceUpdate`

```go
ctx := context.TODO()
id := patch.NewProvisioningServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue")

payload := patch.TagsResource{
	// ...
}


if err := client.IotDpsResourceUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
