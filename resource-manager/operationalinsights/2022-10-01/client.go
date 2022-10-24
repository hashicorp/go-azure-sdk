package v2022_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2022-10-01/deletedworkspaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2022-10-01/tables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2022-10-01/workspaces"
)

type Client struct {
	DeletedWorkspaces *deletedworkspaces.DeletedWorkspacesClient
	Tables            *tables.TablesClient
	Workspaces        *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	deletedWorkspacesClient := deletedworkspaces.NewDeletedWorkspacesClientWithBaseURI(endpoint)
	configureAuthFunc(&deletedWorkspacesClient.Client)

	tablesClient := tables.NewTablesClientWithBaseURI(endpoint)
	configureAuthFunc(&tablesClient.Client)

	workspacesClient := workspaces.NewWorkspacesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacesClient.Client)

	return Client{
		DeletedWorkspaces: &deletedWorkspacesClient,
		Tables:            &tablesClient,
		Workspaces:        &workspacesClient,
	}
}
