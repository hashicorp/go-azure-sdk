package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppDetectionOperationPredicate struct {
	ODataType *string
}

func (p Win32LobAppDetectionOperationPredicate) Matches(input Win32LobAppDetection) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
