package v1_0

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyactivitybasedtimeoutpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyactivitybasedtimeoutpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyadminconsentrequestpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyappmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyappmanagementpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyauthenticationflowspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyauthenticationmethodspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyauthenticationmethodspolicyauthenticationmethodconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyauthenticationstrengthpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyauthenticationstrengthpolicycombinationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyauthorizationpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyclaimsmappingpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyclaimsmappingpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyconditionalaccesspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policycrosstenantaccesspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policycrosstenantaccesspolicydefault"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policycrosstenantaccesspolicypartner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policycrosstenantaccesspolicypartneridentitysynchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policydefaultappmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyfeaturerolloutpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyfeaturerolloutpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyhomerealmdiscoverypolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyhomerealmdiscoverypolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyidentitysecuritydefaultsenforcementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policypermissiongrantpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policypermissiongrantpolicyexclude"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policypermissiongrantpolicyinclude"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyrolemanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyrolemanagementpolicyassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyrolemanagementpolicyassignmentpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyrolemanagementpolicyeffectiverule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policyrolemanagementpolicyrule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policytokenissuancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policytokenissuancepolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policytokenlifetimepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/v1_0/policytokenlifetimepolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	Policy                                                             *policy.PolicyClient
	PolicyActivityBasedTimeoutPolicy                                   *policyactivitybasedtimeoutpolicy.PolicyActivityBasedTimeoutPolicyClient
	PolicyActivityBasedTimeoutPolicyAppliesTo                          *policyactivitybasedtimeoutpolicyappliesto.PolicyActivityBasedTimeoutPolicyAppliesToClient
	PolicyAdminConsentRequestPolicy                                    *policyadminconsentrequestpolicy.PolicyAdminConsentRequestPolicyClient
	PolicyAppManagementPolicy                                          *policyappmanagementpolicy.PolicyAppManagementPolicyClient
	PolicyAppManagementPolicyAppliesTo                                 *policyappmanagementpolicyappliesto.PolicyAppManagementPolicyAppliesToClient
	PolicyAuthenticationFlowsPolicy                                    *policyauthenticationflowspolicy.PolicyAuthenticationFlowsPolicyClient
	PolicyAuthenticationMethodsPolicy                                  *policyauthenticationmethodspolicy.PolicyAuthenticationMethodsPolicyClient
	PolicyAuthenticationMethodsPolicyAuthenticationMethodConfiguration *policyauthenticationmethodspolicyauthenticationmethodconfiguration.PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationClient
	PolicyAuthenticationStrengthPolicy                                 *policyauthenticationstrengthpolicy.PolicyAuthenticationStrengthPolicyClient
	PolicyAuthenticationStrengthPolicyCombinationConfiguration         *policyauthenticationstrengthpolicycombinationconfiguration.PolicyAuthenticationStrengthPolicyCombinationConfigurationClient
	PolicyAuthorizationPolicy                                          *policyauthorizationpolicy.PolicyAuthorizationPolicyClient
	PolicyClaimsMappingPolicy                                          *policyclaimsmappingpolicy.PolicyClaimsMappingPolicyClient
	PolicyClaimsMappingPolicyAppliesTo                                 *policyclaimsmappingpolicyappliesto.PolicyClaimsMappingPolicyAppliesToClient
	PolicyConditionalAccessPolicy                                      *policyconditionalaccesspolicy.PolicyConditionalAccessPolicyClient
	PolicyCrossTenantAccessPolicy                                      *policycrosstenantaccesspolicy.PolicyCrossTenantAccessPolicyClient
	PolicyCrossTenantAccessPolicyDefault                               *policycrosstenantaccesspolicydefault.PolicyCrossTenantAccessPolicyDefaultClient
	PolicyCrossTenantAccessPolicyPartner                               *policycrosstenantaccesspolicypartner.PolicyCrossTenantAccessPolicyPartnerClient
	PolicyCrossTenantAccessPolicyPartnerIdentitySynchronization        *policycrosstenantaccesspolicypartneridentitysynchronization.PolicyCrossTenantAccessPolicyPartnerIdentitySynchronizationClient
	PolicyDefaultAppManagementPolicy                                   *policydefaultappmanagementpolicy.PolicyDefaultAppManagementPolicyClient
	PolicyFeatureRolloutPolicy                                         *policyfeaturerolloutpolicy.PolicyFeatureRolloutPolicyClient
	PolicyFeatureRolloutPolicyAppliesTo                                *policyfeaturerolloutpolicyappliesto.PolicyFeatureRolloutPolicyAppliesToClient
	PolicyHomeRealmDiscoveryPolicy                                     *policyhomerealmdiscoverypolicy.PolicyHomeRealmDiscoveryPolicyClient
	PolicyHomeRealmDiscoveryPolicyAppliesTo                            *policyhomerealmdiscoverypolicyappliesto.PolicyHomeRealmDiscoveryPolicyAppliesToClient
	PolicyIdentitySecurityDefaultsEnforcementPolicy                    *policyidentitysecuritydefaultsenforcementpolicy.PolicyIdentitySecurityDefaultsEnforcementPolicyClient
	PolicyPermissionGrantPolicy                                        *policypermissiongrantpolicy.PolicyPermissionGrantPolicyClient
	PolicyPermissionGrantPolicyExclude                                 *policypermissiongrantpolicyexclude.PolicyPermissionGrantPolicyExcludeClient
	PolicyPermissionGrantPolicyInclude                                 *policypermissiongrantpolicyinclude.PolicyPermissionGrantPolicyIncludeClient
	PolicyRoleManagementPolicy                                         *policyrolemanagementpolicy.PolicyRoleManagementPolicyClient
	PolicyRoleManagementPolicyAssignment                               *policyrolemanagementpolicyassignment.PolicyRoleManagementPolicyAssignmentClient
	PolicyRoleManagementPolicyAssignmentPolicy                         *policyrolemanagementpolicyassignmentpolicy.PolicyRoleManagementPolicyAssignmentPolicyClient
	PolicyRoleManagementPolicyEffectiveRule                            *policyrolemanagementpolicyeffectiverule.PolicyRoleManagementPolicyEffectiveRuleClient
	PolicyRoleManagementPolicyRule                                     *policyrolemanagementpolicyrule.PolicyRoleManagementPolicyRuleClient
	PolicyTokenIssuancePolicy                                          *policytokenissuancepolicy.PolicyTokenIssuancePolicyClient
	PolicyTokenIssuancePolicyAppliesTo                                 *policytokenissuancepolicyappliesto.PolicyTokenIssuancePolicyAppliesToClient
	PolicyTokenLifetimePolicy                                          *policytokenlifetimepolicy.PolicyTokenLifetimePolicyClient
	PolicyTokenLifetimePolicyAppliesTo                                 *policytokenlifetimepolicyappliesto.PolicyTokenLifetimePolicyAppliesToClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	policyActivityBasedTimeoutPolicyAppliesToClient, err := policyactivitybasedtimeoutpolicyappliesto.NewPolicyActivityBasedTimeoutPolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyActivityBasedTimeoutPolicyAppliesTo client: %+v", err)
	}
	configureFunc(policyActivityBasedTimeoutPolicyAppliesToClient.Client)

	policyActivityBasedTimeoutPolicyClient, err := policyactivitybasedtimeoutpolicy.NewPolicyActivityBasedTimeoutPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyActivityBasedTimeoutPolicy client: %+v", err)
	}
	configureFunc(policyActivityBasedTimeoutPolicyClient.Client)

	policyAdminConsentRequestPolicyClient, err := policyadminconsentrequestpolicy.NewPolicyAdminConsentRequestPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAdminConsentRequestPolicy client: %+v", err)
	}
	configureFunc(policyAdminConsentRequestPolicyClient.Client)

	policyAppManagementPolicyAppliesToClient, err := policyappmanagementpolicyappliesto.NewPolicyAppManagementPolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAppManagementPolicyAppliesTo client: %+v", err)
	}
	configureFunc(policyAppManagementPolicyAppliesToClient.Client)

	policyAppManagementPolicyClient, err := policyappmanagementpolicy.NewPolicyAppManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAppManagementPolicy client: %+v", err)
	}
	configureFunc(policyAppManagementPolicyClient.Client)

	policyAuthenticationFlowsPolicyClient, err := policyauthenticationflowspolicy.NewPolicyAuthenticationFlowsPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAuthenticationFlowsPolicy client: %+v", err)
	}
	configureFunc(policyAuthenticationFlowsPolicyClient.Client)

	policyAuthenticationMethodsPolicyAuthenticationMethodConfigurationClient, err := policyauthenticationmethodspolicyauthenticationmethodconfiguration.NewPolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAuthenticationMethodsPolicyAuthenticationMethodConfiguration client: %+v", err)
	}
	configureFunc(policyAuthenticationMethodsPolicyAuthenticationMethodConfigurationClient.Client)

	policyAuthenticationMethodsPolicyClient, err := policyauthenticationmethodspolicy.NewPolicyAuthenticationMethodsPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAuthenticationMethodsPolicy client: %+v", err)
	}
	configureFunc(policyAuthenticationMethodsPolicyClient.Client)

	policyAuthenticationStrengthPolicyClient, err := policyauthenticationstrengthpolicy.NewPolicyAuthenticationStrengthPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAuthenticationStrengthPolicy client: %+v", err)
	}
	configureFunc(policyAuthenticationStrengthPolicyClient.Client)

	policyAuthenticationStrengthPolicyCombinationConfigurationClient, err := policyauthenticationstrengthpolicycombinationconfiguration.NewPolicyAuthenticationStrengthPolicyCombinationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAuthenticationStrengthPolicyCombinationConfiguration client: %+v", err)
	}
	configureFunc(policyAuthenticationStrengthPolicyCombinationConfigurationClient.Client)

	policyAuthorizationPolicyClient, err := policyauthorizationpolicy.NewPolicyAuthorizationPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAuthorizationPolicy client: %+v", err)
	}
	configureFunc(policyAuthorizationPolicyClient.Client)

	policyClaimsMappingPolicyAppliesToClient, err := policyclaimsmappingpolicyappliesto.NewPolicyClaimsMappingPolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyClaimsMappingPolicyAppliesTo client: %+v", err)
	}
	configureFunc(policyClaimsMappingPolicyAppliesToClient.Client)

	policyClaimsMappingPolicyClient, err := policyclaimsmappingpolicy.NewPolicyClaimsMappingPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyClaimsMappingPolicy client: %+v", err)
	}
	configureFunc(policyClaimsMappingPolicyClient.Client)

	policyClient, err := policy.NewPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Policy client: %+v", err)
	}
	configureFunc(policyClient.Client)

	policyConditionalAccessPolicyClient, err := policyconditionalaccesspolicy.NewPolicyConditionalAccessPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyConditionalAccessPolicy client: %+v", err)
	}
	configureFunc(policyConditionalAccessPolicyClient.Client)

	policyCrossTenantAccessPolicyClient, err := policycrosstenantaccesspolicy.NewPolicyCrossTenantAccessPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyCrossTenantAccessPolicy client: %+v", err)
	}
	configureFunc(policyCrossTenantAccessPolicyClient.Client)

	policyCrossTenantAccessPolicyDefaultClient, err := policycrosstenantaccesspolicydefault.NewPolicyCrossTenantAccessPolicyDefaultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyCrossTenantAccessPolicyDefault client: %+v", err)
	}
	configureFunc(policyCrossTenantAccessPolicyDefaultClient.Client)

	policyCrossTenantAccessPolicyPartnerClient, err := policycrosstenantaccesspolicypartner.NewPolicyCrossTenantAccessPolicyPartnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyCrossTenantAccessPolicyPartner client: %+v", err)
	}
	configureFunc(policyCrossTenantAccessPolicyPartnerClient.Client)

	policyCrossTenantAccessPolicyPartnerIdentitySynchronizationClient, err := policycrosstenantaccesspolicypartneridentitysynchronization.NewPolicyCrossTenantAccessPolicyPartnerIdentitySynchronizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyCrossTenantAccessPolicyPartnerIdentitySynchronization client: %+v", err)
	}
	configureFunc(policyCrossTenantAccessPolicyPartnerIdentitySynchronizationClient.Client)

	policyDefaultAppManagementPolicyClient, err := policydefaultappmanagementpolicy.NewPolicyDefaultAppManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyDefaultAppManagementPolicy client: %+v", err)
	}
	configureFunc(policyDefaultAppManagementPolicyClient.Client)

	policyFeatureRolloutPolicyAppliesToClient, err := policyfeaturerolloutpolicyappliesto.NewPolicyFeatureRolloutPolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyFeatureRolloutPolicyAppliesTo client: %+v", err)
	}
	configureFunc(policyFeatureRolloutPolicyAppliesToClient.Client)

	policyFeatureRolloutPolicyClient, err := policyfeaturerolloutpolicy.NewPolicyFeatureRolloutPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyFeatureRolloutPolicy client: %+v", err)
	}
	configureFunc(policyFeatureRolloutPolicyClient.Client)

	policyHomeRealmDiscoveryPolicyAppliesToClient, err := policyhomerealmdiscoverypolicyappliesto.NewPolicyHomeRealmDiscoveryPolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyHomeRealmDiscoveryPolicyAppliesTo client: %+v", err)
	}
	configureFunc(policyHomeRealmDiscoveryPolicyAppliesToClient.Client)

	policyHomeRealmDiscoveryPolicyClient, err := policyhomerealmdiscoverypolicy.NewPolicyHomeRealmDiscoveryPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyHomeRealmDiscoveryPolicy client: %+v", err)
	}
	configureFunc(policyHomeRealmDiscoveryPolicyClient.Client)

	policyIdentitySecurityDefaultsEnforcementPolicyClient, err := policyidentitysecuritydefaultsenforcementpolicy.NewPolicyIdentitySecurityDefaultsEnforcementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyIdentitySecurityDefaultsEnforcementPolicy client: %+v", err)
	}
	configureFunc(policyIdentitySecurityDefaultsEnforcementPolicyClient.Client)

	policyPermissionGrantPolicyClient, err := policypermissiongrantpolicy.NewPolicyPermissionGrantPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyPermissionGrantPolicy client: %+v", err)
	}
	configureFunc(policyPermissionGrantPolicyClient.Client)

	policyPermissionGrantPolicyExcludeClient, err := policypermissiongrantpolicyexclude.NewPolicyPermissionGrantPolicyExcludeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyPermissionGrantPolicyExclude client: %+v", err)
	}
	configureFunc(policyPermissionGrantPolicyExcludeClient.Client)

	policyPermissionGrantPolicyIncludeClient, err := policypermissiongrantpolicyinclude.NewPolicyPermissionGrantPolicyIncludeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyPermissionGrantPolicyInclude client: %+v", err)
	}
	configureFunc(policyPermissionGrantPolicyIncludeClient.Client)

	policyRoleManagementPolicyAssignmentClient, err := policyrolemanagementpolicyassignment.NewPolicyRoleManagementPolicyAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyRoleManagementPolicyAssignment client: %+v", err)
	}
	configureFunc(policyRoleManagementPolicyAssignmentClient.Client)

	policyRoleManagementPolicyAssignmentPolicyClient, err := policyrolemanagementpolicyassignmentpolicy.NewPolicyRoleManagementPolicyAssignmentPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyRoleManagementPolicyAssignmentPolicy client: %+v", err)
	}
	configureFunc(policyRoleManagementPolicyAssignmentPolicyClient.Client)

	policyRoleManagementPolicyClient, err := policyrolemanagementpolicy.NewPolicyRoleManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyRoleManagementPolicy client: %+v", err)
	}
	configureFunc(policyRoleManagementPolicyClient.Client)

	policyRoleManagementPolicyEffectiveRuleClient, err := policyrolemanagementpolicyeffectiverule.NewPolicyRoleManagementPolicyEffectiveRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyRoleManagementPolicyEffectiveRule client: %+v", err)
	}
	configureFunc(policyRoleManagementPolicyEffectiveRuleClient.Client)

	policyRoleManagementPolicyRuleClient, err := policyrolemanagementpolicyrule.NewPolicyRoleManagementPolicyRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyRoleManagementPolicyRule client: %+v", err)
	}
	configureFunc(policyRoleManagementPolicyRuleClient.Client)

	policyTokenIssuancePolicyAppliesToClient, err := policytokenissuancepolicyappliesto.NewPolicyTokenIssuancePolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyTokenIssuancePolicyAppliesTo client: %+v", err)
	}
	configureFunc(policyTokenIssuancePolicyAppliesToClient.Client)

	policyTokenIssuancePolicyClient, err := policytokenissuancepolicy.NewPolicyTokenIssuancePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyTokenIssuancePolicy client: %+v", err)
	}
	configureFunc(policyTokenIssuancePolicyClient.Client)

	policyTokenLifetimePolicyAppliesToClient, err := policytokenlifetimepolicyappliesto.NewPolicyTokenLifetimePolicyAppliesToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyTokenLifetimePolicyAppliesTo client: %+v", err)
	}
	configureFunc(policyTokenLifetimePolicyAppliesToClient.Client)

	policyTokenLifetimePolicyClient, err := policytokenlifetimepolicy.NewPolicyTokenLifetimePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyTokenLifetimePolicy client: %+v", err)
	}
	configureFunc(policyTokenLifetimePolicyClient.Client)

	return &Client{
		Policy:                           policyClient,
		PolicyActivityBasedTimeoutPolicy: policyActivityBasedTimeoutPolicyClient,
		PolicyActivityBasedTimeoutPolicyAppliesTo:                          policyActivityBasedTimeoutPolicyAppliesToClient,
		PolicyAdminConsentRequestPolicy:                                    policyAdminConsentRequestPolicyClient,
		PolicyAppManagementPolicy:                                          policyAppManagementPolicyClient,
		PolicyAppManagementPolicyAppliesTo:                                 policyAppManagementPolicyAppliesToClient,
		PolicyAuthenticationFlowsPolicy:                                    policyAuthenticationFlowsPolicyClient,
		PolicyAuthenticationMethodsPolicy:                                  policyAuthenticationMethodsPolicyClient,
		PolicyAuthenticationMethodsPolicyAuthenticationMethodConfiguration: policyAuthenticationMethodsPolicyAuthenticationMethodConfigurationClient,
		PolicyAuthenticationStrengthPolicy:                                 policyAuthenticationStrengthPolicyClient,
		PolicyAuthenticationStrengthPolicyCombinationConfiguration:         policyAuthenticationStrengthPolicyCombinationConfigurationClient,
		PolicyAuthorizationPolicy:                                          policyAuthorizationPolicyClient,
		PolicyClaimsMappingPolicy:                                          policyClaimsMappingPolicyClient,
		PolicyClaimsMappingPolicyAppliesTo:                                 policyClaimsMappingPolicyAppliesToClient,
		PolicyConditionalAccessPolicy:                                      policyConditionalAccessPolicyClient,
		PolicyCrossTenantAccessPolicy:                                      policyCrossTenantAccessPolicyClient,
		PolicyCrossTenantAccessPolicyDefault:                               policyCrossTenantAccessPolicyDefaultClient,
		PolicyCrossTenantAccessPolicyPartner:                               policyCrossTenantAccessPolicyPartnerClient,
		PolicyCrossTenantAccessPolicyPartnerIdentitySynchronization:        policyCrossTenantAccessPolicyPartnerIdentitySynchronizationClient,
		PolicyDefaultAppManagementPolicy:                                   policyDefaultAppManagementPolicyClient,
		PolicyFeatureRolloutPolicy:                                         policyFeatureRolloutPolicyClient,
		PolicyFeatureRolloutPolicyAppliesTo:                                policyFeatureRolloutPolicyAppliesToClient,
		PolicyHomeRealmDiscoveryPolicy:                                     policyHomeRealmDiscoveryPolicyClient,
		PolicyHomeRealmDiscoveryPolicyAppliesTo:                            policyHomeRealmDiscoveryPolicyAppliesToClient,
		PolicyIdentitySecurityDefaultsEnforcementPolicy:                    policyIdentitySecurityDefaultsEnforcementPolicyClient,
		PolicyPermissionGrantPolicy:                                        policyPermissionGrantPolicyClient,
		PolicyPermissionGrantPolicyExclude:                                 policyPermissionGrantPolicyExcludeClient,
		PolicyPermissionGrantPolicyInclude:                                 policyPermissionGrantPolicyIncludeClient,
		PolicyRoleManagementPolicy:                                         policyRoleManagementPolicyClient,
		PolicyRoleManagementPolicyAssignment:                               policyRoleManagementPolicyAssignmentClient,
		PolicyRoleManagementPolicyAssignmentPolicy:                         policyRoleManagementPolicyAssignmentPolicyClient,
		PolicyRoleManagementPolicyEffectiveRule:                            policyRoleManagementPolicyEffectiveRuleClient,
		PolicyRoleManagementPolicyRule:                                     policyRoleManagementPolicyRuleClient,
		PolicyTokenIssuancePolicy:                                          policyTokenIssuancePolicyClient,
		PolicyTokenIssuancePolicyAppliesTo:                                 policyTokenIssuancePolicyAppliesToClient,
		PolicyTokenLifetimePolicy:                                          policyTokenLifetimePolicyClient,
		PolicyTokenLifetimePolicyAppliesTo:                                 policyTokenLifetimePolicyAppliesToClient,
	}, nil
}
