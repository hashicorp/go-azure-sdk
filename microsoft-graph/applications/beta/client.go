package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationappmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationconnectorgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationcreatedonbehalfof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationextensionproperty"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationfederatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationhomerealmdiscoverypolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationlogo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationowner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationsynchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationsynchronizationjob"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationsynchronizationjobbulkupload"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationsynchronizationjobschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationsynchronizationjobschemadirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationsynchronizationsecret"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationsynchronizationtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationsynchronizationtemplateschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationsynchronizationtemplateschemadirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationtokenissuancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/applicationtokenlifetimepolicy"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	Application                                       *application.ApplicationClient
	ApplicationAppManagementPolicy                    *applicationappmanagementpolicy.ApplicationAppManagementPolicyClient
	ApplicationConnectorGroup                         *applicationconnectorgroup.ApplicationConnectorGroupClient
	ApplicationCreatedOnBehalfOf                      *applicationcreatedonbehalfof.ApplicationCreatedOnBehalfOfClient
	ApplicationExtensionProperty                      *applicationextensionproperty.ApplicationExtensionPropertyClient
	ApplicationFederatedIdentityCredential            *applicationfederatedidentitycredential.ApplicationFederatedIdentityCredentialClient
	ApplicationHomeRealmDiscoveryPolicy               *applicationhomerealmdiscoverypolicy.ApplicationHomeRealmDiscoveryPolicyClient
	ApplicationLogo                                   *applicationlogo.ApplicationLogoClient
	ApplicationOwner                                  *applicationowner.ApplicationOwnerClient
	ApplicationSynchronization                        *applicationsynchronization.ApplicationSynchronizationClient
	ApplicationSynchronizationJob                     *applicationsynchronizationjob.ApplicationSynchronizationJobClient
	ApplicationSynchronizationJobBulkUpload           *applicationsynchronizationjobbulkupload.ApplicationSynchronizationJobBulkUploadClient
	ApplicationSynchronizationJobSchema               *applicationsynchronizationjobschema.ApplicationSynchronizationJobSchemaClient
	ApplicationSynchronizationJobSchemaDirectory      *applicationsynchronizationjobschemadirectory.ApplicationSynchronizationJobSchemaDirectoryClient
	ApplicationSynchronizationSecret                  *applicationsynchronizationsecret.ApplicationSynchronizationSecretClient
	ApplicationSynchronizationTemplate                *applicationsynchronizationtemplate.ApplicationSynchronizationTemplateClient
	ApplicationSynchronizationTemplateSchema          *applicationsynchronizationtemplateschema.ApplicationSynchronizationTemplateSchemaClient
	ApplicationSynchronizationTemplateSchemaDirectory *applicationsynchronizationtemplateschemadirectory.ApplicationSynchronizationTemplateSchemaDirectoryClient
	ApplicationTokenIssuancePolicy                    *applicationtokenissuancepolicy.ApplicationTokenIssuancePolicyClient
	ApplicationTokenLifetimePolicy                    *applicationtokenlifetimepolicy.ApplicationTokenLifetimePolicyClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	applicationAppManagementPolicyClient, err := applicationappmanagementpolicy.NewApplicationAppManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationAppManagementPolicy client: %+v", err)
	}
	configureFunc(applicationAppManagementPolicyClient.Client)

	applicationClient, err := application.NewApplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Application client: %+v", err)
	}
	configureFunc(applicationClient.Client)

	applicationConnectorGroupClient, err := applicationconnectorgroup.NewApplicationConnectorGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationConnectorGroup client: %+v", err)
	}
	configureFunc(applicationConnectorGroupClient.Client)

	applicationCreatedOnBehalfOfClient, err := applicationcreatedonbehalfof.NewApplicationCreatedOnBehalfOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationCreatedOnBehalfOf client: %+v", err)
	}
	configureFunc(applicationCreatedOnBehalfOfClient.Client)

	applicationExtensionPropertyClient, err := applicationextensionproperty.NewApplicationExtensionPropertyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationExtensionProperty client: %+v", err)
	}
	configureFunc(applicationExtensionPropertyClient.Client)

	applicationFederatedIdentityCredentialClient, err := applicationfederatedidentitycredential.NewApplicationFederatedIdentityCredentialClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationFederatedIdentityCredential client: %+v", err)
	}
	configureFunc(applicationFederatedIdentityCredentialClient.Client)

	applicationHomeRealmDiscoveryPolicyClient, err := applicationhomerealmdiscoverypolicy.NewApplicationHomeRealmDiscoveryPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationHomeRealmDiscoveryPolicy client: %+v", err)
	}
	configureFunc(applicationHomeRealmDiscoveryPolicyClient.Client)

	applicationLogoClient, err := applicationlogo.NewApplicationLogoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationLogo client: %+v", err)
	}
	configureFunc(applicationLogoClient.Client)

	applicationOwnerClient, err := applicationowner.NewApplicationOwnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationOwner client: %+v", err)
	}
	configureFunc(applicationOwnerClient.Client)

	applicationSynchronizationClient, err := applicationsynchronization.NewApplicationSynchronizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationSynchronization client: %+v", err)
	}
	configureFunc(applicationSynchronizationClient.Client)

	applicationSynchronizationJobBulkUploadClient, err := applicationsynchronizationjobbulkupload.NewApplicationSynchronizationJobBulkUploadClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationSynchronizationJobBulkUpload client: %+v", err)
	}
	configureFunc(applicationSynchronizationJobBulkUploadClient.Client)

	applicationSynchronizationJobClient, err := applicationsynchronizationjob.NewApplicationSynchronizationJobClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationSynchronizationJob client: %+v", err)
	}
	configureFunc(applicationSynchronizationJobClient.Client)

	applicationSynchronizationJobSchemaClient, err := applicationsynchronizationjobschema.NewApplicationSynchronizationJobSchemaClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationSynchronizationJobSchema client: %+v", err)
	}
	configureFunc(applicationSynchronizationJobSchemaClient.Client)

	applicationSynchronizationJobSchemaDirectoryClient, err := applicationsynchronizationjobschemadirectory.NewApplicationSynchronizationJobSchemaDirectoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationSynchronizationJobSchemaDirectory client: %+v", err)
	}
	configureFunc(applicationSynchronizationJobSchemaDirectoryClient.Client)

	applicationSynchronizationSecretClient, err := applicationsynchronizationsecret.NewApplicationSynchronizationSecretClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationSynchronizationSecret client: %+v", err)
	}
	configureFunc(applicationSynchronizationSecretClient.Client)

	applicationSynchronizationTemplateClient, err := applicationsynchronizationtemplate.NewApplicationSynchronizationTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationSynchronizationTemplate client: %+v", err)
	}
	configureFunc(applicationSynchronizationTemplateClient.Client)

	applicationSynchronizationTemplateSchemaClient, err := applicationsynchronizationtemplateschema.NewApplicationSynchronizationTemplateSchemaClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationSynchronizationTemplateSchema client: %+v", err)
	}
	configureFunc(applicationSynchronizationTemplateSchemaClient.Client)

	applicationSynchronizationTemplateSchemaDirectoryClient, err := applicationsynchronizationtemplateschemadirectory.NewApplicationSynchronizationTemplateSchemaDirectoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationSynchronizationTemplateSchemaDirectory client: %+v", err)
	}
	configureFunc(applicationSynchronizationTemplateSchemaDirectoryClient.Client)

	applicationTokenIssuancePolicyClient, err := applicationtokenissuancepolicy.NewApplicationTokenIssuancePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationTokenIssuancePolicy client: %+v", err)
	}
	configureFunc(applicationTokenIssuancePolicyClient.Client)

	applicationTokenLifetimePolicyClient, err := applicationtokenlifetimepolicy.NewApplicationTokenLifetimePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationTokenLifetimePolicy client: %+v", err)
	}
	configureFunc(applicationTokenLifetimePolicyClient.Client)

	return &Client{
		Application:                                       applicationClient,
		ApplicationAppManagementPolicy:                    applicationAppManagementPolicyClient,
		ApplicationConnectorGroup:                         applicationConnectorGroupClient,
		ApplicationCreatedOnBehalfOf:                      applicationCreatedOnBehalfOfClient,
		ApplicationExtensionProperty:                      applicationExtensionPropertyClient,
		ApplicationFederatedIdentityCredential:            applicationFederatedIdentityCredentialClient,
		ApplicationHomeRealmDiscoveryPolicy:               applicationHomeRealmDiscoveryPolicyClient,
		ApplicationLogo:                                   applicationLogoClient,
		ApplicationOwner:                                  applicationOwnerClient,
		ApplicationSynchronization:                        applicationSynchronizationClient,
		ApplicationSynchronizationJob:                     applicationSynchronizationJobClient,
		ApplicationSynchronizationJobBulkUpload:           applicationSynchronizationJobBulkUploadClient,
		ApplicationSynchronizationJobSchema:               applicationSynchronizationJobSchemaClient,
		ApplicationSynchronizationJobSchemaDirectory:      applicationSynchronizationJobSchemaDirectoryClient,
		ApplicationSynchronizationSecret:                  applicationSynchronizationSecretClient,
		ApplicationSynchronizationTemplate:                applicationSynchronizationTemplateClient,
		ApplicationSynchronizationTemplateSchema:          applicationSynchronizationTemplateSchemaClient,
		ApplicationSynchronizationTemplateSchemaDirectory: applicationSynchronizationTemplateSchemaDirectoryClient,
		ApplicationTokenIssuancePolicy:                    applicationTokenIssuancePolicyClient,
		ApplicationTokenLifetimePolicy:                    applicationTokenLifetimePolicyClient,
	}, nil
}
