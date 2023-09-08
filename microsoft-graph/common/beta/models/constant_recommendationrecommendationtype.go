package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationRecommendationType string

const (
	RecommendationRecommendationTypeadfsAppsMigration           RecommendationRecommendationType = "AdfsAppsMigration"
	RecommendationRecommendationTypeadminMFAV2                  RecommendationRecommendationType = "AdminMFAV2"
	RecommendationRecommendationTypeappRoleAssignmentsGroups    RecommendationRecommendationType = "AppRoleAssignmentsGroups"
	RecommendationRecommendationTypeappRoleAssignmentsUsers     RecommendationRecommendationType = "AppRoleAssignmentsUsers"
	RecommendationRecommendationTypeapplicationCredentialExpiry RecommendationRecommendationType = "ApplicationCredentialExpiry"
	RecommendationRecommendationTypeblockLegacyAuthentication   RecommendationRecommendationType = "BlockLegacyAuthentication"
	RecommendationRecommendationTypeenableDesktopSSO            RecommendationRecommendationType = "EnableDesktopSSO"
	RecommendationRecommendationTypeenablePHS                   RecommendationRecommendationType = "EnablePHS"
	RecommendationRecommendationTypeenableProvisioning          RecommendationRecommendationType = "EnableProvisioning"
	RecommendationRecommendationTypeintegratedApps              RecommendationRecommendationType = "IntegratedApps"
	RecommendationRecommendationTypemanagedIdentity             RecommendationRecommendationType = "ManagedIdentity"
	RecommendationRecommendationTypemfaRegistrationV2           RecommendationRecommendationType = "MfaRegistrationV2"
	RecommendationRecommendationTypeoneAdmin                    RecommendationRecommendationType = "OneAdmin"
	RecommendationRecommendationTypeoverprivilegedApps          RecommendationRecommendationType = "OverprivilegedApps"
	RecommendationRecommendationTypepasswordHashSync            RecommendationRecommendationType = "PasswordHashSync"
	RecommendationRecommendationTypeprivateLinkForAAD           RecommendationRecommendationType = "PrivateLinkForAAD"
	RecommendationRecommendationTypepwagePolicyNew              RecommendationRecommendationType = "PwagePolicyNew"
	RecommendationRecommendationTyperoleOverlap                 RecommendationRecommendationType = "RoleOverlap"
	RecommendationRecommendationTypeselfServicePasswordReset    RecommendationRecommendationType = "SelfServicePasswordReset"
	RecommendationRecommendationTypeservicePrincipalKeyExpiry   RecommendationRecommendationType = "ServicePrincipalKeyExpiry"
	RecommendationRecommendationTypesigninRiskPolicy            RecommendationRecommendationType = "SigninRiskPolicy"
	RecommendationRecommendationTypestaleAppCreds               RecommendationRecommendationType = "StaleAppCreds"
	RecommendationRecommendationTypestaleApps                   RecommendationRecommendationType = "StaleApps"
	RecommendationRecommendationTypeswitchFromPerUserMFA        RecommendationRecommendationType = "SwitchFromPerUserMFA"
	RecommendationRecommendationTypetenantMFA                   RecommendationRecommendationType = "TenantMFA"
	RecommendationRecommendationTypethirdPartyApps              RecommendationRecommendationType = "ThirdPartyApps"
	RecommendationRecommendationTypeturnOffPerUserMFA           RecommendationRecommendationType = "TurnOffPerUserMFA"
	RecommendationRecommendationTypeuseAuthenticatorApp         RecommendationRecommendationType = "UseAuthenticatorApp"
	RecommendationRecommendationTypeuseMyApps                   RecommendationRecommendationType = "UseMyApps"
	RecommendationRecommendationTypeuserRiskPolicy              RecommendationRecommendationType = "UserRiskPolicy"
	RecommendationRecommendationTypeverifyAppPublisher          RecommendationRecommendationType = "VerifyAppPublisher"
)

