package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkMicrophoneConfigurationOperationPredicate struct {
	IsMicrophoneOptional *bool
	ODataType            *string
}

func (p TeamworkMicrophoneConfigurationOperationPredicate) Matches(input TeamworkMicrophoneConfiguration) bool {

	if p.IsMicrophoneOptional != nil && (input.IsMicrophoneOptional == nil || *p.IsMicrophoneOptional != *input.IsMicrophoneOptional) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
