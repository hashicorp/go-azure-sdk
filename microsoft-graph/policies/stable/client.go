package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/activitybasedtimeoutpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/activitybasedtimeoutpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/adminconsentrequestpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/appmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/appmanagementpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationflowspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationmethodspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationmethodspolicyauthenticationmethodconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationstrengthpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationstrengthpolicycombinationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authorizationpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/claimsmappingpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/claimsmappingpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/conditionalaccesspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/crosstenantaccesspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/crosstenantaccesspolicydefault"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/crosstenantaccesspolicypartner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/crosstenantaccesspolicypartneridentitysynchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/defaultappmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/featurerolloutpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/featurerolloutpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/homerealmdiscoverypolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/homerealmdiscoverypolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/identitysecuritydefaultsenforcementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/permissiongrantpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/permissiongrantpolicyexclude"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/permissiongrantpolicyinclude"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/policy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicyassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicyassignmentpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicyeffectiverule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicyrule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/tokenissuancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/tokenissuancepolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/tokenlifetimepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/tokenlifetimepolicyappliesto"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	ActivityBasedTimeoutPolicy                                   *activitybasedtimeoutpolicy.ActivityBasedTimeoutPolicyClient
	ActivityBasedTimeoutPolicyAppliesTo                          *activitybasedtimeoutpolicyappliesto.ActivityBasedTimeoutPolicyAppliesToClient
	AdminConsentRequestPolicy                                    *adminconsentrequestpolicy.AdminConsentRequestPolicyClient
	AppManagementPolicy                                          *appmanagementpolicy.AppManagementPolicyClient
	AppManagementPolicyAppliesTo                                 *appmanagementpolicyappliesto.AppManagementPolicyAppliesToClient
	AuthenticationFlowsPolicy                                    *authenticationflowspolicy.AuthenticationFlowsPolicyClient
	AuthenticationMethodsPolicy                                  *authenticationmethodspolicy.AuthenticationMethodsPolicyClient
	AuthenticationMethodsPolicyAuthenticationMethodConfiguration *authenticationmethodspolicyauthenticationmethodconfiguration.AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient
	AuthenticationStrengthPolicy                                 *authenticationstrengthpolicy.AuthenticationStrengthPolicyClient
	AuthenticationStrengthPolicyCombinationConfiguration         *authenticationstrengthpolicycombinationconfiguration.AuthenticationStrengthPolicyCombinationConfigurationClient
	AuthorizationPolicy                                          *authorizationpolicy.AuthorizationPolicyClient
	ClaimsMappingPolicy                                          *claimsmappingpolicy.ClaimsMappingPolicyClient
	ClaimsMappingPolicyAppliesTo                                 *claimsmappingpolicyappliesto.ClaimsMappingPolicyAppliesToClient
	ConditionalAccessPolicy                                      *conditionalaccesspolicy.ConditionalAccessPolicyClient
	CrossTenantAccessPolicy                                      *crosstenantaccesspolicy.CrossTenantAccessPolicyClient
	CrossTenantAccessPolicyDefault                               *crosstenantaccesspolicydefault.CrossTenantAccessPolicyDefaultClient
	CrossTenantAccessPolicyPartner                               *crosstenantaccesspolicypartner.CrossTenantAccessPolicyPartnerClient
	CrossTenantAccessPolicyPartnerIdentitySynchronization        *crosstenantaccesspolicypartneridentitysynchronization.CrossTenantAccessPolicyPartnerIdentitySynchronizationClient
	DefaultAppManagementPolicy                                   *defaultappmanagementpolicy.DefaultAppManagementPolicyClient
	FeatureRolloutPolicy                                         *featurerolloutpolicy.FeatureRolloutPolicyClient
	FeatureRolloutPolicyAppliesTo                                *featurerolloutpolicyappliesto.FeatureRolloutPolicyAppliesToClient
	HomeRealmDiscoveryPolicy                                     *homerealmdiscoverypolicy.HomeRealmDiscoveryPolicyClient
	HomeRealmDiscoveryPolicyAppliesTo                            *homerealmdiscoverypolicyappliesto.HomeRealmDiscoveryPolicyAppliesToClient
	IdentitySecurityDefaultsEnforcementPolicy                    *identitysecuritydefaultsenforcementpolicy.IdentitySecurityDefaultsEnforcementPolicyClient
	PermissionGrantPolicy                                        *permissiongrantpolicy.PermissionGrantPolicyClient
	PermissionGrantPolicyExclude                                 *permissiongrantpolicyexclude.PermissionGrantPolicyExcludeClient
	PermissionGrantPolicyInclude                                 *permissiongrantpolicyinclude.PermissionGrantPolicyIncludeClient
	Policy                                                       *policy.PolicyClient
	RoleManagementPolicy                                         *rolemanagementpolicy.RoleManagementPolicyClient
	RoleManagementPolicyAssignment                               *rolemanagementpolicyassignment.RoleManagementPolicyAssignmentClient
	RoleManagementPolicyAssignmentPolicy                         *rolemanagementpolicyassignmentpolicy.RoleManagementPolicyAssignmentPolicyClient
	RoleManagementPolicyEffectiveRule                            *rolemanagementpolicyeffectiverule.RoleManagementPolicyEffectiveRuleClient
	RoleManagementPolicyRule                                     *rolemanagementpolicyrule.RoleManagementPolicyRuleClient
	TokenIssuancePolicy                                          *tokenissuancepolicy.TokenIssuancePolicyClient
	TokenIssuancePolicyAppliesTo                                 *tokenissuancepolicyappliesto.TokenIssuancePolicyAppliesToClient
	TokenLifetimePolicy                                          *tokenlifetimepolicy.TokenLifetimePolicyClient
	TokenLifetimePolicyAppliesTo                                 *tokenlifetimepolicyappliesto.TokenLifetimePolicyAppliesToClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
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

	defaultAppManagementPolicyClient, err := defaultappmanagementpolicy.NewDefaultAppManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefaultAppManagementPolicy client: %+v", err)
	}
	configureFunc(defaultAppManagementPolicyClient.Client)

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
		ClaimsMappingPolicy:                                          claimsMappingPolicyClient,
		ClaimsMappingPolicyAppliesTo:                                 claimsMappingPolicyAppliesToClient,
		ConditionalAccessPolicy:                                      conditionalAccessPolicyClient,
		CrossTenantAccessPolicy:                                      crossTenantAccessPolicyClient,
		CrossTenantAccessPolicyDefault:                               crossTenantAccessPolicyDefaultClient,
		CrossTenantAccessPolicyPartner:                               crossTenantAccessPolicyPartnerClient,
		CrossTenantAccessPolicyPartnerIdentitySynchronization:        crossTenantAccessPolicyPartnerIdentitySynchronizationClient,
		DefaultAppManagementPolicy:                                   defaultAppManagementPolicyClient,
		FeatureRolloutPolicy:                                         featureRolloutPolicyClient,
		FeatureRolloutPolicyAppliesTo:                                featureRolloutPolicyAppliesToClient,
		HomeRealmDiscoveryPolicy:                                     homeRealmDiscoveryPolicyClient,
		HomeRealmDiscoveryPolicyAppliesTo:                            homeRealmDiscoveryPolicyAppliesToClient,
		IdentitySecurityDefaultsEnforcementPolicy:                    identitySecurityDefaultsEnforcementPolicyClient,
		PermissionGrantPolicy:                                        permissionGrantPolicyClient,
		PermissionGrantPolicyExclude:                                 permissionGrantPolicyExcludeClient,
		PermissionGrantPolicyInclude:                                 permissionGrantPolicyIncludeClient,
		Policy:                                                       policyClient,
		RoleManagementPolicy:                                         roleManagementPolicyClient,
		RoleManagementPolicyAssignment:                               roleManagementPolicyAssignmentClient,
		RoleManagementPolicyAssignmentPolicy:                         roleManagementPolicyAssignmentPolicyClient,
		RoleManagementPolicyEffectiveRule:                            roleManagementPolicyEffectiveRuleClient,
		RoleManagementPolicyRule:                                     roleManagementPolicyRuleClient,
		TokenIssuancePolicy:                                          tokenIssuancePolicyClient,
		TokenIssuancePolicyAppliesTo:                                 tokenIssuancePolicyAppliesToClient,
		TokenLifetimePolicy:                                          tokenLifetimePolicyClient,
		TokenLifetimePolicyAppliesTo:                                 tokenLifetimePolicyAppliesToClient,
	}, nil
}
