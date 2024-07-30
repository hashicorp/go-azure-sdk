package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSKerberosSingleSignOnExtension struct {
	ActiveDirectorySiteCode                  *string   `json:"activeDirectorySiteCode,omitempty"`
	BlockActiveDirectorySiteAutoDiscovery    *bool     `json:"blockActiveDirectorySiteAutoDiscovery,omitempty"`
	BlockAutomaticLogin                      *bool     `json:"blockAutomaticLogin,omitempty"`
	CacheName                                *string   `json:"cacheName,omitempty"`
	CredentialBundleIdAccessControlList      *[]string `json:"credentialBundleIdAccessControlList,omitempty"`
	CredentialsCacheMonitored                *bool     `json:"credentialsCacheMonitored,omitempty"`
	DomainRealms                             *[]string `json:"domainRealms,omitempty"`
	Domains                                  *[]string `json:"domains,omitempty"`
	IsDefaultRealm                           *bool     `json:"isDefaultRealm,omitempty"`
	KerberosAppsInBundleIdACLIncluded        *bool     `json:"kerberosAppsInBundleIdACLIncluded,omitempty"`
	ManagedAppsInBundleIdACLIncluded         *bool     `json:"managedAppsInBundleIdACLIncluded,omitempty"`
	ModeCredentialUsed                       *string   `json:"modeCredentialUsed,omitempty"`
	ODataType                                *string   `json:"@odata.type,omitempty"`
	PasswordBlockModification                *bool     `json:"passwordBlockModification,omitempty"`
	PasswordChangeUrl                        *string   `json:"passwordChangeUrl,omitempty"`
	PasswordEnableLocalSync                  *bool     `json:"passwordEnableLocalSync,omitempty"`
	PasswordExpirationDays                   *int64    `json:"passwordExpirationDays,omitempty"`
	PasswordExpirationNotificationDays       *int64    `json:"passwordExpirationNotificationDays,omitempty"`
	PasswordMinimumAgeDays                   *int64    `json:"passwordMinimumAgeDays,omitempty"`
	PasswordMinimumLength                    *int64    `json:"passwordMinimumLength,omitempty"`
	PasswordPreviousPasswordBlockCount       *int64    `json:"passwordPreviousPasswordBlockCount,omitempty"`
	PasswordRequireActiveDirectoryComplexity *bool     `json:"passwordRequireActiveDirectoryComplexity,omitempty"`
	PasswordRequirementsDescription          *string   `json:"passwordRequirementsDescription,omitempty"`
	PreferredKDCs                            *[]string `json:"preferredKDCs,omitempty"`
	Realm                                    *string   `json:"realm,omitempty"`
	RequireUserPresence                      *bool     `json:"requireUserPresence,omitempty"`
	SignInHelpText                           *string   `json:"signInHelpText,omitempty"`
	TlsForLDAPRequired                       *bool     `json:"tlsForLDAPRequired,omitempty"`
	UserPrincipalName                        *string   `json:"userPrincipalName,omitempty"`
	UserSetupDelayed                         *bool     `json:"userSetupDelayed,omitempty"`
	UsernameLabelCustom                      *string   `json:"usernameLabelCustom,omitempty"`
}
