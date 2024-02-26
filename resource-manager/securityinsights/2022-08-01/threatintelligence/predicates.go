package threatintelligence

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatIntelligenceInformationOperationPredicate struct {
}

func (p ThreatIntelligenceInformationOperationPredicate) Matches(input ThreatIntelligenceInformation) bool {

	return true
}

type ThreatIntelligenceInformationListOperationPredicate struct {
	NextLink *string
}

func (p ThreatIntelligenceInformationListOperationPredicate) Matches(input ThreatIntelligenceInformationList) bool {

	if p.NextLink != nil && (input.NextLink == nil || *p.NextLink != *input.NextLink) {
		return false
	}

	return true
}
