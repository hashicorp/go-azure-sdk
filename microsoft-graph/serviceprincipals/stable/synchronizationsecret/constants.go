package synchronizationsecret

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationSecretKeyStringValuePairKey string

const (
	SynchronizationSecretKeyStringValuePairKeyAppKey                          SynchronizationSecretKeyStringValuePairKey = "AppKey"
	SynchronizationSecretKeyStringValuePairKeyApplicationTemplateIdentifier   SynchronizationSecretKeyStringValuePairKey = "ApplicationTemplateIdentifier"
	SynchronizationSecretKeyStringValuePairKeyAuthenticationType              SynchronizationSecretKeyStringValuePairKey = "AuthenticationType"
	SynchronizationSecretKeyStringValuePairKeyBaseAddress                     SynchronizationSecretKeyStringValuePairKey = "BaseAddress"
	SynchronizationSecretKeyStringValuePairKeyClientIdentifier                SynchronizationSecretKeyStringValuePairKey = "ClientIdentifier"
	SynchronizationSecretKeyStringValuePairKeyClientSecret                    SynchronizationSecretKeyStringValuePairKey = "ClientSecret"
	SynchronizationSecretKeyStringValuePairKeyCompanyId                       SynchronizationSecretKeyStringValuePairKey = "CompanyId"
	SynchronizationSecretKeyStringValuePairKeyConnectionString                SynchronizationSecretKeyStringValuePairKey = "ConnectionString"
	SynchronizationSecretKeyStringValuePairKeyConsumerKey                     SynchronizationSecretKeyStringValuePairKey = "ConsumerKey"
	SynchronizationSecretKeyStringValuePairKeyConsumerSecret                  SynchronizationSecretKeyStringValuePairKey = "ConsumerSecret"
	SynchronizationSecretKeyStringValuePairKeyDomain                          SynchronizationSecretKeyStringValuePairKey = "Domain"
	SynchronizationSecretKeyStringValuePairKeyEnforceDomain                   SynchronizationSecretKeyStringValuePairKey = "EnforceDomain"
	SynchronizationSecretKeyStringValuePairKeyHardDeletesEnabled              SynchronizationSecretKeyStringValuePairKey = "HardDeletesEnabled"
	SynchronizationSecretKeyStringValuePairKeyInstanceName                    SynchronizationSecretKeyStringValuePairKey = "InstanceName"
	SynchronizationSecretKeyStringValuePairKeyNone                            SynchronizationSecretKeyStringValuePairKey = "None"
	SynchronizationSecretKeyStringValuePairKeyOauth2AccessToken               SynchronizationSecretKeyStringValuePairKey = "Oauth2AccessToken"
	SynchronizationSecretKeyStringValuePairKeyOauth2AccessTokenCreationTime   SynchronizationSecretKeyStringValuePairKey = "Oauth2AccessTokenCreationTime"
	SynchronizationSecretKeyStringValuePairKeyOauth2AuthorizationCode         SynchronizationSecretKeyStringValuePairKey = "Oauth2AuthorizationCode"
	SynchronizationSecretKeyStringValuePairKeyOauth2AuthorizationUri          SynchronizationSecretKeyStringValuePairKey = "Oauth2AuthorizationUri"
	SynchronizationSecretKeyStringValuePairKeyOauth2ClientId                  SynchronizationSecretKeyStringValuePairKey = "Oauth2ClientId"
	SynchronizationSecretKeyStringValuePairKeyOauth2ClientSecret              SynchronizationSecretKeyStringValuePairKey = "Oauth2ClientSecret"
	SynchronizationSecretKeyStringValuePairKeyOauth2RedirectUri               SynchronizationSecretKeyStringValuePairKey = "Oauth2RedirectUri"
	SynchronizationSecretKeyStringValuePairKeyOauth2RefreshToken              SynchronizationSecretKeyStringValuePairKey = "Oauth2RefreshToken"
	SynchronizationSecretKeyStringValuePairKeyOauth2TokenExchangeUri          SynchronizationSecretKeyStringValuePairKey = "Oauth2TokenExchangeUri"
	SynchronizationSecretKeyStringValuePairKeyPassword                        SynchronizationSecretKeyStringValuePairKey = "Password"
	SynchronizationSecretKeyStringValuePairKeyPerformInboundEntitlementGrants SynchronizationSecretKeyStringValuePairKey = "PerformInboundEntitlementGrants"
	SynchronizationSecretKeyStringValuePairKeySandbox                         SynchronizationSecretKeyStringValuePairKey = "Sandbox"
	SynchronizationSecretKeyStringValuePairKeySandboxName                     SynchronizationSecretKeyStringValuePairKey = "SandboxName"
	SynchronizationSecretKeyStringValuePairKeySecretToken                     SynchronizationSecretKeyStringValuePairKey = "SecretToken"
	SynchronizationSecretKeyStringValuePairKeyServer                          SynchronizationSecretKeyStringValuePairKey = "Server"
	SynchronizationSecretKeyStringValuePairKeySingleSignOnType                SynchronizationSecretKeyStringValuePairKey = "SingleSignOnType"
	SynchronizationSecretKeyStringValuePairKeySkipOutOfScopeDeletions         SynchronizationSecretKeyStringValuePairKey = "SkipOutOfScopeDeletions"
	SynchronizationSecretKeyStringValuePairKeySyncAgentADContainer            SynchronizationSecretKeyStringValuePairKey = "SyncAgentADContainer"
	SynchronizationSecretKeyStringValuePairKeySyncAgentCompatibilityKey       SynchronizationSecretKeyStringValuePairKey = "SyncAgentCompatibilityKey"
	SynchronizationSecretKeyStringValuePairKeySyncAll                         SynchronizationSecretKeyStringValuePairKey = "SyncAll"
	SynchronizationSecretKeyStringValuePairKeySyncNotificationSettings        SynchronizationSecretKeyStringValuePairKey = "SyncNotificationSettings"
	SynchronizationSecretKeyStringValuePairKeySynchronizationSchedule         SynchronizationSecretKeyStringValuePairKey = "SynchronizationSchedule"
	SynchronizationSecretKeyStringValuePairKeySystemOfRecord                  SynchronizationSecretKeyStringValuePairKey = "SystemOfRecord"
	SynchronizationSecretKeyStringValuePairKeyTestReferences                  SynchronizationSecretKeyStringValuePairKey = "TestReferences"
	SynchronizationSecretKeyStringValuePairKeyTokenExpiration                 SynchronizationSecretKeyStringValuePairKey = "TokenExpiration"
	SynchronizationSecretKeyStringValuePairKeyTokenKey                        SynchronizationSecretKeyStringValuePairKey = "TokenKey"
	SynchronizationSecretKeyStringValuePairKeyUpdateKeyOnSoftDelete           SynchronizationSecretKeyStringValuePairKey = "UpdateKeyOnSoftDelete"
	SynchronizationSecretKeyStringValuePairKeyUrl                             SynchronizationSecretKeyStringValuePairKey = "Url"
	SynchronizationSecretKeyStringValuePairKeyUserName                        SynchronizationSecretKeyStringValuePairKey = "UserName"
	SynchronizationSecretKeyStringValuePairKeyValidateDomain                  SynchronizationSecretKeyStringValuePairKey = "ValidateDomain"
)

