package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalappmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalapproleassignedto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalapproleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalclaimsmappingpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalcreatedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipaldelegatedpermissionclassification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalendpoint"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalfederatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalhomerealmdiscoverypolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipallicensedetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalmemberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipaloauth2permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalownedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalowner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalsynchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalsynchronizationjob"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalsynchronizationjobbulkupload"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalsynchronizationjobschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalsynchronizationjobschemadirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalsynchronizationsecret"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalsynchronizationtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalsynchronizationtemplateschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipalsynchronizationtemplateschemadirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipaltokenissuancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipaltokenlifetimepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipaltransitivememberof"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	ServicePrincipal                                       *serviceprincipal.ServicePrincipalClient
	ServicePrincipalAppManagementPolicy                    *serviceprincipalappmanagementpolicy.ServicePrincipalAppManagementPolicyClient
	ServicePrincipalAppRoleAssignedTo                      *serviceprincipalapproleassignedto.ServicePrincipalAppRoleAssignedToClient
	ServicePrincipalAppRoleAssignment                      *serviceprincipalapproleassignment.ServicePrincipalAppRoleAssignmentClient
	ServicePrincipalClaimsMappingPolicy                    *serviceprincipalclaimsmappingpolicy.ServicePrincipalClaimsMappingPolicyClient
	ServicePrincipalCreatedObject                          *serviceprincipalcreatedobject.ServicePrincipalCreatedObjectClient
	ServicePrincipalDelegatedPermissionClassification      *serviceprincipaldelegatedpermissionclassification.ServicePrincipalDelegatedPermissionClassificationClient
	ServicePrincipalEndpoint                               *serviceprincipalendpoint.ServicePrincipalEndpointClient
	ServicePrincipalFederatedIdentityCredential            *serviceprincipalfederatedidentitycredential.ServicePrincipalFederatedIdentityCredentialClient
	ServicePrincipalHomeRealmDiscoveryPolicy               *serviceprincipalhomerealmdiscoverypolicy.ServicePrincipalHomeRealmDiscoveryPolicyClient
	ServicePrincipalLicenseDetail                          *serviceprincipallicensedetail.ServicePrincipalLicenseDetailClient
	ServicePrincipalMemberOf                               *serviceprincipalmemberof.ServicePrincipalMemberOfClient
	ServicePrincipalOauth2PermissionGrant                  *serviceprincipaloauth2permissiongrant.ServicePrincipalOauth2PermissionGrantClient
	ServicePrincipalOwnedObject                            *serviceprincipalownedobject.ServicePrincipalOwnedObjectClient
	ServicePrincipalOwner                                  *serviceprincipalowner.ServicePrincipalOwnerClient
	ServicePrincipalSynchronization                        *serviceprincipalsynchronization.ServicePrincipalSynchronizationClient
	ServicePrincipalSynchronizationJob                     *serviceprincipalsynchronizationjob.ServicePrincipalSynchronizationJobClient
	ServicePrincipalSynchronizationJobBulkUpload           *serviceprincipalsynchronizationjobbulkupload.ServicePrincipalSynchronizationJobBulkUploadClient
	ServicePrincipalSynchronizationJobSchema               *serviceprincipalsynchronizationjobschema.ServicePrincipalSynchronizationJobSchemaClient
	ServicePrincipalSynchronizationJobSchemaDirectory      *serviceprincipalsynchronizationjobschemadirectory.ServicePrincipalSynchronizationJobSchemaDirectoryClient
	ServicePrincipalSynchronizationSecret                  *serviceprincipalsynchronizationsecret.ServicePrincipalSynchronizationSecretClient
	ServicePrincipalSynchronizationTemplate                *serviceprincipalsynchronizationtemplate.ServicePrincipalSynchronizationTemplateClient
	ServicePrincipalSynchronizationTemplateSchema          *serviceprincipalsynchronizationtemplateschema.ServicePrincipalSynchronizationTemplateSchemaClient
	ServicePrincipalSynchronizationTemplateSchemaDirectory *serviceprincipalsynchronizationtemplateschemadirectory.ServicePrincipalSynchronizationTemplateSchemaDirectoryClient
	ServicePrincipalTokenIssuancePolicy                    *serviceprincipaltokenissuancepolicy.ServicePrincipalTokenIssuancePolicyClient
	ServicePrincipalTokenLifetimePolicy                    *serviceprincipaltokenlifetimepolicy.ServicePrincipalTokenLifetimePolicyClient
	ServicePrincipalTransitiveMemberOf                     *serviceprincipaltransitivememberof.ServicePrincipalTransitiveMemberOfClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	servicePrincipalAppManagementPolicyClient, err := serviceprincipalappmanagementpolicy.NewServicePrincipalAppManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalAppManagementPolicy client: %+v", err)
	}
	configureFunc(servicePrincipalAppManagementPolicyClient.Client)

	servicePrincipalAppRoleAssignedToClient, err := serviceprincipalapproleassignedto.NewServicePrincipalAppRoleAssignedToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalAppRoleAssignedTo client: %+v", err)
	}
	configureFunc(servicePrincipalAppRoleAssignedToClient.Client)

	servicePrincipalAppRoleAssignmentClient, err := serviceprincipalapproleassignment.NewServicePrincipalAppRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalAppRoleAssignment client: %+v", err)
	}
	configureFunc(servicePrincipalAppRoleAssignmentClient.Client)

	servicePrincipalClaimsMappingPolicyClient, err := serviceprincipalclaimsmappingpolicy.NewServicePrincipalClaimsMappingPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalClaimsMappingPolicy client: %+v", err)
	}
	configureFunc(servicePrincipalClaimsMappingPolicyClient.Client)

	servicePrincipalClient, err := serviceprincipal.NewServicePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipal client: %+v", err)
	}
	configureFunc(servicePrincipalClient.Client)

	servicePrincipalCreatedObjectClient, err := serviceprincipalcreatedobject.NewServicePrincipalCreatedObjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalCreatedObject client: %+v", err)
	}
	configureFunc(servicePrincipalCreatedObjectClient.Client)

	servicePrincipalDelegatedPermissionClassificationClient, err := serviceprincipaldelegatedpermissionclassification.NewServicePrincipalDelegatedPermissionClassificationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalDelegatedPermissionClassification client: %+v", err)
	}
	configureFunc(servicePrincipalDelegatedPermissionClassificationClient.Client)

	servicePrincipalEndpointClient, err := serviceprincipalendpoint.NewServicePrincipalEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalEndpoint client: %+v", err)
	}
	configureFunc(servicePrincipalEndpointClient.Client)

	servicePrincipalFederatedIdentityCredentialClient, err := serviceprincipalfederatedidentitycredential.NewServicePrincipalFederatedIdentityCredentialClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalFederatedIdentityCredential client: %+v", err)
	}
	configureFunc(servicePrincipalFederatedIdentityCredentialClient.Client)

	servicePrincipalHomeRealmDiscoveryPolicyClient, err := serviceprincipalhomerealmdiscoverypolicy.NewServicePrincipalHomeRealmDiscoveryPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalHomeRealmDiscoveryPolicy client: %+v", err)
	}
	configureFunc(servicePrincipalHomeRealmDiscoveryPolicyClient.Client)

	servicePrincipalLicenseDetailClient, err := serviceprincipallicensedetail.NewServicePrincipalLicenseDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalLicenseDetail client: %+v", err)
	}
	configureFunc(servicePrincipalLicenseDetailClient.Client)

	servicePrincipalMemberOfClient, err := serviceprincipalmemberof.NewServicePrincipalMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalMemberOf client: %+v", err)
	}
	configureFunc(servicePrincipalMemberOfClient.Client)

	servicePrincipalOauth2PermissionGrantClient, err := serviceprincipaloauth2permissiongrant.NewServicePrincipalOauth2PermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalOauth2PermissionGrant client: %+v", err)
	}
	configureFunc(servicePrincipalOauth2PermissionGrantClient.Client)

	servicePrincipalOwnedObjectClient, err := serviceprincipalownedobject.NewServicePrincipalOwnedObjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalOwnedObject client: %+v", err)
	}
	configureFunc(servicePrincipalOwnedObjectClient.Client)

	servicePrincipalOwnerClient, err := serviceprincipalowner.NewServicePrincipalOwnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalOwner client: %+v", err)
	}
	configureFunc(servicePrincipalOwnerClient.Client)

	servicePrincipalSynchronizationClient, err := serviceprincipalsynchronization.NewServicePrincipalSynchronizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalSynchronization client: %+v", err)
	}
	configureFunc(servicePrincipalSynchronizationClient.Client)

	servicePrincipalSynchronizationJobBulkUploadClient, err := serviceprincipalsynchronizationjobbulkupload.NewServicePrincipalSynchronizationJobBulkUploadClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalSynchronizationJobBulkUpload client: %+v", err)
	}
	configureFunc(servicePrincipalSynchronizationJobBulkUploadClient.Client)

	servicePrincipalSynchronizationJobClient, err := serviceprincipalsynchronizationjob.NewServicePrincipalSynchronizationJobClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalSynchronizationJob client: %+v", err)
	}
	configureFunc(servicePrincipalSynchronizationJobClient.Client)

	servicePrincipalSynchronizationJobSchemaClient, err := serviceprincipalsynchronizationjobschema.NewServicePrincipalSynchronizationJobSchemaClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalSynchronizationJobSchema client: %+v", err)
	}
	configureFunc(servicePrincipalSynchronizationJobSchemaClient.Client)

	servicePrincipalSynchronizationJobSchemaDirectoryClient, err := serviceprincipalsynchronizationjobschemadirectory.NewServicePrincipalSynchronizationJobSchemaDirectoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalSynchronizationJobSchemaDirectory client: %+v", err)
	}
	configureFunc(servicePrincipalSynchronizationJobSchemaDirectoryClient.Client)

	servicePrincipalSynchronizationSecretClient, err := serviceprincipalsynchronizationsecret.NewServicePrincipalSynchronizationSecretClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalSynchronizationSecret client: %+v", err)
	}
	configureFunc(servicePrincipalSynchronizationSecretClient.Client)

	servicePrincipalSynchronizationTemplateClient, err := serviceprincipalsynchronizationtemplate.NewServicePrincipalSynchronizationTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalSynchronizationTemplate client: %+v", err)
	}
	configureFunc(servicePrincipalSynchronizationTemplateClient.Client)

	servicePrincipalSynchronizationTemplateSchemaClient, err := serviceprincipalsynchronizationtemplateschema.NewServicePrincipalSynchronizationTemplateSchemaClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalSynchronizationTemplateSchema client: %+v", err)
	}
	configureFunc(servicePrincipalSynchronizationTemplateSchemaClient.Client)

	servicePrincipalSynchronizationTemplateSchemaDirectoryClient, err := serviceprincipalsynchronizationtemplateschemadirectory.NewServicePrincipalSynchronizationTemplateSchemaDirectoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalSynchronizationTemplateSchemaDirectory client: %+v", err)
	}
	configureFunc(servicePrincipalSynchronizationTemplateSchemaDirectoryClient.Client)

	servicePrincipalTokenIssuancePolicyClient, err := serviceprincipaltokenissuancepolicy.NewServicePrincipalTokenIssuancePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalTokenIssuancePolicy client: %+v", err)
	}
	configureFunc(servicePrincipalTokenIssuancePolicyClient.Client)

	servicePrincipalTokenLifetimePolicyClient, err := serviceprincipaltokenlifetimepolicy.NewServicePrincipalTokenLifetimePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalTokenLifetimePolicy client: %+v", err)
	}
	configureFunc(servicePrincipalTokenLifetimePolicyClient.Client)

	servicePrincipalTransitiveMemberOfClient, err := serviceprincipaltransitivememberof.NewServicePrincipalTransitiveMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalTransitiveMemberOf client: %+v", err)
	}
	configureFunc(servicePrincipalTransitiveMemberOfClient.Client)

	return &Client{
		ServicePrincipal:                                       servicePrincipalClient,
		ServicePrincipalAppManagementPolicy:                    servicePrincipalAppManagementPolicyClient,
		ServicePrincipalAppRoleAssignedTo:                      servicePrincipalAppRoleAssignedToClient,
		ServicePrincipalAppRoleAssignment:                      servicePrincipalAppRoleAssignmentClient,
		ServicePrincipalClaimsMappingPolicy:                    servicePrincipalClaimsMappingPolicyClient,
		ServicePrincipalCreatedObject:                          servicePrincipalCreatedObjectClient,
		ServicePrincipalDelegatedPermissionClassification:      servicePrincipalDelegatedPermissionClassificationClient,
		ServicePrincipalEndpoint:                               servicePrincipalEndpointClient,
		ServicePrincipalFederatedIdentityCredential:            servicePrincipalFederatedIdentityCredentialClient,
		ServicePrincipalHomeRealmDiscoveryPolicy:               servicePrincipalHomeRealmDiscoveryPolicyClient,
		ServicePrincipalLicenseDetail:                          servicePrincipalLicenseDetailClient,
		ServicePrincipalMemberOf:                               servicePrincipalMemberOfClient,
		ServicePrincipalOauth2PermissionGrant:                  servicePrincipalOauth2PermissionGrantClient,
		ServicePrincipalOwnedObject:                            servicePrincipalOwnedObjectClient,
		ServicePrincipalOwner:                                  servicePrincipalOwnerClient,
		ServicePrincipalSynchronization:                        servicePrincipalSynchronizationClient,
		ServicePrincipalSynchronizationJob:                     servicePrincipalSynchronizationJobClient,
		ServicePrincipalSynchronizationJobBulkUpload:           servicePrincipalSynchronizationJobBulkUploadClient,
		ServicePrincipalSynchronizationJobSchema:               servicePrincipalSynchronizationJobSchemaClient,
		ServicePrincipalSynchronizationJobSchemaDirectory:      servicePrincipalSynchronizationJobSchemaDirectoryClient,
		ServicePrincipalSynchronizationSecret:                  servicePrincipalSynchronizationSecretClient,
		ServicePrincipalSynchronizationTemplate:                servicePrincipalSynchronizationTemplateClient,
		ServicePrincipalSynchronizationTemplateSchema:          servicePrincipalSynchronizationTemplateSchemaClient,
		ServicePrincipalSynchronizationTemplateSchemaDirectory: servicePrincipalSynchronizationTemplateSchemaDirectoryClient,
		ServicePrincipalTokenIssuancePolicy:                    servicePrincipalTokenIssuancePolicyClient,
		ServicePrincipalTokenLifetimePolicy:                    servicePrincipalTokenLifetimePolicyClient,
		ServicePrincipalTransitiveMemberOf:                     servicePrincipalTransitiveMemberOfClient,
	}, nil
}
