package v2021_05_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2021-05-01/application"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2021-05-01/applicationtype"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2021-05-01/applicationtypeversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2021-05-01/managedcluster"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2021-05-01/managedclusterversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2021-05-01/nodetype"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2021-05-01/service"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2021-05-01/services"
)

type Client struct {
	Application            *application.ApplicationClient
	ApplicationType        *applicationtype.ApplicationTypeClient
	ApplicationTypeVersion *applicationtypeversion.ApplicationTypeVersionClient
	ManagedCluster         *managedcluster.ManagedClusterClient
	ManagedClusterVersion  *managedclusterversion.ManagedClusterVersionClient
	NodeType               *nodetype.NodeTypeClient
	Service                *service.ServiceClient
	Services               *services.ServicesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	applicationClient := application.NewApplicationClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationClient.Client)

	applicationTypeClient := applicationtype.NewApplicationTypeClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationTypeClient.Client)

	applicationTypeVersionClient := applicationtypeversion.NewApplicationTypeVersionClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationTypeVersionClient.Client)

	managedClusterClient := managedcluster.NewManagedClusterClientWithBaseURI(endpoint)
	configureAuthFunc(&managedClusterClient.Client)

	managedClusterVersionClient := managedclusterversion.NewManagedClusterVersionClientWithBaseURI(endpoint)
	configureAuthFunc(&managedClusterVersionClient.Client)

	nodeTypeClient := nodetype.NewNodeTypeClientWithBaseURI(endpoint)
	configureAuthFunc(&nodeTypeClient.Client)

	serviceClient := service.NewServiceClientWithBaseURI(endpoint)
	configureAuthFunc(&serviceClient.Client)

	servicesClient := services.NewServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&servicesClient.Client)

	return Client{
		Application:            &applicationClient,
		ApplicationType:        &applicationTypeClient,
		ApplicationTypeVersion: &applicationTypeVersionClient,
		ManagedCluster:         &managedClusterClient,
		ManagedClusterVersion:  &managedClusterVersionClient,
		NodeType:               &nodeTypeClient,
		Service:                &serviceClient,
		Services:               &servicesClient,
	}
}
