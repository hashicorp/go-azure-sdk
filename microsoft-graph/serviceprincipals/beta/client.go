package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/appmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/approleassignedto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/approleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/claimsmappingpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/claimspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/createdobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/delegatedpermissionclassification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/endpoint"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/federatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/homerealmdiscoverypolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/licensedetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/memberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/oauth2permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/ownedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/owner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/permissiongrantpreapprovalpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/remotedesktopsecurityconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/remotedesktopsecurityconfigurationapprovedclientapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/remotedesktopsecurityconfigurationtargetdevicegroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronizationjob"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronizationjobbulkupload"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronizationjobschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronizationjobschemadirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronizationsecret"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronizationtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronizationtemplateschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronizationtemplateschemadirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/tokenissuancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/tokenlifetimepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/transitivememberof"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AppManagementPolicy                                 *appmanagementpolicy.AppManagementPolicyClient
	AppRoleAssignedTo                                   *approleassignedto.AppRoleAssignedToClient
	AppRoleAssignment                                   *approleassignment.AppRoleAssignmentClient
	ClaimsMappingPolicy                                 *claimsmappingpolicy.ClaimsMappingPolicyClient
	ClaimsPolicy                                        *claimspolicy.ClaimsPolicyClient
	CreatedObject                                       *createdobject.CreatedObjectClient
	DelegatedPermissionClassification                   *delegatedpermissionclassification.DelegatedPermissionClassificationClient
	Endpoint                                            *endpoint.EndpointClient
	FederatedIdentityCredential                         *federatedidentitycredential.FederatedIdentityCredentialClient
	HomeRealmDiscoveryPolicy                            *homerealmdiscoverypolicy.HomeRealmDiscoveryPolicyClient
	LicenseDetail                                       *licensedetail.LicenseDetailClient
	MemberOf                                            *memberof.MemberOfClient
	OAuth2PermissionGrant                               *oauth2permissiongrant.OAuth2PermissionGrantClient
	OwnedObject                                         *ownedobject.OwnedObjectClient
	Owner                                               *owner.OwnerClient
	PermissionGrantPreApprovalPolicy                    *permissiongrantpreapprovalpolicy.PermissionGrantPreApprovalPolicyClient
	RemoteDesktopSecurityConfiguration                  *remotedesktopsecurityconfiguration.RemoteDesktopSecurityConfigurationClient
	RemoteDesktopSecurityConfigurationApprovedClientApp *remotedesktopsecurityconfigurationapprovedclientapp.RemoteDesktopSecurityConfigurationApprovedClientAppClient
	RemoteDesktopSecurityConfigurationTargetDeviceGroup *remotedesktopsecurityconfigurationtargetdevicegroup.RemoteDesktopSecurityConfigurationTargetDeviceGroupClient
	ServicePrincipal                                    *serviceprincipal.ServicePrincipalClient
	Synchronization                                     *synchronization.SynchronizationClient
	SynchronizationJob                                  *synchronizationjob.SynchronizationJobClient
	SynchronizationJobBulkUpload                        *synchronizationjobbulkupload.SynchronizationJobBulkUploadClient
	SynchronizationJobSchema                            *synchronizationjobschema.SynchronizationJobSchemaClient
	SynchronizationJobSchemaDirectory                   *synchronizationjobschemadirectory.SynchronizationJobSchemaDirectoryClient
	SynchronizationSecret                               *synchronizationsecret.SynchronizationSecretClient
	SynchronizationTemplate                             *synchronizationtemplate.SynchronizationTemplateClient
	SynchronizationTemplateSchema                       *synchronizationtemplateschema.SynchronizationTemplateSchemaClient
	SynchronizationTemplateSchemaDirectory              *synchronizationtemplateschemadirectory.SynchronizationTemplateSchemaDirectoryClient
	TokenIssuancePolicy                                 *tokenissuancepolicy.TokenIssuancePolicyClient
	TokenLifetimePolicy                                 *tokenlifetimepolicy.TokenLifetimePolicyClient
	TransitiveMemberOf                                  *transitivememberof.TransitiveMemberOfClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	appManagementPolicyClient, err := appmanagementpolicy.NewAppManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppManagementPolicy client: %+v", err)
	}
	configureFunc(appManagementPolicyClient.Client)

	appRoleAssignedToClient, err := approleassignedto.NewAppRoleAssignedToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppRoleAssignedTo client: %+v", err)
	}
	configureFunc(appRoleAssignedToClient.Client)

	appRoleAssignmentClient, err := approleassignment.NewAppRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppRoleAssignment client: %+v", err)
	}
	configureFunc(appRoleAssignmentClient.Client)

	claimsMappingPolicyClient, err := claimsmappingpolicy.NewClaimsMappingPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ClaimsMappingPolicy client: %+v", err)
	}
	configureFunc(claimsMappingPolicyClient.Client)

	claimsPolicyClient, err := claimspolicy.NewClaimsPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ClaimsPolicy client: %+v", err)
	}
	configureFunc(claimsPolicyClient.Client)

	createdObjectClient, err := createdobject.NewCreatedObjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CreatedObject client: %+v", err)
	}
	configureFunc(createdObjectClient.Client)

	delegatedPermissionClassificationClient, err := delegatedpermissionclassification.NewDelegatedPermissionClassificationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DelegatedPermissionClassification client: %+v", err)
	}
	configureFunc(delegatedPermissionClassificationClient.Client)

	endpointClient, err := endpoint.NewEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Endpoint client: %+v", err)
	}
	configureFunc(endpointClient.Client)

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

	licenseDetailClient, err := licensedetail.NewLicenseDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LicenseDetail client: %+v", err)
	}
	configureFunc(licenseDetailClient.Client)

	memberOfClient, err := memberof.NewMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MemberOf client: %+v", err)
	}
	configureFunc(memberOfClient.Client)

	oAuth2PermissionGrantClient, err := oauth2permissiongrant.NewOAuth2PermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OAuth2PermissionGrant client: %+v", err)
	}
	configureFunc(oAuth2PermissionGrantClient.Client)

	ownedObjectClient, err := ownedobject.NewOwnedObjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OwnedObject client: %+v", err)
	}
	configureFunc(ownedObjectClient.Client)

	ownerClient, err := owner.NewOwnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Owner client: %+v", err)
	}
	configureFunc(ownerClient.Client)

	permissionGrantPreApprovalPolicyClient, err := permissiongrantpreapprovalpolicy.NewPermissionGrantPreApprovalPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PermissionGrantPreApprovalPolicy client: %+v", err)
	}
	configureFunc(permissionGrantPreApprovalPolicyClient.Client)

	remoteDesktopSecurityConfigurationApprovedClientAppClient, err := remotedesktopsecurityconfigurationapprovedclientapp.NewRemoteDesktopSecurityConfigurationApprovedClientAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RemoteDesktopSecurityConfigurationApprovedClientApp client: %+v", err)
	}
	configureFunc(remoteDesktopSecurityConfigurationApprovedClientAppClient.Client)

	remoteDesktopSecurityConfigurationClient, err := remotedesktopsecurityconfiguration.NewRemoteDesktopSecurityConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RemoteDesktopSecurityConfiguration client: %+v", err)
	}
	configureFunc(remoteDesktopSecurityConfigurationClient.Client)

	remoteDesktopSecurityConfigurationTargetDeviceGroupClient, err := remotedesktopsecurityconfigurationtargetdevicegroup.NewRemoteDesktopSecurityConfigurationTargetDeviceGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RemoteDesktopSecurityConfigurationTargetDeviceGroup client: %+v", err)
	}
	configureFunc(remoteDesktopSecurityConfigurationTargetDeviceGroupClient.Client)

	servicePrincipalClient, err := serviceprincipal.NewServicePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipal client: %+v", err)
	}
	configureFunc(servicePrincipalClient.Client)

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

	transitiveMemberOfClient, err := transitivememberof.NewTransitiveMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TransitiveMemberOf client: %+v", err)
	}
	configureFunc(transitiveMemberOfClient.Client)

	return &Client{
		AppManagementPolicy:                appManagementPolicyClient,
		AppRoleAssignedTo:                  appRoleAssignedToClient,
		AppRoleAssignment:                  appRoleAssignmentClient,
		ClaimsMappingPolicy:                claimsMappingPolicyClient,
		ClaimsPolicy:                       claimsPolicyClient,
		CreatedObject:                      createdObjectClient,
		DelegatedPermissionClassification:  delegatedPermissionClassificationClient,
		Endpoint:                           endpointClient,
		FederatedIdentityCredential:        federatedIdentityCredentialClient,
		HomeRealmDiscoveryPolicy:           homeRealmDiscoveryPolicyClient,
		LicenseDetail:                      licenseDetailClient,
		MemberOf:                           memberOfClient,
		OAuth2PermissionGrant:              oAuth2PermissionGrantClient,
		OwnedObject:                        ownedObjectClient,
		Owner:                              ownerClient,
		PermissionGrantPreApprovalPolicy:   permissionGrantPreApprovalPolicyClient,
		RemoteDesktopSecurityConfiguration: remoteDesktopSecurityConfigurationClient,
		RemoteDesktopSecurityConfigurationApprovedClientApp: remoteDesktopSecurityConfigurationApprovedClientAppClient,
		RemoteDesktopSecurityConfigurationTargetDeviceGroup: remoteDesktopSecurityConfigurationTargetDeviceGroupClient,
		ServicePrincipal:                       servicePrincipalClient,
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
		TransitiveMemberOf:                     transitiveMemberOfClient,
	}, nil
}
