
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/csmpublishingcredentialspoliciesentities` Documentation

The `csmpublishingcredentialspoliciesentities` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/csmpublishingcredentialspoliciesentities"
```


### Client Initialization

```go
client := csmpublishingcredentialspoliciesentities.NewCsmPublishingCredentialsPoliciesEntitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CsmPublishingCredentialsPoliciesEntitiesClient.WebAppsGetFtpAllowed`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetFtpAllowed(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CsmPublishingCredentialsPoliciesEntitiesClient.WebAppsListBasicPublishingCredentialsPolicies`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListBasicPublishingCredentialsPolicies(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListBasicPublishingCredentialsPoliciesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CsmPublishingCredentialsPoliciesEntitiesClient.WebAppsUpdateFtpAllowed`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := csmpublishingcredentialspoliciesentities.CsmPublishingCredentialsPoliciesEntity{
	// ...
}


read, err := client.WebAppsUpdateFtpAllowed(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
