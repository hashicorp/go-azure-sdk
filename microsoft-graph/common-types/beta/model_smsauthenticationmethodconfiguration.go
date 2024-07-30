package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SmsAuthenticationMethodConfiguration struct {
	ExcludeTargets *[]ExcludeTarget                           `json:"excludeTargets,omitempty"`
	Id             *string                                    `json:"id,omitempty"`
	IncludeTargets *[]SmsAuthenticationMethodTarget           `json:"includeTargets,omitempty"`
	ODataType      *string                                    `json:"@odata.type,omitempty"`
	State          *SmsAuthenticationMethodConfigurationState `json:"state,omitempty"`
}
