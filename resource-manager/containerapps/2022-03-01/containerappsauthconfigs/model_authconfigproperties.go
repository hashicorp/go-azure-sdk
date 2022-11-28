package containerappsauthconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthConfigProperties struct {
	GlobalValidation  *GlobalValidation  `json:"globalValidation"`
	HTTPSettings      *HTTPSettings      `json:"httpSettings"`
	IdentityProviders *IdentityProviders `json:"identityProviders"`
	Login             *Login             `json:"login"`
	Platform          *AuthPlatform      `json:"platform"`
}
