
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/customauthenticationextension` Documentation

The `customauthenticationextension` SDK allows for interaction with Microsoft Graph `identity` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/customauthenticationextension"
```


### Client Initialization

```go
client := customauthenticationextension.NewCustomAuthenticationExtensionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CustomAuthenticationExtensionClient.CreateCustomAuthenticationExtension`

```go
ctx := context.TODO()

payload := customauthenticationextension.CustomAuthenticationExtension{
	// ...
}


read, err := client.CreateCustomAuthenticationExtension(ctx, payload, customauthenticationextension.DefaultCreateCustomAuthenticationExtensionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAuthenticationExtensionClient.DeleteCustomAuthenticationExtension`

```go
ctx := context.TODO()
id := customauthenticationextension.NewIdentityCustomAuthenticationExtensionID("customAuthenticationExtensionId")

read, err := client.DeleteCustomAuthenticationExtension(ctx, id, customauthenticationextension.DefaultDeleteCustomAuthenticationExtensionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAuthenticationExtensionClient.GetCustomAuthenticationExtension`

```go
ctx := context.TODO()
id := customauthenticationextension.NewIdentityCustomAuthenticationExtensionID("customAuthenticationExtensionId")

read, err := client.GetCustomAuthenticationExtension(ctx, id, customauthenticationextension.DefaultGetCustomAuthenticationExtensionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAuthenticationExtensionClient.GetCustomAuthenticationExtensionsCount`

```go
ctx := context.TODO()


read, err := client.GetCustomAuthenticationExtensionsCount(ctx, customauthenticationextension.DefaultGetCustomAuthenticationExtensionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAuthenticationExtensionClient.ListCustomAuthenticationExtensions`

```go
ctx := context.TODO()


// alternatively `client.ListCustomAuthenticationExtensions(ctx, customauthenticationextension.DefaultListCustomAuthenticationExtensionsOperationOptions())` can be used to do batched pagination
items, err := client.ListCustomAuthenticationExtensionsComplete(ctx, customauthenticationextension.DefaultListCustomAuthenticationExtensionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CustomAuthenticationExtensionClient.UpdateCustomAuthenticationExtension`

```go
ctx := context.TODO()
id := customauthenticationextension.NewIdentityCustomAuthenticationExtensionID("customAuthenticationExtensionId")

payload := customauthenticationextension.CustomAuthenticationExtension{
	// ...
}


read, err := client.UpdateCustomAuthenticationExtension(ctx, id, payload, customauthenticationextension.DefaultUpdateCustomAuthenticationExtensionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAuthenticationExtensionClient.ValidateCustomAuthenticationExtensionAuthenticationConfiguration`

```go
ctx := context.TODO()
id := customauthenticationextension.NewIdentityCustomAuthenticationExtensionID("customAuthenticationExtensionId")

read, err := client.ValidateCustomAuthenticationExtensionAuthenticationConfiguration(ctx, id, customauthenticationextension.DefaultValidateCustomAuthenticationExtensionAuthenticationConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomAuthenticationExtensionClient.ValidateCustomAuthenticationExtensionsAuthenticationConfiguration`

```go
ctx := context.TODO()

payload := customauthenticationextension.ValidateCustomAuthenticationExtensionsAuthenticationConfigurationRequest{
	// ...
}


read, err := client.ValidateCustomAuthenticationExtensionsAuthenticationConfiguration(ctx, payload, customauthenticationextension.DefaultValidateCustomAuthenticationExtensionsAuthenticationConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
