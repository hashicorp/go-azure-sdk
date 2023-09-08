package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationBaseRecommendationType string

const (
	RecommendationBaseRecommendationTypeadfsAppsMigration           RecommendationBaseRecommendationType = "AdfsAppsMigration"
	RecommendationBaseRecommendationTypeadminMFAV2                  RecommendationBaseRecommendationType = "AdminMFAV2"
	RecommendationBaseRecommendationTypeappRoleAssignmentsGroups    RecommendationBaseRecommendationType = "AppRoleAssignmentsGroups"
	RecommendationBaseRecommendationTypeappRoleAssignmentsUsers     RecommendationBaseRecommendationType = "AppRoleAssignmentsUsers"
	RecommendationBaseRecommendationTypeapplicationCredentialExpiry RecommendationBaseRecommendationType = "ApplicationCredentialExpiry"
	RecommendationBaseRecommendationTypeblockLegacyAuthentication   RecommendationBaseRecommendationType = "BlockLegacyAuthentication"
	RecommendationBaseRecommendationTypeenableDesktopSSO            RecommendationBaseRecommendationType = "EnableDesktopSSO"
	RecommendationBaseRecommendationTypeenablePHS                   RecommendationBaseRecommendationType = "EnablePHS"
	RecommendationBaseRecommendationTypeenableProvisioning          RecommendationBaseRecommendationType = "EnableProvisioning"
	RecommendationBaseRecommendationTypeintegratedApps              RecommendationBaseRecommendationType = "IntegratedApps"
	RecommendationBaseRecommendationTypemanagedIdentity             RecommendationBaseRecommendationType = "ManagedIdentity"
	RecommendationBaseRecommendationTypemfaRegistrationV2           RecommendationBaseRecommendationType = "MfaRegistrationV2"
	RecommendationBaseRecommendationTypeoneAdmin                    RecommendationBaseRecommendationType = "OneAdmin"
	RecommendationBaseRecommendationTypeoverprivilegedApps          RecommendationBaseRecommendationType = "OverprivilegedApps"
	RecommendationBaseRecommendationTypepasswordHashSync            RecommendationBaseRecommendationType = "PasswordHashSync"
	RecommendationBaseRecommendationTypeprivateLinkForAAD           RecommendationBaseRecommendationType = "PrivateLinkForAAD"
	RecommendationBaseRecommendationTypepwagePolicyNew              RecommendationBaseRecommendationType = "PwagePolicyNew"
	RecommendationBaseRecommendationTyperoleOverlap                 RecommendationBaseRecommendationType = "RoleOverlap"
	RecommendationBaseRecommendationTypeselfServicePasswordReset    RecommendationBaseRecommendationType = "SelfServicePasswordReset"
	RecommendationBaseRecommendationTypeservicePrincipalKeyExpiry   RecommendationBaseRecommendationType = "ServicePrincipalKeyExpiry"
	RecommendationBaseRecommendationTypesigninRiskPolicy            RecommendationBaseRecommendationType = "SigninRiskPolicy"
	RecommendationBaseRecommendationTypestaleAppCreds               RecommendationBaseRecommendationType = "StaleAppCreds"
	RecommendationBaseRecommendationTypestaleApps                   RecommendationBaseRecommendationType = "StaleApps"
	RecommendationBaseRecommendationTypeswitchFromPerUserMFA        RecommendationBaseRecommendationType = "SwitchFromPerUserMFA"
	RecommendationBaseRecommendationTypetenantMFA                   RecommendationBaseRecommendationType = "TenantMFA"
	RecommendationBaseRecommendationTypethirdPartyApps              RecommendationBaseRecommendationType = "ThirdPartyApps"
	RecommendationBaseRecommendationTypeturnOffPerUserMFA           RecommendationBaseRecommendationType = "TurnOffPerUserMFA"
	RecommendationBaseRecommendationTypeuseAuthenticatorApp         RecommendationBaseRecommendationType = "UseAuthenticatorApp"
	RecommendationBaseRecommendationTypeuseMyApps                   RecommendationBaseRecommendationType = "UseMyApps"
	RecommendationBaseRecommendationTypeuserRiskPolicy              RecommendationBaseRecommendationType = "UserRiskPolicy"
	RecommendationBaseRecommendationTypeverifyAppPublisher          RecommendationBaseRecommendationType = "VerifyAppPublisher"
)

