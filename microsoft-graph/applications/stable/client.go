package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/appmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/createdonbehalfof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/extensionproperty"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/federatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/homerealmdiscoverypolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/logo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/owner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationjob"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationjobbulkupload"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationjobschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationjobschemadirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationsecret"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationtemplateschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationtemplateschemadirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/tokenissuancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/tokenlifetimepolicy"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AppManagementPolicy                    *appmanagementpolicy.AppManagementPolicyClient
	Application                            *application.ApplicationClient
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
