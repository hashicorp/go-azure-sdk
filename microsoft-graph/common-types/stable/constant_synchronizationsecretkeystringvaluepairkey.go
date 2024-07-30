package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationSecretKeyStringValuePairKey string

const (
	SynchronizationSecretKeyStringValuePairKey_AppKey                          SynchronizationSecretKeyStringValuePairKey = "AppKey"
	SynchronizationSecretKeyStringValuePairKey_ApplicationTemplateIdentifier   SynchronizationSecretKeyStringValuePairKey = "ApplicationTemplateIdentifier"
	SynchronizationSecretKeyStringValuePairKey_AuthenticationType              SynchronizationSecretKeyStringValuePairKey = "AuthenticationType"
	SynchronizationSecretKeyStringValuePairKey_BaseAddress                     SynchronizationSecretKeyStringValuePairKey = "BaseAddress"
	SynchronizationSecretKeyStringValuePairKey_ClientIdentifier                SynchronizationSecretKeyStringValuePairKey = "ClientIdentifier"
	SynchronizationSecretKeyStringValuePairKey_ClientSecret                    SynchronizationSecretKeyStringValuePairKey = "ClientSecret"
	SynchronizationSecretKeyStringValuePairKey_CompanyId                       SynchronizationSecretKeyStringValuePairKey = "CompanyId"
	SynchronizationSecretKeyStringValuePairKey_ConnectionString                SynchronizationSecretKeyStringValuePairKey = "ConnectionString"
	SynchronizationSecretKeyStringValuePairKey_ConsumerKey                     SynchronizationSecretKeyStringValuePairKey = "ConsumerKey"
	SynchronizationSecretKeyStringValuePairKey_ConsumerSecret                  SynchronizationSecretKeyStringValuePairKey = "ConsumerSecret"
	SynchronizationSecretKeyStringValuePairKey_Domain                          SynchronizationSecretKeyStringValuePairKey = "Domain"
	SynchronizationSecretKeyStringValuePairKey_EnforceDomain                   SynchronizationSecretKeyStringValuePairKey = "EnforceDomain"
	SynchronizationSecretKeyStringValuePairKey_HardDeletesEnabled              SynchronizationSecretKeyStringValuePairKey = "HardDeletesEnabled"
	SynchronizationSecretKeyStringValuePairKey_InstanceName                    SynchronizationSecretKeyStringValuePairKey = "InstanceName"
	SynchronizationSecretKeyStringValuePairKey_None                            SynchronizationSecretKeyStringValuePairKey = "None"
	SynchronizationSecretKeyStringValuePairKey_Oauth2AccessToken               SynchronizationSecretKeyStringValuePairKey = "Oauth2AccessToken"
	SynchronizationSecretKeyStringValuePairKey_Oauth2AccessTokenCreationTime   SynchronizationSecretKeyStringValuePairKey = "Oauth2AccessTokenCreationTime"
	SynchronizationSecretKeyStringValuePairKey_Oauth2AuthorizationCode         SynchronizationSecretKeyStringValuePairKey = "Oauth2AuthorizationCode"
	SynchronizationSecretKeyStringValuePairKey_Oauth2AuthorizationUri          SynchronizationSecretKeyStringValuePairKey = "Oauth2AuthorizationUri"
	SynchronizationSecretKeyStringValuePairKey_Oauth2ClientId                  SynchronizationSecretKeyStringValuePairKey = "Oauth2ClientId"
	SynchronizationSecretKeyStringValuePairKey_Oauth2ClientSecret              SynchronizationSecretKeyStringValuePairKey = "Oauth2ClientSecret"
	SynchronizationSecretKeyStringValuePairKey_Oauth2RedirectUri               SynchronizationSecretKeyStringValuePairKey = "Oauth2RedirectUri"
	SynchronizationSecretKeyStringValuePairKey_Oauth2RefreshToken              SynchronizationSecretKeyStringValuePairKey = "Oauth2RefreshToken"
	SynchronizationSecretKeyStringValuePairKey_Oauth2TokenExchangeUri          SynchronizationSecretKeyStringValuePairKey = "Oauth2TokenExchangeUri"
	SynchronizationSecretKeyStringValuePairKey_Password                        SynchronizationSecretKeyStringValuePairKey = "Password"
	SynchronizationSecretKeyStringValuePairKey_PerformInboundEntitlementGrants SynchronizationSecretKeyStringValuePairKey = "PerformInboundEntitlementGrants"
	SynchronizationSecretKeyStringValuePairKey_Sandbox                         SynchronizationSecretKeyStringValuePairKey = "Sandbox"
	SynchronizationSecretKeyStringValuePairKey_SandboxName                     SynchronizationSecretKeyStringValuePairKey = "SandboxName"
	SynchronizationSecretKeyStringValuePairKey_SecretToken                     SynchronizationSecretKeyStringValuePairKey = "SecretToken"
	SynchronizationSecretKeyStringValuePairKey_Server                          SynchronizationSecretKeyStringValuePairKey = "Server"
	SynchronizationSecretKeyStringValuePairKey_SingleSignOnType                SynchronizationSecretKeyStringValuePairKey = "SingleSignOnType"
	SynchronizationSecretKeyStringValuePairKey_SkipOutOfScopeDeletions         SynchronizationSecretKeyStringValuePairKey = "SkipOutOfScopeDeletions"
	SynchronizationSecretKeyStringValuePairKey_SyncAgentADContainer            SynchronizationSecretKeyStringValuePairKey = "SyncAgentADContainer"
	SynchronizationSecretKeyStringValuePairKey_SyncAgentCompatibilityKey       SynchronizationSecretKeyStringValuePairKey = "SyncAgentCompatibilityKey"
	SynchronizationSecretKeyStringValuePairKey_SyncAll                         SynchronizationSecretKeyStringValuePairKey = "SyncAll"
	SynchronizationSecretKeyStringValuePairKey_SyncNotificationSettings        SynchronizationSecretKeyStringValuePairKey = "SyncNotificationSettings"
	SynchronizationSecretKeyStringValuePairKey_SynchronizationSchedule         SynchronizationSecretKeyStringValuePairKey = "SynchronizationSchedule"
	SynchronizationSecretKeyStringValuePairKey_SystemOfRecord                  SynchronizationSecretKeyStringValuePairKey = "SystemOfRecord"
	SynchronizationSecretKeyStringValuePairKey_TestReferences                  SynchronizationSecretKeyStringValuePairKey = "TestReferences"
	SynchronizationSecretKeyStringValuePairKey_TokenExpiration                 SynchronizationSecretKeyStringValuePairKey = "TokenExpiration"
	SynchronizationSecretKeyStringValuePairKey_TokenKey                        SynchronizationSecretKeyStringValuePairKey = "TokenKey"
	SynchronizationSecretKeyStringValuePairKey_UpdateKeyOnSoftDelete           SynchronizationSecretKeyStringValuePairKey = "UpdateKeyOnSoftDelete"
	SynchronizationSecretKeyStringValuePairKey_Url                             SynchronizationSecretKeyStringValuePairKey = "Url"
	SynchronizationSecretKeyStringValuePairKey_UserName                        SynchronizationSecretKeyStringValuePairKey = "UserName"
	SynchronizationSecretKeyStringValuePairKey_ValidateDomain                  SynchronizationSecretKeyStringValuePairKey = "ValidateDomain"
)

