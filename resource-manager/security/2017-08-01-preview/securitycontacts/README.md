
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


### Example Usage: `SecurityContactsClient.SecurityContactsCreate`

```go
ctx := context.TODO()
id := securitycontacts.NewSecurityContactID("12345678-1234-9876-4563-123456789012", "securityContactValue")

payload := securitycontacts.SecurityContact{
	// ...
}


read, err := client.SecurityContactsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityContactsClient.SecurityContactsDelete`

```go
ctx := context.TODO()
id := securitycontacts.NewSecurityContactID("12345678-1234-9876-4563-123456789012", "securityContactValue")

read, err := client.SecurityContactsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityContactsClient.SecurityContactsGet`

```go
ctx := context.TODO()
id := securitycontacts.NewSecurityContactID("12345678-1234-9876-4563-123456789012", "securityContactValue")

read, err := client.SecurityContactsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityContactsClient.SecurityContactsList`

```go
ctx := context.TODO()
id := securitycontacts.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.SecurityContactsList(ctx, id)` can be used to do batched pagination
items, err := client.SecurityContactsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SecurityContactsClient.SecurityContactsUpdate`

```go
ctx := context.TODO()
id := securitycontacts.NewSecurityContactID("12345678-1234-9876-4563-123456789012", "securityContactValue")

payload := securitycontacts.SecurityContact{
	// ...
}


read, err := client.SecurityContactsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
