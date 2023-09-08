package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerExternalBucketSourceOperationPredicate struct {
	ContextScenarioId *string
	ExternalContextId *string
	ExternalObjectId  *string
	ODataType         *string
}

func (p PlannerExternalBucketSourceOperationPredicate) Matches(input PlannerExternalBucketSource) bool {

	if p.ContextScenarioId != nil && (input.ContextScenarioId == nil || *p.ContextScenarioId != *input.ContextScenarioId) {
		return false
	}

	if p.ExternalContextId != nil && (input.ExternalContextId == nil || *p.ExternalContextId != *input.ExternalContextId) {
		return false
	}

	if p.ExternalObjectId != nil && (input.ExternalObjectId == nil || *p.ExternalObjectId != *input.ExternalObjectId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