func PossibleValuesForSynchronizationSecretKeyStringValuePairKey() []string {
	return []string{
		string(SynchronizationSecretKeyStringValuePairKey_AppKey),
		string(SynchronizationSecretKeyStringValuePairKey_ApplicationTemplateIdentifier),
		string(SynchronizationSecretKeyStringValuePairKey_AuthenticationType),
		string(SynchronizationSecretKeyStringValuePairKey_BaseAddress),
		string(SynchronizationSecretKeyStringValuePairKey_ClientIdentifier),
		string(SynchronizationSecretKeyStringValuePairKey_ClientSecret),
		string(SynchronizationSecretKeyStringValuePairKey_CompanyId),
		string(SynchronizationSecretKeyStringValuePairKey_ConnectionString),
		string(SynchronizationSecretKeyStringValuePairKey_ConsumerKey),
		string(SynchronizationSecretKeyStringValuePairKey_ConsumerSecret),
		string(SynchronizationSecretKeyStringValuePairKey_Domain),
		string(SynchronizationSecretKeyStringValuePairKey_EnforceDomain),
		string(SynchronizationSecretKeyStringValuePairKey_HardDeletesEnabled),
		string(SynchronizationSecretKeyStringValuePairKey_InstanceName),
		string(SynchronizationSecretKeyStringValuePairKey_None),
		string(SynchronizationSecretKeyStringValuePairKey_Oauth2AccessToken),
		string(SynchronizationSecretKeyStringValuePairKey_Oauth2AccessTokenCreationTime),
		string(SynchronizationSecretKeyStringValuePairKey_Oauth2AuthorizationCode),
		string(SynchronizationSecretKeyStringValuePairKey_Oauth2AuthorizationUri),
		string(SynchronizationSecretKeyStringValuePairKey_Oauth2ClientId),
		string(SynchronizationSecretKeyStringValuePairKey_Oauth2ClientSecret),
		string(SynchronizationSecretKeyStringValuePairKey_Oauth2RedirectUri),
		string(SynchronizationSecretKeyStringValuePairKey_Oauth2RefreshToken),
		string(SynchronizationSecretKeyStringValuePairKey_Oauth2TokenExchangeUri),
		string(SynchronizationSecretKeyStringValuePairKey_Password),
		string(SynchronizationSecretKeyStringValuePairKey_PerformInboundEntitlementGrants),
		string(SynchronizationSecretKeyStringValuePairKey_Sandbox),
		string(SynchronizationSecretKeyStringValuePairKey_SandboxName),
		string(SynchronizationSecretKeyStringValuePairKey_SecretToken),
		string(SynchronizationSecretKeyStringValuePairKey_Server),
		string(SynchronizationSecretKeyStringValuePairKey_SingleSignOnType),
		string(SynchronizationSecretKeyStringValuePairKey_SkipOutOfScopeDeletions),
		string(SynchronizationSecretKeyStringValuePairKey_SyncAgentADContainer),
		string(SynchronizationSecretKeyStringValuePairKey_SyncAgentCompatibilityKey),
		string(SynchronizationSecretKeyStringValuePairKey_SyncAll),
		string(SynchronizationSecretKeyStringValuePairKey_SyncNotificationSettings),
		string(SynchronizationSecretKeyStringValuePairKey_SynchronizationSchedule),
		string(SynchronizationSecretKeyStringValuePairKey_SystemOfRecord),
		string(SynchronizationSecretKeyStringValuePairKey_TestReferences),
		string(SynchronizationSecretKeyStringValuePairKey_TokenExpiration),
		string(SynchronizationSecretKeyStringValuePairKey_TokenKey),
		string(SynchronizationSecretKeyStringValuePairKey_UpdateKeyOnSoftDelete),
		string(SynchronizationSecretKeyStringValuePairKey_Url),
		string(SynchronizationSecretKeyStringValuePairKey_UserName),
		string(SynchronizationSecretKeyStringValuePairKey_ValidateDomain),
	}
}

