package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoginPageTextVisibilitySettings struct {
	HideAccountResetCredentials *bool   `json:"hideAccountResetCredentials,omitempty"`
	HideCannotAccessYourAccount *bool   `json:"hideCannotAccessYourAccount,omitempty"`
	HideForgotMyPassword        *bool   `json:"hideForgotMyPassword,omitempty"`
	HidePrivacyAndCookies       *bool   `json:"hidePrivacyAndCookies,omitempty"`
	HideResetItNow              *bool   `json:"hideResetItNow,omitempty"`
	HideTermsOfUse              *bool   `json:"hideTermsOfUse,omitempty"`
	ODataType                   *string `json:"@odata.type,omitempty"`
}
