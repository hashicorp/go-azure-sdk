package containerappsauthconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthConfigProperties struct {
	GlobalValidation  *GlobalValidation  `json:"globalValidation,omitempty"`
	HTTPSettings      *HTTPSettings      `json:"httpSettings,omitempty"`
	IdentityProviders *IdentityProviders `json:"identityProviders,omitempty"`
	Login             *Login             `json:"login,omitempty"`
	Platform          *AuthPlatform      `json:"platform,omitempty"`
}
