
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/taskscommons` Documentation

The `taskscommons` SDK allows for interaction with the Azure Resource Manager Service `datamigration` (API Version `2018-04-19`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/taskscommons"
```


### Client Initialization

```go
client := taskscommons.NewTasksCommonsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```

