
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/afdprofiles` Documentation

The `afdprofiles` SDK allows for interaction with Azure Resource Manager `cdn` (API Version `2025-04-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-04-15/afdprofiles"
```


### Client Initialization

```go
client := afdprofiles.NewAFDProfilesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AFDProfilesClient.CheckEndpointNameAvailability`

```go
ctx := context.TODO()
id := afdprofiles.NewProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName")

payload := afdprofiles.CheckEndpointNameAvailabilityInput{
	// ...
}


read, err := client.CheckEndpointNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AFDProfilesClient.CheckHostNameAvailability`

```go
ctx := context.TODO()
id := afdprofiles.NewProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName")

payload := afdprofiles.CheckHostNameAvailabilityInput{
	// ...
}


read, err := client.CheckHostNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AFDProfilesClient.ListResourceUsage`

```go
ctx := context.TODO()
id := afdprofiles.NewProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName")

// alternatively `client.ListResourceUsage(ctx, id)` can be used to do batched pagination
items, err := client.ListResourceUsageComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AFDProfilesClient.Upgrade`

```go
ctx := context.TODO()
id := afdprofiles.NewProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName")

payload := afdprofiles.ProfileUpgradeParameters{
	// ...
}


if err := client.UpgradeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AFDProfilesClient.ValidateSecret`

```go
ctx := context.TODO()
id := afdprofiles.NewProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName")

payload := afdprofiles.ValidateSecretInput{
	// ...
}


read, err := client.ValidateSecret(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
