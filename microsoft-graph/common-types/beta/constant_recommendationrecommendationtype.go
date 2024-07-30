package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationRecommendationType string

const (
	RecommendationRecommendationType_AadConnectDeprecated        RecommendationRecommendationType = "aadConnectDeprecated"
	RecommendationRecommendationType_AdalToMsalMigration         RecommendationRecommendationType = "adalToMsalMigration"
	RecommendationRecommendationType_AdfsAppsMigration           RecommendationRecommendationType = "adfsAppsMigration"
	RecommendationRecommendationType_AdminMFAV2                  RecommendationRecommendationType = "adminMFAV2"
	RecommendationRecommendationType_AppRoleAssignmentsGroups    RecommendationRecommendationType = "appRoleAssignmentsGroups"
	RecommendationRecommendationType_AppRoleAssignmentsUsers     RecommendationRecommendationType = "appRoleAssignmentsUsers"
	RecommendationRecommendationType_ApplicationCredentialExpiry RecommendationRecommendationType = "applicationCredentialExpiry"
	RecommendationRecommendationType_BlockLegacyAuthentication   RecommendationRecommendationType = "blockLegacyAuthentication"
	RecommendationRecommendationType_EnableDesktopSSO            RecommendationRecommendationType = "enableDesktopSSO"
	RecommendationRecommendationType_EnablePHS                   RecommendationRecommendationType = "enablePHS"
	RecommendationRecommendationType_EnableProvisioning          RecommendationRecommendationType = "enableProvisioning"
	RecommendationRecommendationType_InactiveGuests              RecommendationRecommendationType = "inactiveGuests"
	RecommendationRecommendationType_IntegratedApps              RecommendationRecommendationType = "integratedApps"
	RecommendationRecommendationType_LongLivedCredentials        RecommendationRecommendationType = "longLivedCredentials"
	RecommendationRecommendationType_ManagedIdentity             RecommendationRecommendationType = "managedIdentity"
	RecommendationRecommendationType_MfaRegistrationV2           RecommendationRecommendationType = "mfaRegistrationV2"
	RecommendationRecommendationType_OneAdmin                    RecommendationRecommendationType = "oneAdmin"
	RecommendationRecommendationType_OverprivilegedApps          RecommendationRecommendationType = "overprivilegedApps"
	RecommendationRecommendationType_OwnerlessApps               RecommendationRecommendationType = "ownerlessApps"
	RecommendationRecommendationType_PasswordHashSync            RecommendationRecommendationType = "passwordHashSync"
	RecommendationRecommendationType_PrivateLinkForAAD           RecommendationRecommendationType = "privateLinkForAAD"
	RecommendationRecommendationType_PwagePolicyNew              RecommendationRecommendationType = "pwagePolicyNew"
	RecommendationRecommendationType_RoleOverlap                 RecommendationRecommendationType = "roleOverlap"
	RecommendationRecommendationType_SelfServicePasswordReset    RecommendationRecommendationType = "selfServicePasswordReset"
	RecommendationRecommendationType_ServicePrincipalKeyExpiry   RecommendationRecommendationType = "servicePrincipalKeyExpiry"
	RecommendationRecommendationType_SigninRiskPolicy            RecommendationRecommendationType = "signinRiskPolicy"
	RecommendationRecommendationType_StaleAppCreds               RecommendationRecommendationType = "staleAppCreds"
	RecommendationRecommendationType_StaleApps                   RecommendationRecommendationType = "staleApps"
	RecommendationRecommendationType_SwitchFromPerUserMFA        RecommendationRecommendationType = "switchFromPerUserMFA"
	RecommendationRecommendationType_TenantMFA                   RecommendationRecommendationType = "tenantMFA"
	RecommendationRecommendationType_ThirdPartyApps              RecommendationRecommendationType = "thirdPartyApps"
	RecommendationRecommendationType_TurnOffPerUserMFA           RecommendationRecommendationType = "turnOffPerUserMFA"
	RecommendationRecommendationType_UseAuthenticatorApp         RecommendationRecommendationType = "useAuthenticatorApp"
	RecommendationRecommendationType_UseMyApps                   RecommendationRecommendationType = "useMyApps"
	RecommendationRecommendationType_UserRiskPolicy              RecommendationRecommendationType = "userRiskPolicy"
	RecommendationRecommendationType_VerifyAppPublisher          RecommendationRecommendationType = "verifyAppPublisher"
)

