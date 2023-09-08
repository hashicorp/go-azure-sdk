package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerCollectionAssignmentTargetOperationPredicate struct {
	CollectionId *string
	ODataType    *string
}

func (p ConfigurationManagerCollectionAssignmentTargetOperationPredicate) Matches(input ConfigurationManagerCollectionAssignmentTarget) bool {

	if p.CollectionId != nil && (input.CollectionId == nil || *p.CollectionId != *input.CollectionId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
