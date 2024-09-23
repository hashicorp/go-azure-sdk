package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/appmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/connectorgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/createdonbehalfof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/extensionproperty"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/federatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/homerealmdiscoverypolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/logo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/owner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/synchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/synchronizationjob"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/synchronizationjobbulkupload"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/synchronizationjobschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/synchronizationjobschemadirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/synchronizationsecret"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/synchronizationtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/synchronizationtemplateschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/synchronizationtemplateschemadirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/tokenissuancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/tokenlifetimepolicy"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AppManagementPolicy                    *appmanagementpolicy.AppManagementPolicyClient
	Application                            *application.ApplicationClient
	ConnectorGroup                         *connectorgroup.ConnectorGroupClient
	CreatedOnBehalfOf                      *createdonbehalfof.CreatedOnBehalfOfClient
	ExtensionProperty                      *extensionproperty.ExtensionPropertyClient
	FederatedIdentityCredential            *federatedidentitycredential.FederatedIdentityCredentialClient
	HomeRealmDiscoveryPolicy               *homerealmdiscoverypolicy.HomeRealmDiscoveryPolicyClient
	Logo                                   *logo.LogoClient
	Owner                                  *owner.OwnerClient
	Synchronization                        *synchronization.SynchronizationClient
	SynchronizationJob                     *synchronizationjob.SynchronizationJobClient
	SynchronizationJobBulkUpload           *synchronizationjobbulkupload.SynchronizationJobBulkUploadClient
	SynchronizationJobSchema               *synchronizationjobschema.SynchronizationJobSchemaClient
	SynchronizationJobSchemaDirectory      *synchronizationjobschemadirectory.SynchronizationJobSchemaDirectoryClient
	SynchronizationSecret                  *synchronizationsecret.SynchronizationSecretClient
	SynchronizationTemplate                *synchronizationtemplate.SynchronizationTemplateClient
	SynchronizationTemplateSchema          *synchronizationtemplateschema.SynchronizationTemplateSchemaClient
	SynchronizationTemplateSchemaDirectory *synchronizationtemplateschemadirectory.SynchronizationTemplateSchemaDirectoryClient
	TokenIssuancePolicy                    *tokenissuancepolicy.TokenIssuancePolicyClient
	TokenLifetimePolicy                    *tokenlifetimepolicy.TokenLifetimePolicyClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	appManagementPolicyClient, err := appmanagementpolicy.NewAppManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppManagementPolicy client: %+v", err)
	}
	configureFunc(appManagementPolicyClient.Client)

	applicationClient, err := application.NewApplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Application client: %+v", err)
	}
	configureFunc(applicationClient.Client)

	connectorGroupClient, err := connectorgroup.NewConnectorGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectorGroup client: %+v", err)
	}
	configureFunc(connectorGroupClient.Client)

	createdOnBehalfOfClient, err := createdonbehalfof.NewCreatedOnBehalfOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CreatedOnBehalfOf client: %+v", err)
	}
	configureFunc(createdOnBehalfOfClient.Client)

	extensionPropertyClient, err := extensionproperty.NewExtensionPropertyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExtensionProperty client: %+v", err)
	}
	configureFunc(extensionPropertyClient.Client)

	federatedIdentityCredentialClient, err := federatedidentitycredential.NewFederatedIdentityCredentialClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FederatedIdentityCredential client: %+v", err)
	}
	configureFunc(federatedIdentityCredentialClient.Client)

	homeRealmDiscoveryPolicyClient, err := homerealmdiscoverypolicy.NewHomeRealmDiscoveryPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HomeRealmDiscoveryPolicy client: %+v", err)
	}
	configureFunc(homeRealmDiscoveryPolicyClient.Client)

	logoClient, err := logo.NewLogoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Logo client: %+v", err)
	}
	configureFunc(logoClient.Client)

	ownerClient, err := owner.NewOwnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Owner client: %+v", err)
	}
	configureFunc(ownerClient.Client)

	synchronizationClient, err := synchronization.NewSynchronizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Synchronization client: %+v", err)
	}
	configureFunc(synchronizationClient.Client)

	synchronizationJobBulkUploadClient, err := synchronizationjobbulkupload.NewSynchronizationJobBulkUploadClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SynchronizationJobBulkUpload client: %+v", err)
	}
	configureFunc(synchronizationJobBulkUploadClient.Client)

	synchronizationJobClient, err := synchronizationjob.NewSynchronizationJobClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SynchronizationJob client: %+v", err)
	}
	configureFunc(synchronizationJobClient.Client)

	synchronizationJobSchemaClient, err := synchronizationjobschema.NewSynchronizationJobSchemaClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SynchronizationJobSchema client: %+v", err)
	}
	configureFunc(synchronizationJobSchemaClient.Client)

	synchronizationJobSchemaDirectoryClient, err := synchronizationjobschemadirectory.NewSynchronizationJobSchemaDirectoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SynchronizationJobSchemaDirectory client: %+v", err)
	}
	configureFunc(synchronizationJobSchemaDirectoryClient.Client)

	synchronizationSecretClient, err := synchronizationsecret.NewSynchronizationSecretClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SynchronizationSecret client: %+v", err)
	}
	configureFunc(synchronizationSecretClient.Client)

	synchronizationTemplateClient, err := synchronizationtemplate.NewSynchronizationTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SynchronizationTemplate client: %+v", err)
	}
	configureFunc(synchronizationTemplateClient.Client)

	synchronizationTemplateSchemaClient, err := synchronizationtemplateschema.NewSynchronizationTemplateSchemaClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SynchronizationTemplateSchema client: %+v", err)
	}
	configureFunc(synchronizationTemplateSchemaClient.Client)

	synchronizationTemplateSchemaDirectoryClient, err := synchronizationtemplateschemadirectory.NewSynchronizationTemplateSchemaDirectoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SynchronizationTemplateSchemaDirectory client: %+v", err)
	}
	configureFunc(synchronizationTemplateSchemaDirectoryClient.Client)

	tokenIssuancePolicyClient, err := tokenissuancepolicy.NewTokenIssuancePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TokenIssuancePolicy client: %+v", err)
	}
	configureFunc(tokenIssuancePolicyClient.Client)

	tokenLifetimePolicyClient, err := tokenlifetimepolicy.NewTokenLifetimePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TokenLifetimePolicy client: %+v", err)
	}
	configureFunc(tokenLifetimePolicyClient.Client)

	return &Client{
		AppManagementPolicy:                    appManagementPolicyClient,
		Application:                            applicationClient,
		ConnectorGroup:                         connectorGroupClient,
		CreatedOnBehalfOf:                      createdOnBehalfOfClient,
		ExtensionProperty:                      extensionPropertyClient,
		FederatedIdentityCredential:            federatedIdentityCredentialClient,
		HomeRealmDiscoveryPolicy:               homeRealmDiscoveryPolicyClient,
		Logo:                                   logoClient,
		Owner:                                  ownerClient,
		Synchronization:                        synchronizationClient,
		SynchronizationJob:                     synchronizationJobClient,
		SynchronizationJobBulkUpload:           synchronizationJobBulkUploadClient,
		SynchronizationJobSchema:               synchronizationJobSchemaClient,
		SynchronizationJobSchemaDirectory:      synchronizationJobSchemaDirectoryClient,
		SynchronizationSecret:                  synchronizationSecretClient,
		SynchronizationTemplate:                synchronizationTemplateClient,
		SynchronizationTemplateSchema:          synchronizationTemplateSchemaClient,
		SynchronizationTemplateSchemaDirectory: synchronizationTemplateSchemaDirectoryClient,
		TokenIssuancePolicy:                    tokenIssuancePolicyClient,
		TokenLifetimePolicy:                    tokenLifetimePolicyClient,
	}, nil
}
