package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInPreferences struct {
	IsSystemPreferredAuthenticationMethodEnabled  *bool                                                           `json:"isSystemPreferredAuthenticationMethodEnabled,omitempty"`
	ODataType                                     *string                                                         `json:"@odata.type,omitempty"`
	UserPreferredMethodForSecondaryAuthentication *SignInPreferencesUserPreferredMethodForSecondaryAuthentication `json:"userPreferredMethodForSecondaryAuthentication,omitempty"`
}