func PossibleValuesForSynchronizationSecretKeyStringValuePairKey() []string {
	return []string{
		string(SynchronizationSecretKeyStringValuePairKeyAppKey),
		string(SynchronizationSecretKeyStringValuePairKeyApplicationTemplateIdentifier),
		string(SynchronizationSecretKeyStringValuePairKeyAuthenticationType),
		string(SynchronizationSecretKeyStringValuePairKeyBaseAddress),
		string(SynchronizationSecretKeyStringValuePairKeyClientIdentifier),
		string(SynchronizationSecretKeyStringValuePairKeyClientSecret),
		string(SynchronizationSecretKeyStringValuePairKeyCompanyId),
		string(SynchronizationSecretKeyStringValuePairKeyConnectionString),
		string(SynchronizationSecretKeyStringValuePairKeyConsumerKey),
		string(SynchronizationSecretKeyStringValuePairKeyConsumerSecret),
		string(SynchronizationSecretKeyStringValuePairKeyDomain),
		string(SynchronizationSecretKeyStringValuePairKeyEnforceDomain),
		string(SynchronizationSecretKeyStringValuePairKeyHardDeletesEnabled),
		string(SynchronizationSecretKeyStringValuePairKeyInstanceName),
		string(SynchronizationSecretKeyStringValuePairKeyNone),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2AccessToken),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2AccessTokenCreationTime),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2AuthorizationCode),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2AuthorizationUri),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2ClientId),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2ClientSecret),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2RedirectUri),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2RefreshToken),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2TokenExchangeUri),
		string(SynchronizationSecretKeyStringValuePairKeyPassword),
		string(SynchronizationSecretKeyStringValuePairKeyPerformInboundEntitlementGrants),
		string(SynchronizationSecretKeyStringValuePairKeySandbox),
		string(SynchronizationSecretKeyStringValuePairKeySandboxName),
		string(SynchronizationSecretKeyStringValuePairKeySecretToken),
		string(SynchronizationSecretKeyStringValuePairKeyServer),
		string(SynchronizationSecretKeyStringValuePairKeySingleSignOnType),
		string(SynchronizationSecretKeyStringValuePairKeySkipOutOfScopeDeletions),
		string(SynchronizationSecretKeyStringValuePairKeySyncAgentADContainer),
		string(SynchronizationSecretKeyStringValuePairKeySyncAgentCompatibilityKey),
		string(SynchronizationSecretKeyStringValuePairKeySyncAll),
		string(SynchronizationSecretKeyStringValuePairKeySyncNotificationSettings),
		string(SynchronizationSecretKeyStringValuePairKeySynchronizationSchedule),
		string(SynchronizationSecretKeyStringValuePairKeySystemOfRecord),
		string(SynchronizationSecretKeyStringValuePairKeyTestReferences),
		string(SynchronizationSecretKeyStringValuePairKeyTokenExpiration),
		string(SynchronizationSecretKeyStringValuePairKeyTokenKey),
		string(SynchronizationSecretKeyStringValuePairKeyUpdateKeyOnSoftDelete),
		string(SynchronizationSecretKeyStringValuePairKeyUrl),
		string(SynchronizationSecretKeyStringValuePairKeyUserName),
		string(SynchronizationSecretKeyStringValuePairKeyValidateDomain),
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
		"appkey":                          SynchronizationSecretKeyStringValuePairKeyAppKey,
		"applicationtemplateidentifier":   SynchronizationSecretKeyStringValuePairKeyApplicationTemplateIdentifier,
		"authenticationtype":              SynchronizationSecretKeyStringValuePairKeyAuthenticationType,
		"baseaddress":                     SynchronizationSecretKeyStringValuePairKeyBaseAddress,
		"clientidentifier":                SynchronizationSecretKeyStringValuePairKeyClientIdentifier,
		"clientsecret":                    SynchronizationSecretKeyStringValuePairKeyClientSecret,
		"companyid":                       SynchronizationSecretKeyStringValuePairKeyCompanyId,
		"connectionstring":                SynchronizationSecretKeyStringValuePairKeyConnectionString,
		"consumerkey":                     SynchronizationSecretKeyStringValuePairKeyConsumerKey,
		"consumersecret":                  SynchronizationSecretKeyStringValuePairKeyConsumerSecret,
		"domain":                          SynchronizationSecretKeyStringValuePairKeyDomain,
		"enforcedomain":                   SynchronizationSecretKeyStringValuePairKeyEnforceDomain,
		"harddeletesenabled":              SynchronizationSecretKeyStringValuePairKeyHardDeletesEnabled,
		"instancename":                    SynchronizationSecretKeyStringValuePairKeyInstanceName,
		"none":                            SynchronizationSecretKeyStringValuePairKeyNone,
		"oauth2accesstoken":               SynchronizationSecretKeyStringValuePairKeyOauth2AccessToken,
		"oauth2accesstokencreationtime":   SynchronizationSecretKeyStringValuePairKeyOauth2AccessTokenCreationTime,
		"oauth2authorizationcode":         SynchronizationSecretKeyStringValuePairKeyOauth2AuthorizationCode,
		"oauth2authorizationuri":          SynchronizationSecretKeyStringValuePairKeyOauth2AuthorizationUri,
		"oauth2clientid":                  SynchronizationSecretKeyStringValuePairKeyOauth2ClientId,
		"oauth2clientsecret":              SynchronizationSecretKeyStringValuePairKeyOauth2ClientSecret,
		"oauth2redirecturi":               SynchronizationSecretKeyStringValuePairKeyOauth2RedirectUri,
		"oauth2refreshtoken":              SynchronizationSecretKeyStringValuePairKeyOauth2RefreshToken,
		"oauth2tokenexchangeuri":          SynchronizationSecretKeyStringValuePairKeyOauth2TokenExchangeUri,
		"password":                        SynchronizationSecretKeyStringValuePairKeyPassword,
		"performinboundentitlementgrants": SynchronizationSecretKeyStringValuePairKeyPerformInboundEntitlementGrants,
		"sandbox":                         SynchronizationSecretKeyStringValuePairKeySandbox,
		"sandboxname":                     SynchronizationSecretKeyStringValuePairKeySandboxName,
		"secrettoken":                     SynchronizationSecretKeyStringValuePairKeySecretToken,
		"server":                          SynchronizationSecretKeyStringValuePairKeyServer,
		"singlesignontype":                SynchronizationSecretKeyStringValuePairKeySingleSignOnType,
		"skipoutofscopedeletions":         SynchronizationSecretKeyStringValuePairKeySkipOutOfScopeDeletions,
		"syncagentadcontainer":            SynchronizationSecretKeyStringValuePairKeySyncAgentADContainer,
		"syncagentcompatibilitykey":       SynchronizationSecretKeyStringValuePairKeySyncAgentCompatibilityKey,
		"syncall":                         SynchronizationSecretKeyStringValuePairKeySyncAll,
		"syncnotificationsettings":        SynchronizationSecretKeyStringValuePairKeySyncNotificationSettings,
		"synchronizationschedule":         SynchronizationSecretKeyStringValuePairKeySynchronizationSchedule,
		"systemofrecord":                  SynchronizationSecretKeyStringValuePairKeySystemOfRecord,
		"testreferences":                  SynchronizationSecretKeyStringValuePairKeyTestReferences,
		"tokenexpiration":                 SynchronizationSecretKeyStringValuePairKeyTokenExpiration,
		"tokenkey":                        SynchronizationSecretKeyStringValuePairKeyTokenKey,
		"updatekeyonsoftdelete":           SynchronizationSecretKeyStringValuePairKeyUpdateKeyOnSoftDelete,
		"url":                             SynchronizationSecretKeyStringValuePairKeyUrl,
		"username":                        SynchronizationSecretKeyStringValuePairKeyUserName,
		"validatedomain":                  SynchronizationSecretKeyStringValuePairKeyValidateDomain,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationSecretKeyStringValuePairKey(input)
	return &out, nil
}
