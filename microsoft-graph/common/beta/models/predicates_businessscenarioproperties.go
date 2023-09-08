package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioPropertiesOperationPredicate struct {
	ExternalBucketId      *string
	ExternalContextId     *string
	ExternalObjectId      *string
	ExternalObjectVersion *string
	ODataType             *string
	WebUrl                *string
}

func (p BusinessScenarioPropertiesOperationPredicate) Matches(input BusinessScenarioProperties) bool {

	if p.ExternalBucketId != nil && (input.ExternalBucketId == nil || *p.ExternalBucketId != *input.ExternalBucketId) {
		return false
	}

	if p.ExternalContextId != nil && (input.ExternalContextId == nil || *p.ExternalContextId != *input.ExternalContextId) {
		return false
	}

	if p.ExternalObjectId != nil && (input.ExternalObjectId == nil || *p.ExternalObjectId != *input.ExternalObjectId) {
		return false
	}

	if p.ExternalObjectVersion != nil && (input.ExternalObjectVersion == nil || *p.ExternalObjectVersion != *input.ExternalObjectVersion) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
