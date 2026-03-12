
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/csmpublishingcredentialspoliciesentityftpallowedslot` Documentation

The `csmpublishingcredentialspoliciesentityftpallowedslot` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/csmpublishingcredentialspoliciesentityftpallowedslot"
```


### Client Initialization

```go
client := csmpublishingcredentialspoliciesentityftpallowedslot.NewCsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient.WebAppsGetFtpAllowedSlot`

```go
ctx := context.TODO()
id := csmpublishingcredentialspoliciesentityftpallowedslot.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.WebAppsGetFtpAllowedSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient.WebAppsListBasicPublishingCredentialsPoliciesSlot`

```go
ctx := context.TODO()
id := csmpublishingcredentialspoliciesentityftpallowedslot.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListBasicPublishingCredentialsPoliciesSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListBasicPublishingCredentialsPoliciesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient.WebAppsUpdateFtpAllowedSlot`

```go
ctx := context.TODO()
id := csmpublishingcredentialspoliciesentityftpallowedslot.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := csmpublishingcredentialspoliciesentityftpallowedslot.CsmPublishingCredentialsPoliciesEntity{
	// ...
}


read, err := client.WebAppsUpdateFtpAllowedSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
