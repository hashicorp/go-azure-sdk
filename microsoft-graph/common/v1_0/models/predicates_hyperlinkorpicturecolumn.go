package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperlinkOrPictureColumnOperationPredicate struct {
	IsPicture *bool
	ODataType *string
}

func (p HyperlinkOrPictureColumnOperationPredicate) Matches(input HyperlinkOrPictureColumn) bool {

	if p.IsPicture != nil && (input.IsPicture == nil || *p.IsPicture != *input.IsPicture) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
