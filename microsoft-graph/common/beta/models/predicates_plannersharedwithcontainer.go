package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerSharedWithContainerOperationPredicate struct {
	ContainerId *string
	ODataType   *string
	Url         *string
}

func (p PlannerSharedWithContainerOperationPredicate) Matches(input PlannerSharedWithContainer) bool {

	if p.ContainerId != nil && (input.ContainerId == nil || *p.ContainerId != *input.ContainerId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	return true
}
