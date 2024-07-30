package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/accessreviewpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/activitybasedtimeoutpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/activitybasedtimeoutpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/adminconsentrequestpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/appmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/appmanagementpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/authenticationflowspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/authenticationmethodspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/authenticationmethodspolicyauthenticationmethodconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/authenticationstrengthpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/authenticationstrengthpolicycombinationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/authorizationpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/authorizationpolicydefaultuserroleoverride"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/b2cauthenticationmethodspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/claimsmappingpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/claimsmappingpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/conditionalaccesspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/crosstenantaccesspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/crosstenantaccesspolicydefault"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/crosstenantaccesspolicypartner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/crosstenantaccesspolicypartneridentitysynchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/crosstenantaccesspolicytemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/crosstenantaccesspolicytemplatemultitenantorganizationidentitysynchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/crosstenantaccesspolicytemplatemultitenantorganizationpartnerconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/defaultappmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/deviceregistrationpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/directoryroleaccessreviewpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/externalidentitiespolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/featurerolloutpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/featurerolloutpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/federatedtokenvalidationpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/homerealmdiscoverypolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/homerealmdiscoverypolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/identitysecuritydefaultsenforcementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/mobileappmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/mobileappmanagementpolicyincludedgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/mobileappmanagementpolicyincludedgroupserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/mobiledevicemanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/mobiledevicemanagementpolicyincludedgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/mobiledevicemanagementpolicyincludedgroupserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/permissiongrantpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/permissiongrantpolicyexclude"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/permissiongrantpolicyinclude"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/rolemanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/rolemanagementpolicyassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/rolemanagementpolicyassignmentpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/rolemanagementpolicyeffectiverule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/rolemanagementpolicyrule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/serviceprincipalcreationpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/serviceprincipalcreationpolicyexclude"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/serviceprincipalcreationpolicyinclude"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/tokenissuancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/tokenissuancepolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/tokenlifetimepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/tokenlifetimepolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AccessReviewPolicy                                                            *accessreviewpolicy.AccessReviewPolicyClient
	ActivityBasedTimeoutPolicy                                                    *activitybasedtimeoutpolicy.ActivityBasedTimeoutPolicyClient
	ActivityBasedTimeoutPolicyAppliesTo                                           *activitybasedtimeoutpolicyappliesto.ActivityBasedTimeoutPolicyAppliesToClient
	AdminConsentRequestPolicy                                                     *adminconsentrequestpolicy.AdminConsentRequestPolicyClient
	AppManagementPolicy                                                           *appmanagementpolicy.AppManagementPolicyClient
	AppManagementPolicyAppliesTo                                                  *appmanagementpolicyappliesto.AppManagementPolicyAppliesToClient
	AuthenticationFlowsPolicy                                                     *authenticationflowspolicy.AuthenticationFlowsPolicyClient
	AuthenticationMethodsPolicy                                                   *authenticationmethodspolicy.AuthenticationMethodsPolicyClient
	AuthenticationMethodsPolicyAuthenticationMethodConfiguration                  *authenticationmethodspolicyauthenticationmethodconfiguration.AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient
	AuthenticationStrengthPolicy                                                  *authenticationstrengthpolicy.AuthenticationStrengthPolicyClient
	AuthenticationStrengthPolicyCombinationConfiguration                          *authenticationstrengthpolicycombinationconfiguration.AuthenticationStrengthPolicyCombinationConfigurationClient
	AuthorizationPolicy                                                           *authorizationpolicy.AuthorizationPolicyClient
	AuthorizationPolicyDefaultUserRoleOverride                                    *authorizationpolicydefaultuserroleoverride.AuthorizationPolicyDefaultUserRoleOverrideClient
	B2cAuthenticationMethodsPolicy                                                *b2cauthenticationmethodspolicy.B2cAuthenticationMethodsPolicyClient
	ClaimsMappingPolicy                                                           *claimsmappingpolicy.ClaimsMappingPolicyClient
	ClaimsMappingPolicyAppliesTo                                                  *claimsmappingpolicyappliesto.ClaimsMappingPolicyAppliesToClient
	ConditionalAccessPolicy                                                       *conditionalaccesspolicy.ConditionalAccessPolicyClient
	CrossTenantAccessPolicy                                                       *crosstenantaccesspolicy.CrossTenantAccessPolicyClient
	CrossTenantAccessPolicyDefault                                                *crosstenantaccesspolicydefault.CrossTenantAccessPolicyDefaultClient
	CrossTenantAccessPolicyPartner                                                *crosstenantaccesspolicypartner.CrossTenantAccessPolicyPartnerClient
	CrossTenantAccessPolicyPartnerIdentitySynchronization                         *crosstenantaccesspolicypartneridentitysynchronization.CrossTenantAccessPolicyPartnerIdentitySynchronizationClient
	CrossTenantAccessPolicyTemplate                                               *crosstenantaccesspolicytemplate.CrossTenantAccessPolicyTemplateClient
	CrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronization *crosstenantaccesspolicytemplatemultitenantorganizationidentitysynchronization.CrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClient
	CrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfiguration    *crosstenantaccesspolicytemplatemultitenantorganizationpartnerconfiguration.CrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationClient
	DefaultAppManagementPolicy                                                    *defaultappmanagementpolicy.DefaultAppManagementPolicyClient
	DeviceRegistrationPolicy                                                      *deviceregistrationpolicy.DeviceRegistrationPolicyClient
	DirectoryRoleAccessReviewPolicy                                               *directoryroleaccessreviewpolicy.DirectoryRoleAccessReviewPolicyClient
	ExternalIdentitiesPolicy                                                      *externalidentitiespolicy.ExternalIdentitiesPolicyClient
	FeatureRolloutPolicy                                                          *featurerolloutpolicy.FeatureRolloutPolicyClient
	FeatureRolloutPolicyAppliesTo                                                 *featurerolloutpolicyappliesto.FeatureRolloutPolicyAppliesToClient
	FederatedTokenValidationPolicy                                                *federatedtokenvalidationpolicy.FederatedTokenValidationPolicyClient
	HomeRealmDiscoveryPolicy                                                      *homerealmdiscoverypolicy.HomeRealmDiscoveryPolicyClient
	HomeRealmDiscoveryPolicyAppliesTo                                             *homerealmdiscoverypolicyappliesto.HomeRealmDiscoveryPolicyAppliesToClient
	IdentitySecurityDefaultsEnforcementPolicy                                     *identitysecuritydefaultsenforcementpolicy.IdentitySecurityDefaultsEnforcementPolicyClient
	MobileAppManagementPolicy                                                     *mobileappmanagementpolicy.MobileAppManagementPolicyClient
	MobileAppManagementPolicyIncludedGroup                                        *mobileappmanagementpolicyincludedgroup.MobileAppManagementPolicyIncludedGroupClient
	MobileAppManagementPolicyIncludedGroupServiceProvisioningError                *mobileappmanagementpolicyincludedgroupserviceprovisioningerror.MobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClient
	MobileDeviceManagementPolicy                                                  *mobiledevicemanagementpolicy.MobileDeviceManagementPolicyClient
	MobileDeviceManagementPolicyIncludedGroup                                     *mobiledevicemanagementpolicyincludedgroup.MobileDeviceManagementPolicyIncludedGroupClient
	MobileDeviceManagementPolicyIncludedGroupServiceProvisioningError             *mobiledevicemanagementpolicyincludedgroupserviceprovisioningerror.MobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClient
	PermissionGrantPolicy                                                         *permissiongrantpolicy.PermissionGrantPolicyClient
	PermissionGrantPolicyExclude                                                  *permissiongrantpolicyexclude.PermissionGrantPolicyExcludeClient
	PermissionGrantPolicyInclude                                                  *permissiongrantpolicyinclude.PermissionGrantPolicyIncludeClient
	Policy                                                                        *policy.PolicyClient
	RoleManagementPolicy                                                          *rolemanagementpolicy.RoleManagementPolicyClient
	RoleManagementPolicyAssignment                                                *rolemanagementpolicyassignment.RoleManagementPolicyAssignmentClient
	RoleManagementPolicyAssignmentPolicy                                          *rolemanagementpolicyassignmentpolicy.RoleManagementPolicyAssignmentPolicyClient
	RoleManagementPolicyEffectiveRule                                             *rolemanagementpolicyeffectiverule.RoleManagementPolicyEffectiveRuleClient
	RoleManagementPolicyRule                                                      *rolemanagementpolicyrule.RoleManagementPolicyRuleClient
	ServicePrincipalCreationPolicy                                                *serviceprincipalcreationpolicy.ServicePrincipalCreationPolicyClient
	ServicePrincipalCreationPolicyExclude                                         *serviceprincipalcreationpolicyexclude.ServicePrincipalCreationPolicyExcludeClient
	ServicePrincipalCreationPolicyInclude                                         *serviceprincipalcreationpolicyinclude.ServicePrincipalCreationPolicyIncludeClient
	TokenIssuancePolicy                                                           *tokenissuancepolicy.TokenIssuancePolicyClient
	TokenIssuancePolicyAppliesTo                                                  *tokenissuancepolicyappliesto.TokenIssuancePolicyAppliesToClient
	TokenLifetimePolicy                                                           *tokenlifetimepolicy.TokenLifetimePolicyClient
	TokenLifetimePolicyAppliesTo                                                  *tokenlifetimepolicyappliesto.TokenLifetimePolicyAppliesToClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	accessReviewPolicyClient, err := accessreviewpolicy.NewAccessReviewPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessReviewPolicy client: %+v", err)
	}
	configureFunc(accessReviewPolicyClient.Client)

	activityBasedTimeoutPolicyAppliesToClient, err := activitybasedtimeoutpolicyappliesto.NewActivityBasedTimeoutPolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ActivityBasedTimeoutPolicyAppliesTo client: %+v", err)
	}
	configureFunc(activityBasedTimeoutPolicyAppliesToClient.Client)

	activityBasedTimeoutPolicyClient, err := activitybasedtimeoutpolicy.NewActivityBasedTimeoutPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ActivityBasedTimeoutPolicy client: %+v", err)
	}
	configureFunc(activityBasedTimeoutPolicyClient.Client)

	adminConsentRequestPolicyClient, err := adminconsentrequestpolicy.NewAdminConsentRequestPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdminConsentRequestPolicy client: %+v", err)
	}
	configureFunc(adminConsentRequestPolicyClient.Client)

	appManagementPolicyAppliesToClient, err := appmanagementpolicyappliesto.NewAppManagementPolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppManagementPolicyAppliesTo client: %+v", err)
	}
	configureFunc(appManagementPolicyAppliesToClient.Client)

	appManagementPolicyClient, err := appmanagementpolicy.NewAppManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppManagementPolicy client: %+v", err)
	}
	configureFunc(appManagementPolicyClient.Client)

	authenticationFlowsPolicyClient, err := authenticationflowspolicy.NewAuthenticationFlowsPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationFlowsPolicy client: %+v", err)
	}
	configureFunc(authenticationFlowsPolicyClient.Client)

	authenticationMethodsPolicyAuthenticationMethodConfigurationClient, err := authenticationmethodspolicyauthenticationmethodconfiguration.NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationMethodsPolicyAuthenticationMethodConfiguration client: %+v", err)
	}
	configureFunc(authenticationMethodsPolicyAuthenticationMethodConfigurationClient.Client)

	authenticationMethodsPolicyClient, err := authenticationmethodspolicy.NewAuthenticationMethodsPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationMethodsPolicy client: %+v", err)
	}
	configureFunc(authenticationMethodsPolicyClient.Client)

	authenticationStrengthPolicyClient, err := authenticationstrengthpolicy.NewAuthenticationStrengthPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationStrengthPolicy client: %+v", err)
	}
	configureFunc(authenticationStrengthPolicyClient.Client)

	authenticationStrengthPolicyCombinationConfigurationClient, err := authenticationstrengthpolicycombinationconfiguration.NewAuthenticationStrengthPolicyCombinationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationStrengthPolicyCombinationConfiguration client: %+v", err)
	}
	configureFunc(authenticationStrengthPolicyCombinationConfigurationClient.Client)

	authorizationPolicyClient, err := authorizationpolicy.NewAuthorizationPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationPolicy client: %+v", err)
	}
	configureFunc(authorizationPolicyClient.Client)

	authorizationPolicyDefaultUserRoleOverrideClient, err := authorizationpolicydefaultuserroleoverride.NewAuthorizationPolicyDefaultUserRoleOverrideClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationPolicyDefaultUserRoleOverride client: %+v", err)
	}
	configureFunc(authorizationPolicyDefaultUserRoleOverrideClient.Client)

	b2cAuthenticationMethodsPolicyClient, err := b2cauthenticationmethodspolicy.NewB2cAuthenticationMethodsPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building B2cAuthenticationMethodsPolicy client: %+v", err)
	}
	configureFunc(b2cAuthenticationMethodsPolicyClient.Client)

	claimsMappingPolicyAppliesToClient, err := claimsmappingpolicyappliesto.NewClaimsMappingPolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ClaimsMappingPolicyAppliesTo client: %+v", err)
	}
	configureFunc(claimsMappingPolicyAppliesToClient.Client)

	claimsMappingPolicyClient, err := claimsmappingpolicy.NewClaimsMappingPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ClaimsMappingPolicy client: %+v", err)
	}
	configureFunc(claimsMappingPolicyClient.Client)

	conditionalAccessPolicyClient, err := conditionalaccesspolicy.NewConditionalAccessPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccessPolicy client: %+v", err)
	}
	configureFunc(conditionalAccessPolicyClient.Client)

	crossTenantAccessPolicyClient, err := crosstenantaccesspolicy.NewCrossTenantAccessPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CrossTenantAccessPolicy client: %+v", err)
	}
	configureFunc(crossTenantAccessPolicyClient.Client)

	crossTenantAccessPolicyDefaultClient, err := crosstenantaccesspolicydefault.NewCrossTenantAccessPolicyDefaultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CrossTenantAccessPolicyDefault client: %+v", err)
	}
	configureFunc(crossTenantAccessPolicyDefaultClient.Client)

	crossTenantAccessPolicyPartnerClient, err := crosstenantaccesspolicypartner.NewCrossTenantAccessPolicyPartnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CrossTenantAccessPolicyPartner client: %+v", err)
	}
	configureFunc(crossTenantAccessPolicyPartnerClient.Client)

	crossTenantAccessPolicyPartnerIdentitySynchronizationClient, err := crosstenantaccesspolicypartneridentitysynchronization.NewCrossTenantAccessPolicyPartnerIdentitySynchronizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CrossTenantAccessPolicyPartnerIdentitySynchronization client: %+v", err)
	}
	configureFunc(crossTenantAccessPolicyPartnerIdentitySynchronizationClient.Client)

	crossTenantAccessPolicyTemplateClient, err := crosstenantaccesspolicytemplate.NewCrossTenantAccessPolicyTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CrossTenantAccessPolicyTemplate client: %+v", err)
	}
	configureFunc(crossTenantAccessPolicyTemplateClient.Client)

	crossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClient, err := crosstenantaccesspolicytemplatemultitenantorganizationidentitysynchronization.NewCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronization client: %+v", err)
	}
	configureFunc(crossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClient.Client)

	crossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationClient, err := crosstenantaccesspolicytemplatemultitenantorganizationpartnerconfiguration.NewCrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfiguration client: %+v", err)
	}
	configureFunc(crossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationClient.Client)

	defaultAppManagementPolicyClient, err := defaultappmanagementpolicy.NewDefaultAppManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefaultAppManagementPolicy client: %+v", err)
	}
	configureFunc(defaultAppManagementPolicyClient.Client)

	deviceRegistrationPolicyClient, err := deviceregistrationpolicy.NewDeviceRegistrationPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceRegistrationPolicy client: %+v", err)
	}
	configureFunc(deviceRegistrationPolicyClient.Client)

	directoryRoleAccessReviewPolicyClient, err := directoryroleaccessreviewpolicy.NewDirectoryRoleAccessReviewPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAccessReviewPolicy client: %+v", err)
	}
	configureFunc(directoryRoleAccessReviewPolicyClient.Client)

	externalIdentitiesPolicyClient, err := externalidentitiespolicy.NewExternalIdentitiesPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExternalIdentitiesPolicy client: %+v", err)
	}
	configureFunc(externalIdentitiesPolicyClient.Client)

	featureRolloutPolicyAppliesToClient, err := featurerolloutpolicyappliesto.NewFeatureRolloutPolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FeatureRolloutPolicyAppliesTo client: %+v", err)
	}
	configureFunc(featureRolloutPolicyAppliesToClient.Client)

	featureRolloutPolicyClient, err := featurerolloutpolicy.NewFeatureRolloutPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FeatureRolloutPolicy client: %+v", err)
	}
	configureFunc(featureRolloutPolicyClient.Client)

	federatedTokenValidationPolicyClient, err := federatedtokenvalidationpolicy.NewFederatedTokenValidationPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FederatedTokenValidationPolicy client: %+v", err)
	}
	configureFunc(federatedTokenValidationPolicyClient.Client)

	homeRealmDiscoveryPolicyAppliesToClient, err := homerealmdiscoverypolicyappliesto.NewHomeRealmDiscoveryPolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HomeRealmDiscoveryPolicyAppliesTo client: %+v", err)
	}
	configureFunc(homeRealmDiscoveryPolicyAppliesToClient.Client)

	homeRealmDiscoveryPolicyClient, err := homerealmdiscoverypolicy.NewHomeRealmDiscoveryPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HomeRealmDiscoveryPolicy client: %+v", err)
	}
	configureFunc(homeRealmDiscoveryPolicyClient.Client)

	identitySecurityDefaultsEnforcementPolicyClient, err := identitysecuritydefaultsenforcementpolicy.NewIdentitySecurityDefaultsEnforcementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentitySecurityDefaultsEnforcementPolicy client: %+v", err)
	}
	configureFunc(identitySecurityDefaultsEnforcementPolicyClient.Client)

	mobileAppManagementPolicyClient, err := mobileappmanagementpolicy.NewMobileAppManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileAppManagementPolicy client: %+v", err)
	}
	configureFunc(mobileAppManagementPolicyClient.Client)

	mobileAppManagementPolicyIncludedGroupClient, err := mobileappmanagementpolicyincludedgroup.NewMobileAppManagementPolicyIncludedGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileAppManagementPolicyIncludedGroup client: %+v", err)
	}
	configureFunc(mobileAppManagementPolicyIncludedGroupClient.Client)

	mobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClient, err := mobileappmanagementpolicyincludedgroupserviceprovisioningerror.NewMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileAppManagementPolicyIncludedGroupServiceProvisioningError client: %+v", err)
	}
	configureFunc(mobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClient.Client)

	mobileDeviceManagementPolicyClient, err := mobiledevicemanagementpolicy.NewMobileDeviceManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileDeviceManagementPolicy client: %+v", err)
	}
	configureFunc(mobileDeviceManagementPolicyClient.Client)

	mobileDeviceManagementPolicyIncludedGroupClient, err := mobiledevicemanagementpolicyincludedgroup.NewMobileDeviceManagementPolicyIncludedGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileDeviceManagementPolicyIncludedGroup client: %+v", err)
	}
	configureFunc(mobileDeviceManagementPolicyIncludedGroupClient.Client)

	mobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClient, err := mobiledevicemanagementpolicyincludedgroupserviceprovisioningerror.NewMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileDeviceManagementPolicyIncludedGroupServiceProvisioningError client: %+v", err)
	}
	configureFunc(mobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClient.Client)

	permissionGrantPolicyClient, err := permissiongrantpolicy.NewPermissionGrantPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PermissionGrantPolicy client: %+v", err)
	}
	configureFunc(permissionGrantPolicyClient.Client)

	permissionGrantPolicyExcludeClient, err := permissiongrantpolicyexclude.NewPermissionGrantPolicyExcludeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PermissionGrantPolicyExclude client: %+v", err)
	}
	configureFunc(permissionGrantPolicyExcludeClient.Client)

	permissionGrantPolicyIncludeClient, err := permissiongrantpolicyinclude.NewPermissionGrantPolicyIncludeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PermissionGrantPolicyInclude client: %+v", err)
	}
	configureFunc(permissionGrantPolicyIncludeClient.Client)

	policyClient, err := policy.NewPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Policy client: %+v", err)
	}
	configureFunc(policyClient.Client)

	roleManagementPolicyAssignmentClient, err := rolemanagementpolicyassignment.NewRoleManagementPolicyAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementPolicyAssignment client: %+v", err)
	}
	configureFunc(roleManagementPolicyAssignmentClient.Client)

	roleManagementPolicyAssignmentPolicyClient, err := rolemanagementpolicyassignmentpolicy.NewRoleManagementPolicyAssignmentPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementPolicyAssignmentPolicy client: %+v", err)
	}
	configureFunc(roleManagementPolicyAssignmentPolicyClient.Client)

	roleManagementPolicyClient, err := rolemanagementpolicy.NewRoleManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementPolicy client: %+v", err)
	}
	configureFunc(roleManagementPolicyClient.Client)

	roleManagementPolicyEffectiveRuleClient, err := rolemanagementpolicyeffectiverule.NewRoleManagementPolicyEffectiveRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementPolicyEffectiveRule client: %+v", err)
	}
	configureFunc(roleManagementPolicyEffectiveRuleClient.Client)

	roleManagementPolicyRuleClient, err := rolemanagementpolicyrule.NewRoleManagementPolicyRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementPolicyRule client: %+v", err)
	}
	configureFunc(roleManagementPolicyRuleClient.Client)

	servicePrincipalCreationPolicyClient, err := serviceprincipalcreationpolicy.NewServicePrincipalCreationPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalCreationPolicy client: %+v", err)
	}
	configureFunc(servicePrincipalCreationPolicyClient.Client)

	servicePrincipalCreationPolicyExcludeClient, err := serviceprincipalcreationpolicyexclude.NewServicePrincipalCreationPolicyExcludeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalCreationPolicyExclude client: %+v", err)
	}
	configureFunc(servicePrincipalCreationPolicyExcludeClient.Client)

	servicePrincipalCreationPolicyIncludeClient, err := serviceprincipalcreationpolicyinclude.NewServicePrincipalCreationPolicyIncludeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipalCreationPolicyInclude client: %+v", err)
	}
	configureFunc(servicePrincipalCreationPolicyIncludeClient.Client)

	tokenIssuancePolicyAppliesToClient, err := tokenissuancepolicyappliesto.NewTokenIssuancePolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TokenIssuancePolicyAppliesTo client: %+v", err)
	}
	configureFunc(tokenIssuancePolicyAppliesToClient.Client)

	tokenIssuancePolicyClient, err := tokenissuancepolicy.NewTokenIssuancePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TokenIssuancePolicy client: %+v", err)
	}
	configureFunc(tokenIssuancePolicyClient.Client)

	tokenLifetimePolicyAppliesToClient, err := tokenlifetimepolicyappliesto.NewTokenLifetimePolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TokenLifetimePolicyAppliesTo client: %+v", err)
	}
	configureFunc(tokenLifetimePolicyAppliesToClient.Client)

	tokenLifetimePolicyClient, err := tokenlifetimepolicy.NewTokenLifetimePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TokenLifetimePolicy client: %+v", err)
	}
	configureFunc(tokenLifetimePolicyClient.Client)

	return &Client{
		AccessReviewPolicy:                                           accessReviewPolicyClient,
		ActivityBasedTimeoutPolicy:                                   activityBasedTimeoutPolicyClient,
		ActivityBasedTimeoutPolicyAppliesTo:                          activityBasedTimeoutPolicyAppliesToClient,
		AdminConsentRequestPolicy:                                    adminConsentRequestPolicyClient,
		AppManagementPolicy:                                          appManagementPolicyClient,
		AppManagementPolicyAppliesTo:                                 appManagementPolicyAppliesToClient,
		AuthenticationFlowsPolicy:                                    authenticationFlowsPolicyClient,
		AuthenticationMethodsPolicy:                                  authenticationMethodsPolicyClient,
		AuthenticationMethodsPolicyAuthenticationMethodConfiguration: authenticationMethodsPolicyAuthenticationMethodConfigurationClient,
		AuthenticationStrengthPolicy:                                 authenticationStrengthPolicyClient,
		AuthenticationStrengthPolicyCombinationConfiguration:         authenticationStrengthPolicyCombinationConfigurationClient,
		AuthorizationPolicy:                                          authorizationPolicyClient,
		AuthorizationPolicyDefaultUserRoleOverride:                   authorizationPolicyDefaultUserRoleOverrideClient,
		B2cAuthenticationMethodsPolicy:                               b2cAuthenticationMethodsPolicyClient,
		ClaimsMappingPolicy:                                          claimsMappingPolicyClient,
		ClaimsMappingPolicyAppliesTo:                                 claimsMappingPolicyAppliesToClient,
		ConditionalAccessPolicy:                                      conditionalAccessPolicyClient,
		CrossTenantAccessPolicy:                                      crossTenantAccessPolicyClient,
		CrossTenantAccessPolicyDefault:                               crossTenantAccessPolicyDefaultClient,
		CrossTenantAccessPolicyPartner:                               crossTenantAccessPolicyPartnerClient,
		CrossTenantAccessPolicyPartnerIdentitySynchronization:        crossTenantAccessPolicyPartnerIdentitySynchronizationClient,
		CrossTenantAccessPolicyTemplate:                              crossTenantAccessPolicyTemplateClient,
		CrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronization: crossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClient,
		CrossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfiguration:    crossTenantAccessPolicyTemplateMultiTenantOrganizationPartnerConfigurationClient,
		DefaultAppManagementPolicy:                                        defaultAppManagementPolicyClient,
		DeviceRegistrationPolicy:                                          deviceRegistrationPolicyClient,
		DirectoryRoleAccessReviewPolicy:                                   directoryRoleAccessReviewPolicyClient,
		ExternalIdentitiesPolicy:                                          externalIdentitiesPolicyClient,
		FeatureRolloutPolicy:                                              featureRolloutPolicyClient,
		FeatureRolloutPolicyAppliesTo:                                     featureRolloutPolicyAppliesToClient,
		FederatedTokenValidationPolicy:                                    federatedTokenValidationPolicyClient,
		HomeRealmDiscoveryPolicy:                                          homeRealmDiscoveryPolicyClient,
		HomeRealmDiscoveryPolicyAppliesTo:                                 homeRealmDiscoveryPolicyAppliesToClient,
		IdentitySecurityDefaultsEnforcementPolicy:                         identitySecurityDefaultsEnforcementPolicyClient,
		MobileAppManagementPolicy:                                         mobileAppManagementPolicyClient,
		MobileAppManagementPolicyIncludedGroup:                            mobileAppManagementPolicyIncludedGroupClient,
		MobileAppManagementPolicyIncludedGroupServiceProvisioningError:    mobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClient,
		MobileDeviceManagementPolicy:                                      mobileDeviceManagementPolicyClient,
		MobileDeviceManagementPolicyIncludedGroup:                         mobileDeviceManagementPolicyIncludedGroupClient,
		MobileDeviceManagementPolicyIncludedGroupServiceProvisioningError: mobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClient,
		PermissionGrantPolicy:                                             permissionGrantPolicyClient,
		PermissionGrantPolicyExclude:                                      permissionGrantPolicyExcludeClient,
		PermissionGrantPolicyInclude:                                      permissionGrantPolicyIncludeClient,
		Policy:                                                            policyClient,
		RoleManagementPolicy:                                              roleManagementPolicyClient,
		RoleManagementPolicyAssignment:                                    roleManagementPolicyAssignmentClient,
		RoleManagementPolicyAssignmentPolicy:                              roleManagementPolicyAssignmentPolicyClient,
		RoleManagementPolicyEffectiveRule:                                 roleManagementPolicyEffectiveRuleClient,
		RoleManagementPolicyRule:                                          roleManagementPolicyRuleClient,
		ServicePrincipalCreationPolicy:                                    servicePrincipalCreationPolicyClient,
		ServicePrincipalCreationPolicyExclude:                             servicePrincipalCreationPolicyExcludeClient,
		ServicePrincipalCreationPolicyInclude:                             servicePrincipalCreationPolicyIncludeClient,
		TokenIssuancePolicy:                                               tokenIssuancePolicyClient,
		TokenIssuancePolicyAppliesTo:                                      tokenIssuancePolicyAppliesToClient,
		TokenLifetimePolicy:                                               tokenLifetimePolicyClient,
		TokenLifetimePolicyAppliesTo:                                      tokenLifetimePolicyAppliesToClient,
	}, nil
}
