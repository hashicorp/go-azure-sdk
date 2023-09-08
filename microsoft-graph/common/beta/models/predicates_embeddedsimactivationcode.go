package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMActivationCodeOperationPredicate struct {
	IntegratedCircuitCardIdentifier *string
	MatchingIdentifier              *string
	ODataType                       *string
	SmdpPlusServerAddress           *string
}

func (p EmbeddedSIMActivationCodeOperationPredicate) Matches(input EmbeddedSIMActivationCode) bool {

	if p.IntegratedCircuitCardIdentifier != nil && (input.IntegratedCircuitCardIdentifier == nil || *p.IntegratedCircuitCardIdentifier != *input.IntegratedCircuitCardIdentifier) {
		return false
	}

	if p.MatchingIdentifier != nil && (input.MatchingIdentifier == nil || *p.MatchingIdentifier != *input.MatchingIdentifier) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SmdpPlusServerAddress != nil && (input.SmdpPlusServerAddress == nil || *p.SmdpPlusServerAddress != *input.SmdpPlusServerAddress) {
		return false
	}

	return true
}
