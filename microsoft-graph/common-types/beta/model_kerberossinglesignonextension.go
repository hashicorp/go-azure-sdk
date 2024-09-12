package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SingleSignOnExtension = KerberosSingleSignOnExtension{}

type KerberosSingleSignOnExtension struct {
	// Gets or sets the Active Directory site.
	ActiveDirectorySiteCode nullable.Type[string] `json:"activeDirectorySiteCode,omitempty"`

	// Enables or disables whether the Kerberos extension can automatically determine its site name.
	BlockActiveDirectorySiteAutoDiscovery *bool `json:"blockActiveDirectorySiteAutoDiscovery,omitempty"`

	// Enables or disables Keychain usage.
	BlockAutomaticLogin *bool `json:"blockAutomaticLogin,omitempty"`

	// Gets or sets the Generic Security Services name of the Kerberos cache to use for this profile.
	CacheName nullable.Type[string] `json:"cacheName,omitempty"`

	// Gets or sets a list of app Bundle IDs allowed to access the Kerberos Ticket Granting Ticket.
	CredentialBundleIdAccessControlList *[]string `json:"credentialBundleIdAccessControlList,omitempty"`

	// Gets or sets a list of realms for custom domain-realm mapping. Realms are case sensitive.
	DomainRealms *[]string `json:"domainRealms,omitempty"`

	// Gets or sets a list of hosts or domain names for which the app extension performs SSO.
	Domains *[]string `json:"domains,omitempty"`

	// When true, this profile's realm will be selected as the default. Necessary if multiple Kerberos-type profiles are
	// configured.
	IsDefaultRealm *bool `json:"isDefaultRealm,omitempty"`

	// Enables or disables password changes.
	PasswordBlockModification *bool `json:"passwordBlockModification,omitempty"`

	// Gets or sets the URL that the user will be sent to when they initiate a password change.
	PasswordChangeUrl nullable.Type[string] `json:"passwordChangeUrl,omitempty"`

	// Enables or disables password syncing. This won't affect users logged in with a mobile account on macOS.
	PasswordEnableLocalSync *bool `json:"passwordEnableLocalSync,omitempty"`

	// Overrides the default password expiration in days. For most domains, this value is calculated automatically.
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// Gets or sets the number of days until the user is notified that their password will expire (default is 15).
	PasswordExpirationNotificationDays nullable.Type[int64] `json:"passwordExpirationNotificationDays,omitempty"`

	// Gets or sets the minimum number of days until a user can change their password again.
	PasswordMinimumAgeDays nullable.Type[int64] `json:"passwordMinimumAgeDays,omitempty"`

	// Gets or sets the minimum length of a password.
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Gets or sets the number of previous passwords to block.
	PasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"passwordPreviousPasswordBlockCount,omitempty"`

	// Enables or disables whether passwords must meet Active Directory's complexity requirements.
	PasswordRequireActiveDirectoryComplexity *bool `json:"passwordRequireActiveDirectoryComplexity,omitempty"`

	// Gets or sets a description of the password complexity requirements.
	PasswordRequirementsDescription nullable.Type[string] `json:"passwordRequirementsDescription,omitempty"`

	// Gets or sets the case-sensitive realm name for this profile.
	Realm *string `json:"realm,omitempty"`

	// Gets or sets whether to require authentication via Touch ID, Face ID, or a passcode to access the keychain entry.
	RequireUserPresence *bool `json:"requireUserPresence,omitempty"`

	// Gets or sets the principle user name to use for this profile. The realm name does not need to be included.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// Fields inherited from SingleSignOnExtension

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s KerberosSingleSignOnExtension) SingleSignOnExtension() BaseSingleSignOnExtensionImpl {
	return BaseSingleSignOnExtensionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = KerberosSingleSignOnExtension{}

func (s KerberosSingleSignOnExtension) MarshalJSON() ([]byte, error) {
	type wrapper KerberosSingleSignOnExtension
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KerberosSingleSignOnExtension: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KerberosSingleSignOnExtension: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.kerberosSingleSignOnExtension"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KerberosSingleSignOnExtension: %+v", err)
	}

	return encoded, nil
}