func PossibleValuesForRecommendationRecommendationType() []string {
	return []string{
		string(RecommendationRecommendationType_AadConnectDeprecated),
		string(RecommendationRecommendationType_AdalToMsalMigration),
		string(RecommendationRecommendationType_AdfsAppsMigration),
		string(RecommendationRecommendationType_AdminMFAV2),
		string(RecommendationRecommendationType_AppRoleAssignmentsGroups),
		string(RecommendationRecommendationType_AppRoleAssignmentsUsers),
		string(RecommendationRecommendationType_ApplicationCredentialExpiry),
		string(RecommendationRecommendationType_BlockLegacyAuthentication),
		string(RecommendationRecommendationType_EnableDesktopSSO),
		string(RecommendationRecommendationType_EnablePHS),
		string(RecommendationRecommendationType_EnableProvisioning),
		string(RecommendationRecommendationType_InactiveGuests),
		string(RecommendationRecommendationType_IntegratedApps),
		string(RecommendationRecommendationType_LongLivedCredentials),
		string(RecommendationRecommendationType_ManagedIdentity),
		string(RecommendationRecommendationType_MfaRegistrationV2),
		string(RecommendationRecommendationType_OneAdmin),
		string(RecommendationRecommendationType_OverprivilegedApps),
		string(RecommendationRecommendationType_OwnerlessApps),
		string(RecommendationRecommendationType_PasswordHashSync),
		string(RecommendationRecommendationType_PrivateLinkForAAD),
		string(RecommendationRecommendationType_PwagePolicyNew),
		string(RecommendationRecommendationType_RoleOverlap),
		string(RecommendationRecommendationType_SelfServicePasswordReset),
		string(RecommendationRecommendationType_ServicePrincipalKeyExpiry),
		string(RecommendationRecommendationType_SigninRiskPolicy),
		string(RecommendationRecommendationType_StaleAppCreds),
		string(RecommendationRecommendationType_StaleApps),
		string(RecommendationRecommendationType_SwitchFromPerUserMFA),
		string(RecommendationRecommendationType_TenantMFA),
		string(RecommendationRecommendationType_ThirdPartyApps),
		string(RecommendationRecommendationType_TurnOffPerUserMFA),
		string(RecommendationRecommendationType_UseAuthenticatorApp),
		string(RecommendationRecommendationType_UseMyApps),
		string(RecommendationRecommendationType_UserRiskPolicy),
		string(RecommendationRecommendationType_VerifyAppPublisher),
	}
}

func (s *RecommendationRecommendationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationRecommendationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationRecommendationType(input string) (*RecommendationRecommendationType, error) {
	vals := map[string]RecommendationRecommendationType{
		"aadconnectdeprecated":        RecommendationRecommendationType_AadConnectDeprecated,
		"adaltomsalmigration":         RecommendationRecommendationType_AdalToMsalMigration,
		"adfsappsmigration":           RecommendationRecommendationType_AdfsAppsMigration,
		"adminmfav2":                  RecommendationRecommendationType_AdminMFAV2,
		"approleassignmentsgroups":    RecommendationRecommendationType_AppRoleAssignmentsGroups,
		"approleassignmentsusers":     RecommendationRecommendationType_AppRoleAssignmentsUsers,
		"applicationcredentialexpiry": RecommendationRecommendationType_ApplicationCredentialExpiry,
		"blocklegacyauthentication":   RecommendationRecommendationType_BlockLegacyAuthentication,
		"enabledesktopsso":            RecommendationRecommendationType_EnableDesktopSSO,
		"enablephs":                   RecommendationRecommendationType_EnablePHS,
		"enableprovisioning":          RecommendationRecommendationType_EnableProvisioning,
		"inactiveguests":              RecommendationRecommendationType_InactiveGuests,
		"integratedapps":              RecommendationRecommendationType_IntegratedApps,
		"longlivedcredentials":        RecommendationRecommendationType_LongLivedCredentials,
		"managedidentity":             RecommendationRecommendationType_ManagedIdentity,
		"mfaregistrationv2":           RecommendationRecommendationType_MfaRegistrationV2,
		"oneadmin":                    RecommendationRecommendationType_OneAdmin,
		"overprivilegedapps":          RecommendationRecommendationType_OverprivilegedApps,
		"ownerlessapps":               RecommendationRecommendationType_OwnerlessApps,
		"passwordhashsync":            RecommendationRecommendationType_PasswordHashSync,
		"privatelinkforaad":           RecommendationRecommendationType_PrivateLinkForAAD,
		"pwagepolicynew":              RecommendationRecommendationType_PwagePolicyNew,
		"roleoverlap":                 RecommendationRecommendationType_RoleOverlap,
		"selfservicepasswordreset":    RecommendationRecommendationType_SelfServicePasswordReset,
		"serviceprincipalkeyexpiry":   RecommendationRecommendationType_ServicePrincipalKeyExpiry,
		"signinriskpolicy":            RecommendationRecommendationType_SigninRiskPolicy,
		"staleappcreds":               RecommendationRecommendationType_StaleAppCreds,
		"staleapps":                   RecommendationRecommendationType_StaleApps,
		"switchfromperusermfa":        RecommendationRecommendationType_SwitchFromPerUserMFA,
		"tenantmfa":                   RecommendationRecommendationType_TenantMFA,
		"thirdpartyapps":              RecommendationRecommendationType_ThirdPartyApps,
		"turnoffperusermfa":           RecommendationRecommendationType_TurnOffPerUserMFA,
		"useauthenticatorapp":         RecommendationRecommendationType_UseAuthenticatorApp,
		"usemyapps":                   RecommendationRecommendationType_UseMyApps,
		"userriskpolicy":              RecommendationRecommendationType_UserRiskPolicy,
		"verifyapppublisher":          RecommendationRecommendationType_VerifyAppPublisher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationRecommendationType(input)
	return &out, nil
}
