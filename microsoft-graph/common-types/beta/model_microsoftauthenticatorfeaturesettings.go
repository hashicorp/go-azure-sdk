package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorFeatureSettings struct {
	CompanionAppAllowedState                *AuthenticationMethodFeatureConfiguration `json:"companionAppAllowedState,omitempty"`
	DisplayAppInformationRequiredState      *AuthenticationMethodFeatureConfiguration `json:"displayAppInformationRequiredState,omitempty"`
	DisplayLocationInformationRequiredState *AuthenticationMethodFeatureConfiguration `json:"displayLocationInformationRequiredState,omitempty"`
	NumberMatchingRequiredState             *AuthenticationMethodFeatureConfiguration `json:"numberMatchingRequiredState,omitempty"`
	ODataType                               *string                                   `json:"@odata.type,omitempty"`
}
