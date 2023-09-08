package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyaccessreviewpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyactivitybasedtimeoutpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyactivitybasedtimeoutpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyadminconsentrequestpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyappmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyappmanagementpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyauthenticationflowspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyauthenticationmethodspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyauthenticationmethodspolicyauthenticationmethodconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyauthenticationstrengthpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyauthenticationstrengthpolicycombinationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyauthorizationpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyauthorizationpolicydefaultuserroleoverride"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyb2cauthenticationmethodspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyclaimsmappingpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyclaimsmappingpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyconditionalaccesspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policycrosstenantaccesspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policycrosstenantaccesspolicydefault"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policycrosstenantaccesspolicypartner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policycrosstenantaccesspolicypartneridentitysynchronization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policydefaultappmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policydeviceregistrationpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policydirectoryroleaccessreviewpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyexternalidentitiespolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyfeaturerolloutpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyfeaturerolloutpolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyfederatedtokenvalidationpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyhomerealmdiscoverypolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyhomerealmdiscoverypolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyidentitysecuritydefaultsenforcementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policymobileappmanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policymobileappmanagementpolicyincludedgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policymobiledevicemanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policymobiledevicemanagementpolicyincludedgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policypermissiongrantpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policypermissiongrantpolicyexclude"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policypermissiongrantpolicyinclude"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyrolemanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyrolemanagementpolicyassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyrolemanagementpolicyassignmentpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyrolemanagementpolicyeffectiverule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyrolemanagementpolicyrule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyserviceprincipalcreationpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyserviceprincipalcreationpolicyexclude"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policyserviceprincipalcreationpolicyinclude"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policytokenissuancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policytokenissuancepolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policytokenlifetimepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/beta/policytokenlifetimepolicyappliesto"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	Policy                                                             *policy.PolicyClient
	PolicyAccessReviewPolicy                                           *policyaccessreviewpolicy.PolicyAccessReviewPolicyClient
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
	PolicyAuthorizationPolicyDefaultUserRoleOverride                   *policyauthorizationpolicydefaultuserroleoverride.PolicyAuthorizationPolicyDefaultUserRoleOverrideClient
	PolicyB2cAuthenticationMethodsPolicy                               *policyb2cauthenticationmethodspolicy.PolicyB2cAuthenticationMethodsPolicyClient
	PolicyClaimsMappingPolicy                                          *policyclaimsmappingpolicy.PolicyClaimsMappingPolicyClient
	PolicyClaimsMappingPolicyAppliesTo                                 *policyclaimsmappingpolicyappliesto.PolicyClaimsMappingPolicyAppliesToClient
	PolicyConditionalAccessPolicy                                      *policyconditionalaccesspolicy.PolicyConditionalAccessPolicyClient
	PolicyCrossTenantAccessPolicy                                      *policycrosstenantaccesspolicy.PolicyCrossTenantAccessPolicyClient
	PolicyCrossTenantAccessPolicyDefault                               *policycrosstenantaccesspolicydefault.PolicyCrossTenantAccessPolicyDefaultClient
	PolicyCrossTenantAccessPolicyPartner                               *policycrosstenantaccesspolicypartner.PolicyCrossTenantAccessPolicyPartnerClient
	PolicyCrossTenantAccessPolicyPartnerIdentitySynchronization        *policycrosstenantaccesspolicypartneridentitysynchronization.PolicyCrossTenantAccessPolicyPartnerIdentitySynchronizationClient
	PolicyDefaultAppManagementPolicy                                   *policydefaultappmanagementpolicy.PolicyDefaultAppManagementPolicyClient
	PolicyDeviceRegistrationPolicy                                     *policydeviceregistrationpolicy.PolicyDeviceRegistrationPolicyClient
	PolicyDirectoryRoleAccessReviewPolicy                              *policydirectoryroleaccessreviewpolicy.PolicyDirectoryRoleAccessReviewPolicyClient
	PolicyExternalIdentitiesPolicy                                     *policyexternalidentitiespolicy.PolicyExternalIdentitiesPolicyClient
	PolicyFeatureRolloutPolicy                                         *policyfeaturerolloutpolicy.PolicyFeatureRolloutPolicyClient
	PolicyFeatureRolloutPolicyAppliesTo                                *policyfeaturerolloutpolicyappliesto.PolicyFeatureRolloutPolicyAppliesToClient
	PolicyFederatedTokenValidationPolicy                               *policyfederatedtokenvalidationpolicy.PolicyFederatedTokenValidationPolicyClient
	PolicyHomeRealmDiscoveryPolicy                                     *policyhomerealmdiscoverypolicy.PolicyHomeRealmDiscoveryPolicyClient
	PolicyHomeRealmDiscoveryPolicyAppliesTo                            *policyhomerealmdiscoverypolicyappliesto.PolicyHomeRealmDiscoveryPolicyAppliesToClient
	PolicyIdentitySecurityDefaultsEnforcementPolicy                    *policyidentitysecuritydefaultsenforcementpolicy.PolicyIdentitySecurityDefaultsEnforcementPolicyClient
	PolicyMobileAppManagementPolicy                                    *policymobileappmanagementpolicy.PolicyMobileAppManagementPolicyClient
	PolicyMobileAppManagementPolicyIncludedGroup                       *policymobileappmanagementpolicyincludedgroup.PolicyMobileAppManagementPolicyIncludedGroupClient
	PolicyMobileDeviceManagementPolicy                                 *policymobiledevicemanagementpolicy.PolicyMobileDeviceManagementPolicyClient
	PolicyMobileDeviceManagementPolicyIncludedGroup                    *policymobiledevicemanagementpolicyincludedgroup.PolicyMobileDeviceManagementPolicyIncludedGroupClient
	PolicyPermissionGrantPolicy                                        *policypermissiongrantpolicy.PolicyPermissionGrantPolicyClient
	PolicyPermissionGrantPolicyExclude                                 *policypermissiongrantpolicyexclude.PolicyPermissionGrantPolicyExcludeClient
	PolicyPermissionGrantPolicyInclude                                 *policypermissiongrantpolicyinclude.PolicyPermissionGrantPolicyIncludeClient
	PolicyRoleManagementPolicy                                         *policyrolemanagementpolicy.PolicyRoleManagementPolicyClient
	PolicyRoleManagementPolicyAssignment                               *policyrolemanagementpolicyassignment.PolicyRoleManagementPolicyAssignmentClient
	PolicyRoleManagementPolicyAssignmentPolicy                         *policyrolemanagementpolicyassignmentpolicy.PolicyRoleManagementPolicyAssignmentPolicyClient
	PolicyRoleManagementPolicyEffectiveRule                            *policyrolemanagementpolicyeffectiverule.PolicyRoleManagementPolicyEffectiveRuleClient
	PolicyRoleManagementPolicyRule                                     *policyrolemanagementpolicyrule.PolicyRoleManagementPolicyRuleClient
	PolicyServicePrincipalCreationPolicy                               *policyserviceprincipalcreationpolicy.PolicyServicePrincipalCreationPolicyClient
	PolicyServicePrincipalCreationPolicyExclude                        *policyserviceprincipalcreationpolicyexclude.PolicyServicePrincipalCreationPolicyExcludeClient
	PolicyServicePrincipalCreationPolicyInclude                        *policyserviceprincipalcreationpolicyinclude.PolicyServicePrincipalCreationPolicyIncludeClient
	PolicyTokenIssuancePolicy                                          *policytokenissuancepolicy.PolicyTokenIssuancePolicyClient
	PolicyTokenIssuancePolicyAppliesTo                                 *policytokenissuancepolicyappliesto.PolicyTokenIssuancePolicyAppliesToClient
	PolicyTokenLifetimePolicy                                          *policytokenlifetimepolicy.PolicyTokenLifetimePolicyClient
	PolicyTokenLifetimePolicyAppliesTo                                 *policytokenlifetimepolicyappliesto.PolicyTokenLifetimePolicyAppliesToClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	policyAccessReviewPolicyClient, err := policyaccessreviewpolicy.NewPolicyAccessReviewPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAccessReviewPolicy client: %+v", err)
	}
	configureFunc(policyAccessReviewPolicyClient.Client)

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

	policyAuthorizationPolicyDefaultUserRoleOverrideClient, err := policyauthorizationpolicydefaultuserroleoverride.NewPolicyAuthorizationPolicyDefaultUserRoleOverrideClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAuthorizationPolicyDefaultUserRoleOverride client: %+v", err)
	}
	configureFunc(policyAuthorizationPolicyDefaultUserRoleOverrideClient.Client)

	policyB2cAuthenticationMethodsPolicyClient, err := policyb2cauthenticationmethodspolicy.NewPolicyB2cAuthenticationMethodsPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyB2cAuthenticationMethodsPolicy client: %+v", err)
	}
	configureFunc(policyB2cAuthenticationMethodsPolicyClient.Client)

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

	policyDeviceRegistrationPolicyClient, err := policydeviceregistrationpolicy.NewPolicyDeviceRegistrationPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyDeviceRegistrationPolicy client: %+v", err)
	}
	configureFunc(policyDeviceRegistrationPolicyClient.Client)

	policyDirectoryRoleAccessReviewPolicyClient, err := policydirectoryroleaccessreviewpolicy.NewPolicyDirectoryRoleAccessReviewPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyDirectoryRoleAccessReviewPolicy client: %+v", err)
	}
	configureFunc(policyDirectoryRoleAccessReviewPolicyClient.Client)

	policyExternalIdentitiesPolicyClient, err := policyexternalidentitiespolicy.NewPolicyExternalIdentitiesPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyExternalIdentitiesPolicy client: %+v", err)
	}
	configureFunc(policyExternalIdentitiesPolicyClient.Client)

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

	policyFederatedTokenValidationPolicyClient, err := policyfederatedtokenvalidationpolicy.NewPolicyFederatedTokenValidationPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyFederatedTokenValidationPolicy client: %+v", err)
	}
	configureFunc(policyFederatedTokenValidationPolicyClient.Client)

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

	policyMobileAppManagementPolicyClient, err := policymobileappmanagementpolicy.NewPolicyMobileAppManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyMobileAppManagementPolicy client: %+v", err)
	}
	configureFunc(policyMobileAppManagementPolicyClient.Client)

	policyMobileAppManagementPolicyIncludedGroupClient, err := policymobileappmanagementpolicyincludedgroup.NewPolicyMobileAppManagementPolicyIncludedGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyMobileAppManagementPolicyIncludedGroup client: %+v", err)
	}
	configureFunc(policyMobileAppManagementPolicyIncludedGroupClient.Client)

	policyMobileDeviceManagementPolicyClient, err := policymobiledevicemanagementpolicy.NewPolicyMobileDeviceManagementPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyMobileDeviceManagementPolicy client: %+v", err)
	}
	configureFunc(policyMobileDeviceManagementPolicyClient.Client)

	policyMobileDeviceManagementPolicyIncludedGroupClient, err := policymobiledevicemanagementpolicyincludedgroup.NewPolicyMobileDeviceManagementPolicyIncludedGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyMobileDeviceManagementPolicyIncludedGroup client: %+v", err)
	}
	configureFunc(policyMobileDeviceManagementPolicyIncludedGroupClient.Client)

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

	policyServicePrincipalCreationPolicyClient, err := policyserviceprincipalcreationpolicy.NewPolicyServicePrincipalCreationPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyServicePrincipalCreationPolicy client: %+v", err)
	}
	configureFunc(policyServicePrincipalCreationPolicyClient.Client)

	policyServicePrincipalCreationPolicyExcludeClient, err := policyserviceprincipalcreationpolicyexclude.NewPolicyServicePrincipalCreationPolicyExcludeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyServicePrincipalCreationPolicyExclude client: %+v", err)
	}
	configureFunc(policyServicePrincipalCreationPolicyExcludeClient.Client)

	policyServicePrincipalCreationPolicyIncludeClient, err := policyserviceprincipalcreationpolicyinclude.NewPolicyServicePrincipalCreationPolicyIncludeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyServicePrincipalCreationPolicyInclude client: %+v", err)
	}
	configureFunc(policyServicePrincipalCreationPolicyIncludeClient.Client)

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
		Policy:                                    policyClient,
		PolicyAccessReviewPolicy:                  policyAccessReviewPolicyClient,
		PolicyActivityBasedTimeoutPolicy:          policyActivityBasedTimeoutPolicyClient,
		PolicyActivityBasedTimeoutPolicyAppliesTo: policyActivityBasedTimeoutPolicyAppliesToClient,
		PolicyAdminConsentRequestPolicy:           policyAdminConsentRequestPolicyClient,
		PolicyAppManagementPolicy:                 policyAppManagementPolicyClient,
		PolicyAppManagementPolicyAppliesTo:        policyAppManagementPolicyAppliesToClient,
		PolicyAuthenticationFlowsPolicy:           policyAuthenticationFlowsPolicyClient,
		PolicyAuthenticationMethodsPolicy:         policyAuthenticationMethodsPolicyClient,
		PolicyAuthenticationMethodsPolicyAuthenticationMethodConfiguration: policyAuthenticationMethodsPolicyAuthenticationMethodConfigurationClient,
		PolicyAuthenticationStrengthPolicy:                                 policyAuthenticationStrengthPolicyClient,
		PolicyAuthenticationStrengthPolicyCombinationConfiguration:         policyAuthenticationStrengthPolicyCombinationConfigurationClient,
		PolicyAuthorizationPolicy:                                          policyAuthorizationPolicyClient,
		PolicyAuthorizationPolicyDefaultUserRoleOverride:                   policyAuthorizationPolicyDefaultUserRoleOverrideClient,
		PolicyB2cAuthenticationMethodsPolicy:                               policyB2cAuthenticationMethodsPolicyClient,
		PolicyClaimsMappingPolicy:                                          policyClaimsMappingPolicyClient,
		PolicyClaimsMappingPolicyAppliesTo:                                 policyClaimsMappingPolicyAppliesToClient,
		PolicyConditionalAccessPolicy:                                      policyConditionalAccessPolicyClient,
		PolicyCrossTenantAccessPolicy:                                      policyCrossTenantAccessPolicyClient,
		PolicyCrossTenantAccessPolicyDefault:                               policyCrossTenantAccessPolicyDefaultClient,
		PolicyCrossTenantAccessPolicyPartner:                               policyCrossTenantAccessPolicyPartnerClient,
		PolicyCrossTenantAccessPolicyPartnerIdentitySynchronization:        policyCrossTenantAccessPolicyPartnerIdentitySynchronizationClient,
		PolicyDefaultAppManagementPolicy:                                   policyDefaultAppManagementPolicyClient,
		PolicyDeviceRegistrationPolicy:                                     policyDeviceRegistrationPolicyClient,
		PolicyDirectoryRoleAccessReviewPolicy:                              policyDirectoryRoleAccessReviewPolicyClient,
		PolicyExternalIdentitiesPolicy:                                     policyExternalIdentitiesPolicyClient,
		PolicyFeatureRolloutPolicy:                                         policyFeatureRolloutPolicyClient,
		PolicyFeatureRolloutPolicyAppliesTo:                                policyFeatureRolloutPolicyAppliesToClient,
		PolicyFederatedTokenValidationPolicy:                               policyFederatedTokenValidationPolicyClient,
		PolicyHomeRealmDiscoveryPolicy:                                     policyHomeRealmDiscoveryPolicyClient,
		PolicyHomeRealmDiscoveryPolicyAppliesTo:                            policyHomeRealmDiscoveryPolicyAppliesToClient,
		PolicyIdentitySecurityDefaultsEnforcementPolicy:                    policyIdentitySecurityDefaultsEnforcementPolicyClient,
		PolicyMobileAppManagementPolicy:                                    policyMobileAppManagementPolicyClient,
		PolicyMobileAppManagementPolicyIncludedGroup:                       policyMobileAppManagementPolicyIncludedGroupClient,
		PolicyMobileDeviceManagementPolicy:                                 policyMobileDeviceManagementPolicyClient,
		PolicyMobileDeviceManagementPolicyIncludedGroup:                    policyMobileDeviceManagementPolicyIncludedGroupClient,
		PolicyPermissionGrantPolicy:                                        policyPermissionGrantPolicyClient,
		PolicyPermissionGrantPolicyExclude:                                 policyPermissionGrantPolicyExcludeClient,
		PolicyPermissionGrantPolicyInclude:                                 policyPermissionGrantPolicyIncludeClient,
		PolicyRoleManagementPolicy:                                         policyRoleManagementPolicyClient,
		PolicyRoleManagementPolicyAssignment:                               policyRoleManagementPolicyAssignmentClient,
		PolicyRoleManagementPolicyAssignmentPolicy:                         policyRoleManagementPolicyAssignmentPolicyClient,
		PolicyRoleManagementPolicyEffectiveRule:                            policyRoleManagementPolicyEffectiveRuleClient,
		PolicyRoleManagementPolicyRule:                                     policyRoleManagementPolicyRuleClient,
		PolicyServicePrincipalCreationPolicy:                               policyServicePrincipalCreationPolicyClient,
		PolicyServicePrincipalCreationPolicyExclude:                        policyServicePrincipalCreationPolicyExcludeClient,
		PolicyServicePrincipalCreationPolicyInclude:                        policyServicePrincipalCreationPolicyIncludeClient,
		PolicyTokenIssuancePolicy:                                          policyTokenIssuancePolicyClient,
		PolicyTokenIssuancePolicyAppliesTo:                                 policyTokenIssuancePolicyAppliesToClient,
		PolicyTokenLifetimePolicy:                                          policyTokenLifetimePolicyClient,
		PolicyTokenLifetimePolicyAppliesTo:                                 policyTokenLifetimePolicyAppliesToClient,
	}, nil
}
