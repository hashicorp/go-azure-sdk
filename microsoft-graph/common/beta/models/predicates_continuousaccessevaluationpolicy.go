package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContinuousAccessEvaluationPolicyOperationPredicate struct {
	Description *string
	DisplayName *string
	Id          *string
	IsEnabled   *bool
	Migrate     *bool
	ODataType   *string
}

func (p ContinuousAccessEvaluationPolicyOperationPredicate) Matches(input ContinuousAccessEvaluationPolicy) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.Migrate != nil && (input.Migrate == nil || *p.Migrate != *input.Migrate) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
