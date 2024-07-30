package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationBaseRecommendationType string

const (
	RecommendationBaseRecommendationType_AadConnectDeprecated        RecommendationBaseRecommendationType = "aadConnectDeprecated"
	RecommendationBaseRecommendationType_AdalToMsalMigration         RecommendationBaseRecommendationType = "adalToMsalMigration"
	RecommendationBaseRecommendationType_AdfsAppsMigration           RecommendationBaseRecommendationType = "adfsAppsMigration"
	RecommendationBaseRecommendationType_AdminMFAV2                  RecommendationBaseRecommendationType = "adminMFAV2"
	RecommendationBaseRecommendationType_AppRoleAssignmentsGroups    RecommendationBaseRecommendationType = "appRoleAssignmentsGroups"
	RecommendationBaseRecommendationType_AppRoleAssignmentsUsers     RecommendationBaseRecommendationType = "appRoleAssignmentsUsers"
	RecommendationBaseRecommendationType_ApplicationCredentialExpiry RecommendationBaseRecommendationType = "applicationCredentialExpiry"
	RecommendationBaseRecommendationType_BlockLegacyAuthentication   RecommendationBaseRecommendationType = "blockLegacyAuthentication"
	RecommendationBaseRecommendationType_EnableDesktopSSO            RecommendationBaseRecommendationType = "enableDesktopSSO"
	RecommendationBaseRecommendationType_EnablePHS                   RecommendationBaseRecommendationType = "enablePHS"
	RecommendationBaseRecommendationType_EnableProvisioning          RecommendationBaseRecommendationType = "enableProvisioning"
	RecommendationBaseRecommendationType_InactiveGuests              RecommendationBaseRecommendationType = "inactiveGuests"
	RecommendationBaseRecommendationType_IntegratedApps              RecommendationBaseRecommendationType = "integratedApps"
	RecommendationBaseRecommendationType_LongLivedCredentials        RecommendationBaseRecommendationType = "longLivedCredentials"
	RecommendationBaseRecommendationType_ManagedIdentity             RecommendationBaseRecommendationType = "managedIdentity"
	RecommendationBaseRecommendationType_MfaRegistrationV2           RecommendationBaseRecommendationType = "mfaRegistrationV2"
	RecommendationBaseRecommendationType_OneAdmin                    RecommendationBaseRecommendationType = "oneAdmin"
	RecommendationBaseRecommendationType_OverprivilegedApps          RecommendationBaseRecommendationType = "overprivilegedApps"
	RecommendationBaseRecommendationType_OwnerlessApps               RecommendationBaseRecommendationType = "ownerlessApps"
	RecommendationBaseRecommendationType_PasswordHashSync            RecommendationBaseRecommendationType = "passwordHashSync"
	RecommendationBaseRecommendationType_PrivateLinkForAAD           RecommendationBaseRecommendationType = "privateLinkForAAD"
	RecommendationBaseRecommendationType_PwagePolicyNew              RecommendationBaseRecommendationType = "pwagePolicyNew"
	RecommendationBaseRecommendationType_RoleOverlap                 RecommendationBaseRecommendationType = "roleOverlap"
	RecommendationBaseRecommendationType_SelfServicePasswordReset    RecommendationBaseRecommendationType = "selfServicePasswordReset"
	RecommendationBaseRecommendationType_ServicePrincipalKeyExpiry   RecommendationBaseRecommendationType = "servicePrincipalKeyExpiry"
	RecommendationBaseRecommendationType_SigninRiskPolicy            RecommendationBaseRecommendationType = "signinRiskPolicy"
	RecommendationBaseRecommendationType_StaleAppCreds               RecommendationBaseRecommendationType = "staleAppCreds"
	RecommendationBaseRecommendationType_StaleApps                   RecommendationBaseRecommendationType = "staleApps"
	RecommendationBaseRecommendationType_SwitchFromPerUserMFA        RecommendationBaseRecommendationType = "switchFromPerUserMFA"
	RecommendationBaseRecommendationType_TenantMFA                   RecommendationBaseRecommendationType = "tenantMFA"
	RecommendationBaseRecommendationType_ThirdPartyApps              RecommendationBaseRecommendationType = "thirdPartyApps"
	RecommendationBaseRecommendationType_TurnOffPerUserMFA           RecommendationBaseRecommendationType = "turnOffPerUserMFA"
	RecommendationBaseRecommendationType_UseAuthenticatorApp         RecommendationBaseRecommendationType = "useAuthenticatorApp"
	RecommendationBaseRecommendationType_UseMyApps                   RecommendationBaseRecommendationType = "useMyApps"
	RecommendationBaseRecommendationType_UserRiskPolicy              RecommendationBaseRecommendationType = "userRiskPolicy"
	RecommendationBaseRecommendationType_VerifyAppPublisher          RecommendationBaseRecommendationType = "verifyAppPublisher"
)

