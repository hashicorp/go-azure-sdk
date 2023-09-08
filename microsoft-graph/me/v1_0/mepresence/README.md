
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mepresence` Documentation

The `mepresence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mepresence"
```


### Client Initialization

```go
client := mepresence.NewMePresenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MePresenceClient.CreateMePresenceClearPresence`

```go
ctx := context.TODO()

payload := mepresence.CreateMePresenceClearPresenceRequest{
	// ...
}


read, err := client.CreateMePresenceClearPresence(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePresenceClient.CreateMePresenceClearUserPreferredPresence`

```go
ctx := context.TODO()


read, err := client.CreateMePresenceClearUserPreferredPresence(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePresenceClient.DeleteMePresence`

```go
ctx := context.TODO()


read, err := client.DeleteMePresence(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePresenceClient.GetMePresence`

```go
ctx := context.TODO()


read, err := client.GetMePresence(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePresenceClient.SetMePresencePresence`

```go
ctx := context.TODO()

payload := mepresence.SetMePresencePresenceRequest{
	// ...
}


read, err := client.SetMePresencePresence(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePresenceClient.SetMePresenceUserPreferredPresence`

```go
ctx := context.TODO()

payload := mepresence.SetMePresenceUserPreferredPresenceRequest{
	// ...
}


read, err := client.SetMePresenceUserPreferredPresence(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePresenceClient.UpdateMePresence`

```go
ctx := context.TODO()

payload := mepresence.Presence{
	// ...
}


read, err := client.UpdateMePresence(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
