package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorAuthenticationMethodTarget struct {
	AuthenticationMode     *MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode `json:"authenticationMode,omitempty"`
	Id                     *string                                                             `json:"id,omitempty"`
	IsRegistrationRequired *bool                                                               `json:"isRegistrationRequired,omitempty"`
	ODataType              *string                                                             `json:"@odata.type,omitempty"`
	TargetType             *MicrosoftAuthenticatorAuthenticationMethodTargetTargetType         `json:"targetType,omitempty"`
}
