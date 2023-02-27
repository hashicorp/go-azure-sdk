// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package threatintelligence

type ThreatIntelligenceInformationOperationPredicate struct {
}

func (p ThreatIntelligenceInformationOperationPredicate) Matches(input ThreatIntelligenceInformation) bool {

	return true
}
