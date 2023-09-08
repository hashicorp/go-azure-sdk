package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDestinationSummaryOperationPredicate struct {
	Count       *int64
	Destination *string
	ODataType   *string
}

func (p NetworkaccessDestinationSummaryOperationPredicate) Matches(input NetworkaccessDestinationSummary) bool {

	if p.Count != nil && (input.Count == nil || *p.Count != *input.Count) {
		return false
	}

	if p.Destination != nil && (input.Destination == nil || *p.Destination != *input.Destination) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