func PossibleValuesForRecommendationBaseRecommendationType() []string {
	return []string{
		string(RecommendationBaseRecommendationTypeadfsAppsMigration),
		string(RecommendationBaseRecommendationTypeadminMFAV2),
		string(RecommendationBaseRecommendationTypeappRoleAssignmentsGroups),
		string(RecommendationBaseRecommendationTypeappRoleAssignmentsUsers),
		string(RecommendationBaseRecommendationTypeapplicationCredentialExpiry),
		string(RecommendationBaseRecommendationTypeblockLegacyAuthentication),
		string(RecommendationBaseRecommendationTypeenableDesktopSSO),
		string(RecommendationBaseRecommendationTypeenablePHS),
		string(RecommendationBaseRecommendationTypeenableProvisioning),
		string(RecommendationBaseRecommendationTypeintegratedApps),
		string(RecommendationBaseRecommendationTypemanagedIdentity),
		string(RecommendationBaseRecommendationTypemfaRegistrationV2),
		string(RecommendationBaseRecommendationTypeoneAdmin),
		string(RecommendationBaseRecommendationTypeoverprivilegedApps),
		string(RecommendationBaseRecommendationTypepasswordHashSync),
		string(RecommendationBaseRecommendationTypeprivateLinkForAAD),
		string(RecommendationBaseRecommendationTypepwagePolicyNew),
		string(RecommendationBaseRecommendationTyperoleOverlap),
		string(RecommendationBaseRecommendationTypeselfServicePasswordReset),
		string(RecommendationBaseRecommendationTypeservicePrincipalKeyExpiry),
		string(RecommendationBaseRecommendationTypesigninRiskPolicy),
		string(RecommendationBaseRecommendationTypestaleAppCreds),
		string(RecommendationBaseRecommendationTypestaleApps),
		string(RecommendationBaseRecommendationTypeswitchFromPerUserMFA),
		string(RecommendationBaseRecommendationTypetenantMFA),
		string(RecommendationBaseRecommendationTypethirdPartyApps),
		string(RecommendationBaseRecommendationTypeturnOffPerUserMFA),
		string(RecommendationBaseRecommendationTypeuseAuthenticatorApp),
		string(RecommendationBaseRecommendationTypeuseMyApps),
		string(RecommendationBaseRecommendationTypeuserRiskPolicy),
		string(RecommendationBaseRecommendationTypeverifyAppPublisher),
	}
}

func parseRecommendationBaseRecommendationType(input string) (*RecommendationBaseRecommendationType, error) {
	vals := map[string]RecommendationBaseRecommendationType{
		"adfsappsmigration":           RecommendationBaseRecommendationTypeadfsAppsMigration,
		"adminmfav2":                  RecommendationBaseRecommendationTypeadminMFAV2,
		"approleassignmentsgroups":    RecommendationBaseRecommendationTypeappRoleAssignmentsGroups,
		"approleassignmentsusers":     RecommendationBaseRecommendationTypeappRoleAssignmentsUsers,
		"applicationcredentialexpiry": RecommendationBaseRecommendationTypeapplicationCredentialExpiry,
		"blocklegacyauthentication":   RecommendationBaseRecommendationTypeblockLegacyAuthentication,
		"enabledesktopsso":            RecommendationBaseRecommendationTypeenableDesktopSSO,
		"enablephs":                   RecommendationBaseRecommendationTypeenablePHS,
		"enableprovisioning":          RecommendationBaseRecommendationTypeenableProvisioning,
		"integratedapps":              RecommendationBaseRecommendationTypeintegratedApps,
		"managedidentity":             RecommendationBaseRecommendationTypemanagedIdentity,
		"mfaregistrationv2":           RecommendationBaseRecommendationTypemfaRegistrationV2,
		"oneadmin":                    RecommendationBaseRecommendationTypeoneAdmin,
		"overprivilegedapps":          RecommendationBaseRecommendationTypeoverprivilegedApps,
		"passwordhashsync":            RecommendationBaseRecommendationTypepasswordHashSync,
		"privatelinkforaad":           RecommendationBaseRecommendationTypeprivateLinkForAAD,
		"pwagepolicynew":              RecommendationBaseRecommendationTypepwagePolicyNew,
		"roleoverlap":                 RecommendationBaseRecommendationTyperoleOverlap,
		"selfservicepasswordreset":    RecommendationBaseRecommendationTypeselfServicePasswordReset,
		"serviceprincipalkeyexpiry":   RecommendationBaseRecommendationTypeservicePrincipalKeyExpiry,
		"signinriskpolicy":            RecommendationBaseRecommendationTypesigninRiskPolicy,
		"staleappcreds":               RecommendationBaseRecommendationTypestaleAppCreds,
		"staleapps":                   RecommendationBaseRecommendationTypestaleApps,
		"switchfromperusermfa":        RecommendationBaseRecommendationTypeswitchFromPerUserMFA,
		"tenantmfa":                   RecommendationBaseRecommendationTypetenantMFA,
		"thirdpartyapps":              RecommendationBaseRecommendationTypethirdPartyApps,
		"turnoffperusermfa":           RecommendationBaseRecommendationTypeturnOffPerUserMFA,
		"useauthenticatorapp":         RecommendationBaseRecommendationTypeuseAuthenticatorApp,
		"usemyapps":                   RecommendationBaseRecommendationTypeuseMyApps,
		"userriskpolicy":              RecommendationBaseRecommendationTypeuserRiskPolicy,
		"verifyapppublisher":          RecommendationBaseRecommendationTypeverifyAppPublisher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationBaseRecommendationType(input)
	return &out, nil
}
