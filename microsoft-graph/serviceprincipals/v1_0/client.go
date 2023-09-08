package v1_0

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalappmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalapproleassignedto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalapproleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalclaimsmappingpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalcreatedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipaldelegatedpermissionclassification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalendpoint"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalfederatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalhomerealmdiscoverypolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalmemberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipaloauth2permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalownedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalowner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalsynchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalsynchronizationjob"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalsynchronizationjobschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalsynchronizationjobschemadirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalsynchronizationsecret"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalsynchronizationtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalsynchronizationtemplateschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipalsynchronizationtemplateschemadirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipaltokenissuancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipaltokenlifetimepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/v1_0/serviceprincipaltransitivememberof"
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
	ServicePrincipalMemberOf                               *serviceprincipalmemberof.ServicePrincipalMemberOfClient
	ServicePrincipalOauth2PermissionGrant                  *serviceprincipaloauth2permissiongrant.ServicePrincipalOauth2PermissionGrantClient
	ServicePrincipalOwnedObject                            *serviceprincipalownedobject.ServicePrincipalOwnedObjectClient
	ServicePrincipalOwner                                  *serviceprincipalowner.ServicePrincipalOwnerClient
	ServicePrincipalSynchronization                        *serviceprincipalsynchronization.ServicePrincipalSynchronizationClient
	ServicePrincipalSynchronizationJob                     *serviceprincipalsynchronizationjob.ServicePrincipalSynchronizationJobClient
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
		ServicePrincipalMemberOf:                               servicePrincipalMemberOfClient,
		ServicePrincipalOauth2PermissionGrant:                  servicePrincipalOauth2PermissionGrantClient,
		ServicePrincipalOwnedObject:                            servicePrincipalOwnedObjectClient,
		ServicePrincipalOwner:                                  servicePrincipalOwnerClient,
		ServicePrincipalSynchronization:                        servicePrincipalSynchronizationClient,
		ServicePrincipalSynchronizationJob:                     servicePrincipalSynchronizationJobClient,
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
