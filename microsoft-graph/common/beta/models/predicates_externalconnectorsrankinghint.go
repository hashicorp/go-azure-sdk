package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsRankingHintOperationPredicate struct {
	ODataType *string
}

func (p ExternalConnectorsRankingHintOperationPredicate) Matches(input ExternalConnectorsRankingHint) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
