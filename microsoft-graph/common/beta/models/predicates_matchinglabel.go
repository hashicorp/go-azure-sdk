package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MatchingLabelOperationPredicate struct {
	Description                 *string
	DisplayName                 *string
	Id                          *string
	IsEndpointProtectionEnabled *bool
	Name                        *string
	ODataType                   *string
	PolicyTip                   *string
	Priority                    *int64
	ToolTip                     *string
}

func (p MatchingLabelOperationPredicate) Matches(input MatchingLabel) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsEndpointProtectionEnabled != nil && (input.IsEndpointProtectionEnabled == nil || *p.IsEndpointProtectionEnabled != *input.IsEndpointProtectionEnabled) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PolicyTip != nil && (input.PolicyTip == nil || *p.PolicyTip != *input.PolicyTip) {
		return false
	}

	if p.Priority != nil && (input.Priority == nil || *p.Priority != *input.Priority) {
		return false
	}

	if p.ToolTip != nil && (input.ToolTip == nil || *p.ToolTip != *input.ToolTip) {
		return false
	}

	return true
}