func (s *SynchronizationSecretKeyStringValuePairKey) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationSecretKeyStringValuePairKey(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationSecretKeyStringValuePairKey(input string) (*SynchronizationSecretKeyStringValuePairKey, error) {
	vals := map[string]SynchronizationSecretKeyStringValuePairKey{
		"appkey":                          SynchronizationSecretKeyStringValuePairKey_AppKey,
		"applicationtemplateidentifier":   SynchronizationSecretKeyStringValuePairKey_ApplicationTemplateIdentifier,
		"authenticationtype":              SynchronizationSecretKeyStringValuePairKey_AuthenticationType,
		"baseaddress":                     SynchronizationSecretKeyStringValuePairKey_BaseAddress,
		"clientidentifier":                SynchronizationSecretKeyStringValuePairKey_ClientIdentifier,
		"clientsecret":                    SynchronizationSecretKeyStringValuePairKey_ClientSecret,
		"companyid":                       SynchronizationSecretKeyStringValuePairKey_CompanyId,
		"connectionstring":                SynchronizationSecretKeyStringValuePairKey_ConnectionString,
		"consumerkey":                     SynchronizationSecretKeyStringValuePairKey_ConsumerKey,
		"consumersecret":                  SynchronizationSecretKeyStringValuePairKey_ConsumerSecret,
		"domain":                          SynchronizationSecretKeyStringValuePairKey_Domain,
		"enforcedomain":                   SynchronizationSecretKeyStringValuePairKey_EnforceDomain,
		"harddeletesenabled":              SynchronizationSecretKeyStringValuePairKey_HardDeletesEnabled,
		"instancename":                    SynchronizationSecretKeyStringValuePairKey_InstanceName,
		"none":                            SynchronizationSecretKeyStringValuePairKey_None,
		"oauth2accesstoken":               SynchronizationSecretKeyStringValuePairKey_Oauth2AccessToken,
		"oauth2accesstokencreationtime":   SynchronizationSecretKeyStringValuePairKey_Oauth2AccessTokenCreationTime,
		"oauth2authorizationcode":         SynchronizationSecretKeyStringValuePairKey_Oauth2AuthorizationCode,
		"oauth2authorizationuri":          SynchronizationSecretKeyStringValuePairKey_Oauth2AuthorizationUri,
		"oauth2clientid":                  SynchronizationSecretKeyStringValuePairKey_Oauth2ClientId,
		"oauth2clientsecret":              SynchronizationSecretKeyStringValuePairKey_Oauth2ClientSecret,
		"oauth2redirecturi":               SynchronizationSecretKeyStringValuePairKey_Oauth2RedirectUri,
		"oauth2refreshtoken":              SynchronizationSecretKeyStringValuePairKey_Oauth2RefreshToken,
		"oauth2tokenexchangeuri":          SynchronizationSecretKeyStringValuePairKey_Oauth2TokenExchangeUri,
		"password":                        SynchronizationSecretKeyStringValuePairKey_Password,
		"performinboundentitlementgrants": SynchronizationSecretKeyStringValuePairKey_PerformInboundEntitlementGrants,
		"sandbox":                         SynchronizationSecretKeyStringValuePairKey_Sandbox,
		"sandboxname":                     SynchronizationSecretKeyStringValuePairKey_SandboxName,
		"secrettoken":                     SynchronizationSecretKeyStringValuePairKey_SecretToken,
		"server":                          SynchronizationSecretKeyStringValuePairKey_Server,
		"singlesignontype":                SynchronizationSecretKeyStringValuePairKey_SingleSignOnType,
		"skipoutofscopedeletions":         SynchronizationSecretKeyStringValuePairKey_SkipOutOfScopeDeletions,
		"syncagentadcontainer":            SynchronizationSecretKeyStringValuePairKey_SyncAgentADContainer,
		"syncagentcompatibilitykey":       SynchronizationSecretKeyStringValuePairKey_SyncAgentCompatibilityKey,
		"syncall":                         SynchronizationSecretKeyStringValuePairKey_SyncAll,
		"syncnotificationsettings":        SynchronizationSecretKeyStringValuePairKey_SyncNotificationSettings,
		"synchronizationschedule":         SynchronizationSecretKeyStringValuePairKey_SynchronizationSchedule,
		"systemofrecord":                  SynchronizationSecretKeyStringValuePairKey_SystemOfRecord,
		"testreferences":                  SynchronizationSecretKeyStringValuePairKey_TestReferences,
		"tokenexpiration":                 SynchronizationSecretKeyStringValuePairKey_TokenExpiration,
		"tokenkey":                        SynchronizationSecretKeyStringValuePairKey_TokenKey,
		"updatekeyonsoftdelete":           SynchronizationSecretKeyStringValuePairKey_UpdateKeyOnSoftDelete,
		"url":                             SynchronizationSecretKeyStringValuePairKey_Url,
		"username":                        SynchronizationSecretKeyStringValuePairKey_UserName,
		"validatedomain":                  SynchronizationSecretKeyStringValuePairKey_ValidateDomain,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationSecretKeyStringValuePairKey(input)
	return &out, nil
}
