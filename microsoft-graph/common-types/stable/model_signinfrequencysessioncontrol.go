package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInFrequencySessionControl struct {
	AuthenticationType *SignInFrequencySessionControlAuthenticationType `json:"authenticationType,omitempty"`
	FrequencyInterval  *SignInFrequencySessionControlFrequencyInterval  `json:"frequencyInterval,omitempty"`
	IsEnabled          *bool                                            `json:"isEnabled,omitempty"`
	ODataType          *string                                          `json:"@odata.type,omitempty"`
	Type               *SignInFrequencySessionControlType               `json:"type,omitempty"`
	Value              *int64                                           `json:"value,omitempty"`
}
