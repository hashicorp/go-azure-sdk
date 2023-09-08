package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationSimulationUserCoverageOperationPredicate struct {
	ClickCount               *int64
	CompromisedCount         *int64
	LatestSimulationDateTime *string
	ODataType                *string
	SimulationCount          *int64
}

func (p AttackSimulationSimulationUserCoverageOperationPredicate) Matches(input AttackSimulationSimulationUserCoverage) bool {

	if p.ClickCount != nil && (input.ClickCount == nil || *p.ClickCount != *input.ClickCount) {
		return false
	}

	if p.CompromisedCount != nil && (input.CompromisedCount == nil || *p.CompromisedCount != *input.CompromisedCount) {
		return false
	}

	if p.LatestSimulationDateTime != nil && (input.LatestSimulationDateTime == nil || *p.LatestSimulationDateTime != *input.LatestSimulationDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SimulationCount != nil && (input.SimulationCount == nil || *p.SimulationCount != *input.SimulationCount) {
		return false
	}

	return true
}
