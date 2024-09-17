
## `github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/logprofilesapis` Documentation

The `logprofilesapis` SDK allows for interaction with Azure Resource Manager `insights` (API Version `2016-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2016-03-01/logprofilesapis"
```


### Client Initialization

```go
client := logprofilesapis.NewLogProfilesAPIsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LogProfilesAPIsClient.LogProfilesUpdate`

```go
ctx := context.TODO()
id := logprofilesapis.NewLogProfileID("12345678-1234-9876-4563-123456789012", "logProfileValue")

payload := logprofilesapis.LogProfileResourcePatch{
	// ...
}


read, err := client.LogProfilesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
