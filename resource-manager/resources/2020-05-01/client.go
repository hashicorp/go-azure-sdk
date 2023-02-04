package v2020_05_01

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-05-01/managementlocks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-05-01/privatelinkassociation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-05-01/resourcemanagementprivatelink"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ManagementLocks               *managementlocks.ManagementLocksClient
	PrivateLinkAssociation        *privatelinkassociation.PrivateLinkAssociationClient
	ResourceManagementPrivateLink *resourcemanagementprivatelink.ResourceManagementPrivateLinkClient
}

func NewClientWithBaseURI(api environments.Api, configureAuthFunc func(c *resourcemanager.Client)) (*Client, error) {
	managementLocksClient, err := managementlocks.NewManagementLocksClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ManagementLocks: %+v", err)
	}
	configureAuthFunc(managementLocksClient.Client)

	privateLinkAssociationClient, err := privatelinkassociation.NewPrivateLinkAssociationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for PrivateLinkAssociation: %+v", err)
	}
	configureAuthFunc(privateLinkAssociationClient.Client)

	resourceManagementPrivateLinkClient, err := resourcemanagementprivatelink.NewResourceManagementPrivateLinkClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ResourceManagementPrivateLink: %+v", err)
	}
	configureAuthFunc(resourceManagementPrivateLinkClient.Client)

	return &Client{
		ManagementLocks:               managementLocksClient,
		PrivateLinkAssociation:        privateLinkAssociationClient,
		ResourceManagementPrivateLink: resourceManagementPrivateLinkClient,
	}, nil
}
