package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkSpeakerConfigurationOperationPredicate struct {
	IsCommunicationSpeakerOptional *bool
	IsSpeakerOptional              *bool
	ODataType                      *string
}

func (p TeamworkSpeakerConfigurationOperationPredicate) Matches(input TeamworkSpeakerConfiguration) bool {

	if p.IsCommunicationSpeakerOptional != nil && (input.IsCommunicationSpeakerOptional == nil || *p.IsCommunicationSpeakerOptional != *input.IsCommunicationSpeakerOptional) {
		return false
	}

	if p.IsSpeakerOptional != nil && (input.IsSpeakerOptional == nil || *p.IsSpeakerOptional != *input.IsSpeakerOptional) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
