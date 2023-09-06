
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/securitycontacts` Documentation

The `securitycontacts` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2017-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/securitycontacts"
```


### Client Initialization

```go
client := securitycontacts.NewSecurityContactsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SecurityContactsClient.Create`

```go
ctx := context.TODO()
id := securitycontacts.NewSecurityContactID("12345678-1234-9876-4563-123456789012", "securityContactValue")

payload := securitycontacts.SecurityContact{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityContactsClient.Delete`

```go
ctx := context.TODO()
id := securitycontacts.NewSecurityContactID("12345678-1234-9876-4563-123456789012", "securityContactValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityContactsClient.Get`

```go
ctx := context.TODO()
id := securitycontacts.NewSecurityContactID("12345678-1234-9876-4563-123456789012", "securityContactValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityContactsClient.List`

```go
ctx := context.TODO()
id := securitycontacts.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SecurityContactsClient.Update`

```go
ctx := context.TODO()
id := securitycontacts.NewSecurityContactID("12345678-1234-9876-4563-123456789012", "securityContactValue")

payload := securitycontacts.SecurityContact{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
