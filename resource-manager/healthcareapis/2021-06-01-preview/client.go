package v2021_06_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/collection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/dicomservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/fhirservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/iotconnectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/proxy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/resource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/workspaces"
)

type Client struct {
	Collection                 *collection.CollectionClient
	DicomServices              *dicomservices.DicomServicesClient
	FhirServices               *fhirservices.FhirServicesClient
	IotConnectors              *iotconnectors.IotConnectorsClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	Proxy                      *proxy.ProxyClient
	Resource                   *resource.ResourceClient
	Workspaces                 *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	collectionClient := collection.NewCollectionClientWithBaseURI(endpoint)
	configureAuthFunc(&collectionClient.Client)

	dicomServicesClient := dicomservices.NewDicomServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&dicomServicesClient.Client)

	fhirServicesClient := fhirservices.NewFhirServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&fhirServicesClient.Client)

	iotConnectorsClient := iotconnectors.NewIotConnectorsClientWithBaseURI(endpoint)
	configureAuthFunc(&iotConnectorsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	proxyClient := proxy.NewProxyClientWithBaseURI(endpoint)
	configureAuthFunc(&proxyClient.Client)

	resourceClient := resource.NewResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&resourceClient.Client)

	workspacesClient := workspaces.NewWorkspacesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacesClient.Client)

	return Client{
		Collection:                 &collectionClient,
		DicomServices:              &dicomServicesClient,
		FhirServices:               &fhirServicesClient,
		IotConnectors:              &iotConnectorsClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
		Proxy:                      &proxyClient,
		Resource:                   &resourceClient,
		Workspaces:                 &workspacesClient,
	}
}
