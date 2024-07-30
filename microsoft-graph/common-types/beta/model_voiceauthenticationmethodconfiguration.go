package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VoiceAuthenticationMethodConfiguration struct {
	ExcludeTargets       *[]ExcludeTarget                             `json:"excludeTargets,omitempty"`
	Id                   *string                                      `json:"id,omitempty"`
	IncludeTargets       *[]VoiceAuthenticationMethodTarget           `json:"includeTargets,omitempty"`
	IsOfficePhoneAllowed *bool                                        `json:"isOfficePhoneAllowed,omitempty"`
	ODataType            *string                                      `json:"@odata.type,omitempty"`
	State                *VoiceAuthenticationMethodConfigurationState `json:"state,omitempty"`
}
