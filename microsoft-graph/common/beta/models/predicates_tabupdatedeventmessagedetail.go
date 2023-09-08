package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TabUpdatedEventMessageDetailOperationPredicate struct {
	ODataType *string
	TabId     *string
}

func (p TabUpdatedEventMessageDetailOperationPredicate) Matches(input TabUpdatedEventMessageDetail) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TabId != nil && (input.TabId == nil || *p.TabId != *input.TabId) {
		return false
	}

	return true
}