func PossibleValuesForRecommendationBaseRecommendationType() []string {
	return []string{
		string(RecommendationBaseRecommendationType_AadConnectDeprecated),
		string(RecommendationBaseRecommendationType_AdalToMsalMigration),
		string(RecommendationBaseRecommendationType_AdfsAppsMigration),
		string(RecommendationBaseRecommendationType_AdminMFAV2),
		string(RecommendationBaseRecommendationType_AppRoleAssignmentsGroups),
		string(RecommendationBaseRecommendationType_AppRoleAssignmentsUsers),
		string(RecommendationBaseRecommendationType_ApplicationCredentialExpiry),
		string(RecommendationBaseRecommendationType_BlockLegacyAuthentication),
		string(RecommendationBaseRecommendationType_EnableDesktopSSO),
		string(RecommendationBaseRecommendationType_EnablePHS),
		string(RecommendationBaseRecommendationType_EnableProvisioning),
		string(RecommendationBaseRecommendationType_InactiveGuests),
		string(RecommendationBaseRecommendationType_IntegratedApps),
		string(RecommendationBaseRecommendationType_LongLivedCredentials),
		string(RecommendationBaseRecommendationType_ManagedIdentity),
		string(RecommendationBaseRecommendationType_MfaRegistrationV2),
		string(RecommendationBaseRecommendationType_OneAdmin),
		string(RecommendationBaseRecommendationType_OverprivilegedApps),
		string(RecommendationBaseRecommendationType_OwnerlessApps),
		string(RecommendationBaseRecommendationType_PasswordHashSync),
		string(RecommendationBaseRecommendationType_PrivateLinkForAAD),
		string(RecommendationBaseRecommendationType_PwagePolicyNew),
		string(RecommendationBaseRecommendationType_RoleOverlap),
		string(RecommendationBaseRecommendationType_SelfServicePasswordReset),
		string(RecommendationBaseRecommendationType_ServicePrincipalKeyExpiry),
		string(RecommendationBaseRecommendationType_SigninRiskPolicy),
		string(RecommendationBaseRecommendationType_StaleAppCreds),
		string(RecommendationBaseRecommendationType_StaleApps),
		string(RecommendationBaseRecommendationType_SwitchFromPerUserMFA),
		string(RecommendationBaseRecommendationType_TenantMFA),
		string(RecommendationBaseRecommendationType_ThirdPartyApps),
		string(RecommendationBaseRecommendationType_TurnOffPerUserMFA),
		string(RecommendationBaseRecommendationType_UseAuthenticatorApp),
		string(RecommendationBaseRecommendationType_UseMyApps),
		string(RecommendationBaseRecommendationType_UserRiskPolicy),
		string(RecommendationBaseRecommendationType_VerifyAppPublisher),
	}
}

func (s *RecommendationBaseRecommendationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationBaseRecommendationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationBaseRecommendationType(input string) (*RecommendationBaseRecommendationType, error) {
	vals := map[string]RecommendationBaseRecommendationType{
		"aadconnectdeprecated":        RecommendationBaseRecommendationType_AadConnectDeprecated,
		"adaltomsalmigration":         RecommendationBaseRecommendationType_AdalToMsalMigration,
		"adfsappsmigration":           RecommendationBaseRecommendationType_AdfsAppsMigration,
		"adminmfav2":                  RecommendationBaseRecommendationType_AdminMFAV2,
		"approleassignmentsgroups":    RecommendationBaseRecommendationType_AppRoleAssignmentsGroups,
		"approleassignmentsusers":     RecommendationBaseRecommendationType_AppRoleAssignmentsUsers,
		"applicationcredentialexpiry": RecommendationBaseRecommendationType_ApplicationCredentialExpiry,
		"blocklegacyauthentication":   RecommendationBaseRecommendationType_BlockLegacyAuthentication,
		"enabledesktopsso":            RecommendationBaseRecommendationType_EnableDesktopSSO,
		"enablephs":                   RecommendationBaseRecommendationType_EnablePHS,
		"enableprovisioning":          RecommendationBaseRecommendationType_EnableProvisioning,
		"inactiveguests":              RecommendationBaseRecommendationType_InactiveGuests,
		"integratedapps":              RecommendationBaseRecommendationType_IntegratedApps,
		"longlivedcredentials":        RecommendationBaseRecommendationType_LongLivedCredentials,
		"managedidentity":             RecommendationBaseRecommendationType_ManagedIdentity,
		"mfaregistrationv2":           RecommendationBaseRecommendationType_MfaRegistrationV2,
		"oneadmin":                    RecommendationBaseRecommendationType_OneAdmin,
		"overprivilegedapps":          RecommendationBaseRecommendationType_OverprivilegedApps,
		"ownerlessapps":               RecommendationBaseRecommendationType_OwnerlessApps,
		"passwordhashsync":            RecommendationBaseRecommendationType_PasswordHashSync,
		"privatelinkforaad":           RecommendationBaseRecommendationType_PrivateLinkForAAD,
		"pwagepolicynew":              RecommendationBaseRecommendationType_PwagePolicyNew,
		"roleoverlap":                 RecommendationBaseRecommendationType_RoleOverlap,
		"selfservicepasswordreset":    RecommendationBaseRecommendationType_SelfServicePasswordReset,
		"serviceprincipalkeyexpiry":   RecommendationBaseRecommendationType_ServicePrincipalKeyExpiry,
		"signinriskpolicy":            RecommendationBaseRecommendationType_SigninRiskPolicy,
		"staleappcreds":               RecommendationBaseRecommendationType_StaleAppCreds,
		"staleapps":                   RecommendationBaseRecommendationType_StaleApps,
		"switchfromperusermfa":        RecommendationBaseRecommendationType_SwitchFromPerUserMFA,
		"tenantmfa":                   RecommendationBaseRecommendationType_TenantMFA,
		"thirdpartyapps":              RecommendationBaseRecommendationType_ThirdPartyApps,
		"turnoffperusermfa":           RecommendationBaseRecommendationType_TurnOffPerUserMFA,
		"useauthenticatorapp":         RecommendationBaseRecommendationType_UseAuthenticatorApp,
		"usemyapps":                   RecommendationBaseRecommendationType_UseMyApps,
		"userriskpolicy":              RecommendationBaseRecommendationType_UserRiskPolicy,
		"verifyapppublisher":          RecommendationBaseRecommendationType_VerifyAppPublisher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationBaseRecommendationType(input)
	return &out, nil
}
