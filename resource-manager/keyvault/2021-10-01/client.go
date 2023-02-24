package v2021_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2021-10-01/keys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2021-10-01/managedhsms"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2021-10-01/mhsmlistprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2021-10-01/mhsmprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2021-10-01/mhsmprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2021-10-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2021-10-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2021-10-01/secrets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2021-10-01/vaults"
)

type Client struct {
	Keys                               *keys.KeysClient
	MHSMListPrivateEndpointConnections *mhsmlistprivateendpointconnections.MHSMListPrivateEndpointConnectionsClient
	MHSMPrivateEndpointConnections     *mhsmprivateendpointconnections.MHSMPrivateEndpointConnectionsClient
	MHSMPrivateLinkResources           *mhsmprivatelinkresources.MHSMPrivateLinkResourcesClient
	ManagedHsms                        *managedhsms.ManagedHsmsClient
	PrivateEndpointConnections         *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources               *privatelinkresources.PrivateLinkResourcesClient
	Secrets                            *secrets.SecretsClient
	Vaults                             *vaults.VaultsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	keysClient := keys.NewKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&keysClient.Client)

	mHSMListPrivateEndpointConnectionsClient := mhsmlistprivateendpointconnections.NewMHSMListPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&mHSMListPrivateEndpointConnectionsClient.Client)

	mHSMPrivateEndpointConnectionsClient := mhsmprivateendpointconnections.NewMHSMPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&mHSMPrivateEndpointConnectionsClient.Client)

	mHSMPrivateLinkResourcesClient := mhsmprivatelinkresources.NewMHSMPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&mHSMPrivateLinkResourcesClient.Client)

	managedHsmsClient := managedhsms.NewManagedHsmsClientWithBaseURI(endpoint)
	configureAuthFunc(&managedHsmsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	secretsClient := secrets.NewSecretsClientWithBaseURI(endpoint)
	configureAuthFunc(&secretsClient.Client)

	vaultsClient := vaults.NewVaultsClientWithBaseURI(endpoint)
	configureAuthFunc(&vaultsClient.Client)

	return Client{
		Keys:                               &keysClient,
		MHSMListPrivateEndpointConnections: &mHSMListPrivateEndpointConnectionsClient,
		MHSMPrivateEndpointConnections:     &mHSMPrivateEndpointConnectionsClient,
		MHSMPrivateLinkResources:           &mHSMPrivateLinkResourcesClient,
		ManagedHsms:                        &managedHsmsClient,
		PrivateEndpointConnections:         &privateEndpointConnectionsClient,
		PrivateLinkResources:               &privateLinkResourcesClient,
		Secrets:                            &secretsClient,
		Vaults:                             &vaultsClient,
	}
}
