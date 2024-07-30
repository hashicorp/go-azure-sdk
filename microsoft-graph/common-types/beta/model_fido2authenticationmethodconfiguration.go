package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2AuthenticationMethodConfiguration struct {
	ExcludeTargets                   *[]ExcludeTarget                             `json:"excludeTargets,omitempty"`
	Id                               *string                                      `json:"id,omitempty"`
	IsAttestationEnforced            *bool                                        `json:"isAttestationEnforced,omitempty"`
	IsSelfServiceRegistrationAllowed *bool                                        `json:"isSelfServiceRegistrationAllowed,omitempty"`
	KeyRestrictions                  *Fido2KeyRestrictions                        `json:"keyRestrictions,omitempty"`
	ODataType                        *string                                      `json:"@odata.type,omitempty"`
	State                            *Fido2AuthenticationMethodConfigurationState `json:"state,omitempty"`
}
