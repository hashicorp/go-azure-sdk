package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorFeatureSettings struct {
	DisplayAppInformationRequiredState      *AuthenticationMethodFeatureConfiguration `json:"displayAppInformationRequiredState,omitempty"`
	DisplayLocationInformationRequiredState *AuthenticationMethodFeatureConfiguration `json:"displayLocationInformationRequiredState,omitempty"`
	ODataType                               *string                                   `json:"@odata.type,omitempty"`
}
