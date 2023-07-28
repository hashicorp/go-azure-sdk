package featuresetversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeaturesetJobOperationPredicate struct {
	CreatedDate  *string
	DisplayName  *string
	Duration     *string
	ExperimentId *string
	JobId        *string
}

func (p FeaturesetJobOperationPredicate) Matches(input FeaturesetJob) bool {

	if p.CreatedDate != nil && (input.CreatedDate == nil || *p.CreatedDate != *input.CreatedDate) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Duration != nil && (input.Duration == nil || *p.Duration != *input.Duration) {
		return false
	}

	if p.ExperimentId != nil && (input.ExperimentId == nil || *p.ExperimentId != *input.ExperimentId) {
		return false
	}

	if p.JobId != nil && (input.JobId == nil || *p.JobId != *input.JobId) {
		return false
	}

	return true
}

type FeaturesetVersionResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p FeaturesetVersionResourceOperationPredicate) Matches(input FeaturesetVersionResource) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
