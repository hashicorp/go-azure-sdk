package v2020_05_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-05-01/managementlocks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-05-01/privatelinkassociation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-05-01/resourcemanagementprivatelink"
)

type Client struct {
	ManagementLocks               *managementlocks.ManagementLocksClient
	PrivateLinkAssociation        *privatelinkassociation.PrivateLinkAssociationClient
	ResourceManagementPrivateLink *resourcemanagementprivatelink.ResourceManagementPrivateLinkClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	managementLocksClient := managementlocks.NewManagementLocksClientWithBaseURI(endpoint)
	configureAuthFunc(&managementLocksClient.Client)

	privateLinkAssociationClient := privatelinkassociation.NewPrivateLinkAssociationClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkAssociationClient.Client)

	resourceManagementPrivateLinkClient := resourcemanagementprivatelink.NewResourceManagementPrivateLinkClientWithBaseURI(endpoint)
	configureAuthFunc(&resourceManagementPrivateLinkClient.Client)

	return Client{
		ManagementLocks:               &managementLocksClient,
		PrivateLinkAssociation:        &privateLinkAssociationClient,
		ResourceManagementPrivateLink: &resourceManagementPrivateLinkClient,
	}
}
