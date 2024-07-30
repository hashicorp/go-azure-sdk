package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorAuthenticationMethodConfiguration struct {
	ExcludeTargets        *[]ExcludeTarget                                              `json:"excludeTargets,omitempty"`
	FeatureSettings       *MicrosoftAuthenticatorFeatureSettings                        `json:"featureSettings,omitempty"`
	Id                    *string                                                       `json:"id,omitempty"`
	IncludeTargets        *[]MicrosoftAuthenticatorAuthenticationMethodTarget           `json:"includeTargets,omitempty"`
	IsSoftwareOathEnabled *bool                                                         `json:"isSoftwareOathEnabled,omitempty"`
	ODataType             *string                                                       `json:"@odata.type,omitempty"`
	State                 *MicrosoftAuthenticatorAuthenticationMethodConfigurationState `json:"state,omitempty"`
}
