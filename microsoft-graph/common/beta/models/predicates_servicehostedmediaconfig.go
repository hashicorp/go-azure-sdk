package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHostedMediaConfigOperationPredicate struct {
	ODataType                   *string
	RemoveFromDefaultAudioGroup *bool
}

func (p ServiceHostedMediaConfigOperationPredicate) Matches(input ServiceHostedMediaConfig) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemoveFromDefaultAudioGroup != nil && (input.RemoveFromDefaultAudioGroup == nil || *p.RemoveFromDefaultAudioGroup != *input.RemoveFromDefaultAudioGroup) {
		return false
	}

	return true
}
