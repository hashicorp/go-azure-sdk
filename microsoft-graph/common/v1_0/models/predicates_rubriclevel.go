package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RubricLevelOperationPredicate struct {
	DisplayName *string
	LevelId     *string
	ODataType   *string
}

func (p RubricLevelOperationPredicate) Matches(input RubricLevel) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.LevelId != nil && (input.LevelId == nil || *p.LevelId != *input.LevelId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
