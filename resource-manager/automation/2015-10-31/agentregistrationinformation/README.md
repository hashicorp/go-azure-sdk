
## `github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/agentregistrationinformation` Documentation

The `agentregistrationinformation` SDK allows for interaction with Azure Resource Manager `automation` (API Version `2015-10-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/agentregistrationinformation"
```


### Client Initialization

```go
client := agentregistrationinformation.NewAgentRegistrationInformationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AgentRegistrationInformationClient.Get`

```go
ctx := context.TODO()
id := agentregistrationinformation.NewAutomationAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentRegistrationInformationClient.RegenerateKey`

```go
ctx := context.TODO()
id := agentregistrationinformation.NewAutomationAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue")

payload := agentregistrationinformation.AgentRegistrationRegenerateKeyParameter{
	// ...
}


read, err := client.RegenerateKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