func PossibleValuesForRecommendationRecommendationType() []string {
	return []string{
		string(RecommendationRecommendationTypeadfsAppsMigration),
		string(RecommendationRecommendationTypeadminMFAV2),
		string(RecommendationRecommendationTypeappRoleAssignmentsGroups),
		string(RecommendationRecommendationTypeappRoleAssignmentsUsers),
		string(RecommendationRecommendationTypeapplicationCredentialExpiry),
		string(RecommendationRecommendationTypeblockLegacyAuthentication),
		string(RecommendationRecommendationTypeenableDesktopSSO),
		string(RecommendationRecommendationTypeenablePHS),
		string(RecommendationRecommendationTypeenableProvisioning),
		string(RecommendationRecommendationTypeintegratedApps),
		string(RecommendationRecommendationTypemanagedIdentity),
		string(RecommendationRecommendationTypemfaRegistrationV2),
		string(RecommendationRecommendationTypeoneAdmin),
		string(RecommendationRecommendationTypeoverprivilegedApps),
		string(RecommendationRecommendationTypepasswordHashSync),
		string(RecommendationRecommendationTypeprivateLinkForAAD),
		string(RecommendationRecommendationTypepwagePolicyNew),
		string(RecommendationRecommendationTyperoleOverlap),
		string(RecommendationRecommendationTypeselfServicePasswordReset),
		string(RecommendationRecommendationTypeservicePrincipalKeyExpiry),
		string(RecommendationRecommendationTypesigninRiskPolicy),
		string(RecommendationRecommendationTypestaleAppCreds),
		string(RecommendationRecommendationTypestaleApps),
		string(RecommendationRecommendationTypeswitchFromPerUserMFA),
		string(RecommendationRecommendationTypetenantMFA),
		string(RecommendationRecommendationTypethirdPartyApps),
		string(RecommendationRecommendationTypeturnOffPerUserMFA),
		string(RecommendationRecommendationTypeuseAuthenticatorApp),
		string(RecommendationRecommendationTypeuseMyApps),
		string(RecommendationRecommendationTypeuserRiskPolicy),
		string(RecommendationRecommendationTypeverifyAppPublisher),
	}
}

func parseRecommendationRecommendationType(input string) (*RecommendationRecommendationType, error) {
	vals := map[string]RecommendationRecommendationType{
		"adfsappsmigration":           RecommendationRecommendationTypeadfsAppsMigration,
		"adminmfav2":                  RecommendationRecommendationTypeadminMFAV2,
		"approleassignmentsgroups":    RecommendationRecommendationTypeappRoleAssignmentsGroups,
		"approleassignmentsusers":     RecommendationRecommendationTypeappRoleAssignmentsUsers,
		"applicationcredentialexpiry": RecommendationRecommendationTypeapplicationCredentialExpiry,
		"blocklegacyauthentication":   RecommendationRecommendationTypeblockLegacyAuthentication,
		"enabledesktopsso":            RecommendationRecommendationTypeenableDesktopSSO,
		"enablephs":                   RecommendationRecommendationTypeenablePHS,
		"enableprovisioning":          RecommendationRecommendationTypeenableProvisioning,
		"integratedapps":              RecommendationRecommendationTypeintegratedApps,
		"managedidentity":             RecommendationRecommendationTypemanagedIdentity,
		"mfaregistrationv2":           RecommendationRecommendationTypemfaRegistrationV2,
		"oneadmin":                    RecommendationRecommendationTypeoneAdmin,
		"overprivilegedapps":          RecommendationRecommendationTypeoverprivilegedApps,
		"passwordhashsync":            RecommendationRecommendationTypepasswordHashSync,
		"privatelinkforaad":           RecommendationRecommendationTypeprivateLinkForAAD,
		"pwagepolicynew":              RecommendationRecommendationTypepwagePolicyNew,
		"roleoverlap":                 RecommendationRecommendationTyperoleOverlap,
		"selfservicepasswordreset":    RecommendationRecommendationTypeselfServicePasswordReset,
		"serviceprincipalkeyexpiry":   RecommendationRecommendationTypeservicePrincipalKeyExpiry,
		"signinriskpolicy":            RecommendationRecommendationTypesigninRiskPolicy,
		"staleappcreds":               RecommendationRecommendationTypestaleAppCreds,
		"staleapps":                   RecommendationRecommendationTypestaleApps,
		"switchfromperusermfa":        RecommendationRecommendationTypeswitchFromPerUserMFA,
		"tenantmfa":                   RecommendationRecommendationTypetenantMFA,
		"thirdpartyapps":              RecommendationRecommendationTypethirdPartyApps,
		"turnoffperusermfa":           RecommendationRecommendationTypeturnOffPerUserMFA,
		"useauthenticatorapp":         RecommendationRecommendationTypeuseAuthenticatorApp,
		"usemyapps":                   RecommendationRecommendationTypeuseMyApps,
		"userriskpolicy":              RecommendationRecommendationTypeuserRiskPolicy,
		"verifyapppublisher":          RecommendationRecommendationTypeverifyAppPublisher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationRecommendationType(input)
	return &out, nil
}
